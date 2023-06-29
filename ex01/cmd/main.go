package main

import (
	"ex01/pkg/client"
	"fmt"
	"log"
)

const url string = "https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=250&page=1"
const coinId string = "bitcoin"

func main() {
	data := client.New()

	err := data.LoadCoinsData(url)
	if err != nil {
		log.Fatal(err)
	}
	PrintAllCouns(data.Coins)

	coin, err := data.GetCoinById(coinId)
	if err != nil {
		log.Fatal(err)
	}
	PrintCoin(coin)
}

func PrintAllCouns(coins []client.Coin) {
	for _, coin := range coins {
		PrintCoin(coin)
	}
}

func PrintCoin(coin client.Coin) {
	fmt.Println("ID:", coin.ID)
	fmt.Println("Symbol:", coin.Symbol)
	fmt.Println("Name:", coin.Name)
	fmt.Println("Image:", coin.Image)
	fmt.Println("Current Price:", coin.CurrentPrice)
	fmt.Println("Market Cap:", coin.MarketCap)
	fmt.Println("Market Cap Rank:", coin.MarketCapRank)
	fmt.Println("Fully Diluted Valuation:", coin.FullyDilutedValuation)
	fmt.Println("Total Volume:", coin.TotalVolume)
	fmt.Println("High 24h:", coin.High24h)
	fmt.Println("Low 24h:", coin.Low24h)
	fmt.Println("Price Change 24h:", coin.PriceChange24h)
	fmt.Println("Price Change Percentage 24h:", coin.PriceChangePercentage)
	fmt.Println("Market Cap Change 24h:", coin.MarketCapChange24h)
	fmt.Println("Market Cap Change Percentage 24h:", coin.MarketCapChangePercentage)
	fmt.Println("Circulating Supply:", coin.CirculatingSupply)
	fmt.Println("Total Supply:", coin.TotalSupply)
	fmt.Println("Max Supply:", coin.MaxSupply)
	fmt.Println("Ath:", coin.Ath)
	fmt.Println("Ath Change Percentage:", coin.AthChangePercentage)
	fmt.Println("Ath Date:", coin.AthDate)
	fmt.Println("Atl:", coin.Atl)
	fmt.Println("Atl Change Percentage:", coin.AtlChangePercentage)
	fmt.Println("Atl Date:", coin.AtlDate)
	fmt.Println("ROI:", coin.Roi)
	fmt.Println("Last Updated:", coin.LastUpdated)
	fmt.Println("---------------------------")
}
