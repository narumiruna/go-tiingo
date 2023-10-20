package api

import (
	"time"

	"github.com/c9s/requestgen"
)

//go:generate requestgen -method GET -url "/tiingo/daily/:ticker/prices" -type PriceRequest -responseType []PriceResponse
type PriceRequest struct {
	Client requestgen.AuthenticatedAPIClient

	Ticker         string     `param:"ticker,slug"`
	StartDate      *time.Time `param:"startDate"`
	EndDate        *time.Time `param:"endDate"`
	ResampleFreq   *string    `param:"resampleFreq"`
	Sort           *string    `param:"sort"`
	ResponseFormat *string    `param:"format"`
	Columns        *[]string  `param:"columns"`
}

type PriceResponse struct {
	Date        time.Time `json:"date"`
	Close       float64   `json:"close"`
	High        float64   `json:"high"`
	Low         float64   `json:"low"`
	Open        float64   `json:"open"`
	Volume      int64     `json:"volume"`
	AdjClose    float64   `json:"adjClose"`
	AdjHigh     float64   `json:"adjHigh"`
	AdjLow      float64   `json:"adjLow"`
	AdjOpen     float64   `json:"adjOpen"`
	AdjVolume   int64     `json:"adjVolume"`
	DivCash     float64   `json:"divCash"`
	SplitFactor float64   `json:"splitFactor"`
}
