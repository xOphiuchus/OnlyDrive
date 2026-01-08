package db

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresSQLDB struct {
	db *gorm.DB
}

func NewPostgresInstance(DBHost string, DBPort int, DBUser string, DBPass string, DBName string) *PostgresSQLDB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		DBHost, DBUser, DBPass, DBName, DBPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create Postgres connection")
	}

	if err := db.Exec("SELECT datname FROM pg_database;").Error; err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to Postgres database")
	}

	log.Info().Msg("Connected to Postgres database successfully")

	return &PostgresSQLDB{
		db: db,
	}
}

func (p *PostgresSQLDB) GetDB() *gorm.DB {
	return p.db
}
