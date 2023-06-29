package client

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type CoinsData struct {
	Coins []Coin
}

type Coin struct {
	ID                        string  `json:"id"`
	Symbol                    string  `json:"symbol"`
	Name                      string  `json:"name"`
	Image                     string  `json:"image"`
	CurrentPrice              float64 `json:"current_price"`
	MarketCap                 int64   `json:"market_cap"`
	MarketCapRank             int     `json:"market_cap_rank"`
	FullyDilutedValuation     int64   `json:"fully_diluted_valuation"`
	TotalVolume               float64 `json:"total_volume"`
	High24h                   float64 `json:"high_24h"`
	Low24h                    float64 `json:"low_24h"`
	PriceChange24h            float64 `json:"price_change_24h"`
	PriceChangePercentage     float64 `json:"price_change_percentage_24h"`
	MarketCapChange24h        float64 `json:"market_cap_change_24h"`
	MarketCapChangePercentage float64 `json:"market_cap_change_percentage_24h"`
	CirculatingSupply         float64 `json:"circulating_supply"`
	TotalSupply               float64 `json:"total_supply"`
	MaxSupply                 float64 `json:"max_supply"`
	Ath                       float64 `json:"ath"`
	AthChangePercentage       float64 `json:"ath_change_percentage"`
	AthDate                   string  `json:"ath_date"`
	Atl                       float64 `json:"atl"`
	AtlChangePercentage       float64 `json:"atl_change_percentage"`
	AtlDate                   string  `json:"atl_date"`
	Roi                       *ROI    `json:"roi,omitempty"`
	LastUpdated               string  `json:"last_updated"`
}

type ROI struct {
	Times      float64 `json:"times"`
	Currency   string  `json:"currency"`
	Percentage float64 `json:"percentage"`
}

func New() *CoinsData {
	return &CoinsData{}
}

func (c *CoinsData) LoadCoinsData(apiURL string) (err error) {
	response, err := http.Get(apiURL)
	if err != nil {
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &c.Coins)
	if err != nil {
		return
	}

	return
}

func (c CoinsData) GetCoinById(coinId string) (Coin, error) {
	for _, coin := range c.Coins {
		if coinId == coin.ID {
			return coin, nil
		}
	}

	return Coin{}, errors.New("Ничего не удалось найти по данному запросу")
}
