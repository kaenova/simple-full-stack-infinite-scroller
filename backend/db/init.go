package db

import (
	"backend/config"
	"backend/utils/errlogger"
	"fmt"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB = nil
var err error

func Init(tableDelete, dataInitialization bool) {

	log.Info().Msg("menginisialisasikan database")

	config := config.GetConfig()

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Database.Host,
		config.Database.Port,
		config.Database.Username,
		config.Database.Password,
		config.Database.Name)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true, Logger: logger.Default.LogMode(logger.Info)})
	errlogger.ErrFatalPanic(err)

	if tableDelete {
		log.Info().Msg("menghapus tabel yang ada")
		trx := db.Exec(`DROP SCHEMA public CASCADE;
		CREATE SCHEMA public; GRANT ALL ON SCHEMA public TO postgres;
		GRANT ALL ON SCHEMA public TO public;`)
		if trx.Error != nil {
			errlogger.ErrFatalPanic(trx.Error)
		}
	}

	Migration(db)

	if dataInitialization {
		initData(db)
	}

	DatabaseFinalCheck(db)

	log.Info().Msg("database terinisialisasi")
}

func GetDB() *gorm.DB {
	if db == nil {
		errlogger.FatalPanicMessage("db belum terinisilisasi")
	}
	return db
}
