package clients

import "time"

type Name struct {
	Text         string `json:"text"`
	LanguageCode string `json:"languageCode"`
}

type Date struct {
	Year  int32 `json:"year"`
	Month int32 `json:"month"`
	Day   int32 `json:"day"`
}

type DayDetails struct {
	Day    int32 `json:"day"`
	Hour   int32 `json:"hour"`
	Minute int32 `json:"minute"`
	Date   Date  `json:"date"`
}

type Day struct {
	Open  DayDetails `json:"open"`
	Close DayDetails `json:"close"`
}

type CurrentOpeningHours struct {
	Periods []Day `json:"periods"`
}

type Place struct {
	FormattedAddress string              `json:"formattedAddress"`
	PriceLevel       string              `json:"priceLevel"`
	DisplayName      Name                `json:"displayName"`
	OpeningHours     CurrentOpeningHours `json:"currentOpeningHours"`
	Open             time.Time
	Close            time.Time
}

type Places struct {
	Places []Place `json:"places"`
}
