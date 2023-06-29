package main

import (
	"ex02/pkg/csv"
	"ex02/pkg/parser"
	"fmt"
)

func main() {
	influencers, err := parser.Parse("https://hypeauditor.com/top-instagram-all-russia/")
	if err != nil {
		fmt.Println(err)
	}

	csv.AddToCSV("../another/csvFiles/csvOutput.csv", influencers)
}
