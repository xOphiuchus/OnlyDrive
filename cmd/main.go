package main

import (
	"github.com/rs/zerolog/log"
	"github.com/xOphiuchus/OnlyDrive/config"
	"github.com/xOphiuchus/OnlyDrive/db"
	"github.com/xOphiuchus/OnlyDrive/logger"
)

func main() {

	config := config.NewConfig()

	logger.Init(config.APPEnv == "development" || config.LOGLevel == "debug")
	log.Info().Msg("Logger has been initialized properly")
	log.Debug().Msg("Debug Mode is set to true")

	log.Debug().Msg("Trying to connect to database")
	db := db.NewPostgresInstance(
		config.DBHost,
		config.DBPort,
		config.DBUser,
		config.DBPass,
		config.DBName,
	)
	log.Info().Msg("Database connection established successfully")
	_ = db.GetDB()

	log.Debug().Msgf("Initial stuff done now moving to API")

}
