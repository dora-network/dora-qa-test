package data

import (
	"encoding/json"
	"time"
)

type Tick struct {
	AssetID   string    `json:"assetID"`
	Timestamp time.Time `json:"timestamp"`
	LastPrice float64   `json:"lastPrice,omitempty"`
	LastSize  float64   `json:"lastSize,omitempty"`
	BestBid   float64   `json:"bestBid,omitempty"`
	BestAsk   float64   `json:"bestAsk,omitempty"`
}

func (t Tick) MarshalBinary() ([]byte, error) {
	return json.Marshal(t)
}
