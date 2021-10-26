package db

import (
	"backend/utils/errlogger"
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
	data, err := os.ReadFile("db/dummy/post.sql")
	errlogger.ErrFatalPanic(err)
	db.Exec(string(data))

	log.Info().Msg("dummy data terinisialisasi")
}
