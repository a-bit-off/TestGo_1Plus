package csv

import (
	"encoding/csv"
	"ex02/pkg/parser"
	"os"
	"strings"
)

func AddToCSV(fileName string, influencers []parser.Influencers) (err error) {
	// Создание нового файла CSV
	file, err := os.Create(fileName)
	if err != nil {
		return
	}
	defer file.Close()

	// Создание писателя CSV
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Запись заголовков столбцов в CSV
	headers := []string{"Rank", "ContributorName", "ContributorTitle", "Category",
		"Subscribers", "Audience", "Authentic", "Engagement"}
	err = writer.Write(headers)
	if err != nil {
		return
	}

	// запись данных
	for _, infl := range influencers {
		if infl.Rank == "" {
			continue
		}
		category := strings.Join(infl.Category, ", ")
		data := []string{
			infl.Rank,
			infl.ContributorName,
			infl.ContributorTitle,
			category,
			infl.Subscribers,
			infl.Audience,
			infl.Authentic,
			infl.Engagement,
		}
		err = writer.Write(data)
		if err != nil {
			return
		}
	}
	return
}
