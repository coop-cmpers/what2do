package clients

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	pb "github.com/coop-cmpers/what2do-backend/protos-gen/what2do/v1"
	"github.com/coop-cmpers/what2do-backend/src/helpers"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Name struct {
	Text string `json:"text"`
	LanguageCode string `json:"languageCode"`
}

type Date struct {
	Year int32 `json:"year"`
	Month int32 `json:"month"`
	Day int32 `json:"day"`
}

type DayDetails struct {
	Day int32 `json:"day"`
	Hour int32 `json:"hour"`
	Minute int32 `json:"minute"`
	Date Date `json:"date"`
}

type Day struct {
	Open DayDetails `json:"open"`
	Close DayDetails `json:"close"`
}

type CurrentOpeningHours struct {
	Periods []Day `json:"periods"`
}

type Place struct {
	FormattedAddress string `json:"formattedAddress"`
	PriceLevel string `json:"priceLevel"`
	DisplayName Name `json:"displayName"`
	OpeningHours CurrentOpeningHours `json:"currentOpeningHours"`
	Open time.Time
	Close time.Time
}

type Places struct {
	Places []Place `json:"places"`
}

func FetchPlacesFromPlacesAPI(ctx context.Context, searchType string, location string, eventTime *timestamppb.Timestamp) ([]*pb.Recommendation, error) {

	envVariables, err := helpers.GetEnvFromCtx(ctx)
	if err != nil {
		log.Printf("Failed to access environment variable - err: %v", err)
	}


	// create the body of the request from the provided arguments
	textQuery := fmt.Sprintf("%s, in %s", searchType, location)
	values := map[string]string{"textQuery": textQuery}
	jsonValue, err := json.Marshal(values)
	if err != nil {
		log.Printf("Failed to convert string (values) to json (jsonValue) - err: %v", err)
		return nil, err
	}

	// create the POST request object
	req, err := http.NewRequest("POST", envVariables["GOOGLE_PLACE_API_URL"], bytes.NewBuffer(jsonValue))
	if err != nil {
		log.Printf("Failed to make post request object - err: %v", err)
		return nil, err
	}

	// set the header of the POST request object to specify what we want from the PlacesAPI
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Goog-Api-Key", envVariables["GOOGLE_PLACE_API_KEY"])
	req.Header.Set("X-Goog-FieldMask", "places.displayName,places.formattedAddress,places.priceLevel,places.currentOpeningHours")

	// create a client and make the request
	placesapiClient := &http.Client{}
    resp, err := placesapiClient.Do(req)
    if err != nil {
        log.Printf("Failed to make post request to placeAPI - err: %v", err)
		return nil, err
    }
    defer resp.Body.Close()

	// log the response status and body
	log.Printf("Response Status: " + resp.Status)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
        log.Printf("Failed to read response body - err: %v", err)
		return nil, err
    }
	

	var places Places
	err = json.Unmarshal(body, &places)
	if err != nil {
		log.Printf("Could not Unmarshal (Decode) the JSON from response - err: %v", err)
		return nil, err
	}

	if len(places.Places) == 0 {
		log.Printf("There were no places returned from the PlacesAPI")
		return nil, errors.New("there were no places returned from the PlacesAPI")
	}

	var finalPlaces Places
	now := eventTime.AsTime()
	var opens time.Time
	var closes time.Time
	// add a binary search to optimise the search process for matching date in each place
	for place := range len(places.Places) {
		dates := places.Places[place].OpeningHours.Periods
		for date := range len(dates) {
			// chris will probably yell at me for this epic variable assignment but lmk if you have a better idea for the assignment lol
			opens = time.Date(int(dates[date].Open.Date.Year), time.Month(dates[date].Open.Date.Month), int(dates[date].Open.Date.Day), int(dates[date].Open.Hour), int(dates[date].Open.Minute), 0, 0, time.Local)
			closes = time.Date(int(dates[date].Close.Date.Year), time.Month(dates[date].Close.Date.Month), int(dates[date].Close.Date.Day), int(dates[date].Close.Hour), int(dates[date].Close.Minute), 0, 0, time.Local)
			if opens.Before(now) && closes.After(now) {
				places.Places[place].Open = opens
				places.Places[place].Close = closes
				finalPlaces.Places = append(finalPlaces.Places, places.Places[place])
			}
		}
	}


	var recommendationReturnList []*pb.Recommendation
	amountOfRecommendations := 0
	if len(finalPlaces.Places) >= 5 {
		amountOfRecommendations = 5
	} else {
		amountOfRecommendations = len(finalPlaces.Places)
	}
	for i := range amountOfRecommendations {
		recommendationCreater := pb.Recommendation{
			Rank: int32(i),
			Name: finalPlaces.Places[i].DisplayName.Text,
			Address: finalPlaces.Places[i].FormattedAddress,
			PriceLevel: finalPlaces.Places[i].PriceLevel,
			Open: timestamppb.New(finalPlaces.Places[i].Open),
			Close: timestamppb.New(finalPlaces.Places[i].Close),
		}
		recommendationReturnList = append(recommendationReturnList, &recommendationCreater)
	}

	
	return recommendationReturnList, nil
}