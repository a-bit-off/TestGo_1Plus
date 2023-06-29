package main

import (
	"ex02/pkg/csv"
	"ex02/pkg/parser"
	"log"
)

func main() {
	influencers, err := parser.Parse("https://hypeauditor.com/top-instagram-all-russia/")
	if err != nil {
		log.Fatal(err)
	}

	err = csv.AddToCSV("../another/csvFiles/csvOutput.csv", influencers)
	if err != nil {
		log.Fatal(err)
	}
}
