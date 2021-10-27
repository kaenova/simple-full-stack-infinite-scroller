package db

import (
	"backend/entity"
	"encoding/csv"
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func initData(db *gorm.DB) {
	/*
		Use this function to make a initial data.
		You need to initialize your data first and the loop through the data.
		To Create Record please refer reading this https://gorm.io/docs/create.html
	*/

	// Kategori
	csvFile, err := os.Open("db/dummy/post.csv")
	if err != nil {
		fmt.Println(err)
	}
	log.Info().Msg("Successfully read csv")
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	for _, line := range csvLines {
		emp := entity.Post{
			Konten: line[0],
		}
		db.Create(&emp)
	}

	log.Info().Msg("dummy data terinisialisasi")
}
