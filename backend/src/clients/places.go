package clients

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	pb "github.com/coop-cmpers/what2do-backend/protos-gen/what2do/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (c *googlePlacesClient) FetchPlaces(ctx context.Context, searchType string, location string, eventTime *timestamppb.Timestamp) ([]*pb.Recommendation, error) {
	// create the body of the request from the provided arguments
	textQuery := fmt.Sprintf("%s, in %s", searchType, location)
	values := map[string]string{"textQuery": textQuery}
	jsonValue, err := json.Marshal(values)
	if err != nil {
		c.logger.Errorf("Failed to convert string (values) to json (jsonValue) - err: %v", err)
		return nil, err
	}

	// create the POST request object
	req, err := http.NewRequest("POST", c.baseURL+":searchText", bytes.NewBuffer(jsonValue))
	if err != nil {
		c.logger.Errorf("Failed to make post request object - err: %v", err)
		return nil, err
	}

	// set the header of the POST request object to specify what we want from the PlacesAPI
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Goog-Api-Key", c.apiKey)
	req.Header.Set("X-Goog-FieldMask", "places.displayName,places.formattedAddress,places.priceLevel,places.currentOpeningHours")

	// create a client and make the request
	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		c.logger.Errorf("Failed to make post request to placeAPI - err: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	// read the response status and body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.logger.Errorf("Failed to read response body - err: %v", err)
		return nil, err
	}

	var places Places
	err = json.Unmarshal(body, &places)
	if err != nil {
		c.logger.Errorf("Could not Unmarshal (Decode) the JSON from response - err: %v", err)
		return nil, err
	}

	if len(places.Places) == 0 {
		return []*pb.Recommendation{}, nil
	}

	var finalPlaces Places
	var opens time.Time
	var closes time.Time
	now := eventTime.AsTime()

	// add a binary search to optimise the search process for matching date in each place
	for place := range len(places.Places) {
		dates := places.Places[place].OpeningHours.Periods

		for date := range len(dates) {
			opens = time.Date(
				int(dates[date].Open.Date.Year),
				time.Month(dates[date].Open.Date.Month),
				int(dates[date].Open.Date.Day),
				int(dates[date].Open.Hour),
				int(dates[date].Open.Minute),
				0,
				0,
				time.Local,
			)
			closes = time.Date(
				int(dates[date].Close.Date.Year),
				time.Month(dates[date].Close.Date.Month),
				int(dates[date].Close.Date.Day),
				int(dates[date].Close.Hour),
				int(dates[date].Close.Minute),
				0,
				0,
				time.Local,
			)

			if opens.Before(now) && closes.After(now) {
				places.Places[place].Open = opens
				places.Places[place].Close = closes
				finalPlaces.Places = append(finalPlaces.Places, places.Places[place])
			}
		}
	}

	numberOfRecommendations := len(finalPlaces.Places)
	if numberOfRecommendations >= 5 {
		numberOfRecommendations = 5
	}

	var recommendations []*pb.Recommendation
	for i := range numberOfRecommendations {
		recommendation := &pb.Recommendation{
			Rank:       int32(i),
			Name:       finalPlaces.Places[i].DisplayName.Text,
			Address:    finalPlaces.Places[i].FormattedAddress,
			PriceLevel: finalPlaces.Places[i].PriceLevel,
			Open:       timestamppb.New(finalPlaces.Places[i].Open),
			Close:      timestamppb.New(finalPlaces.Places[i].Close),
		}
		recommendations = append(recommendations, recommendation)
	}

	return recommendations, nil
}
