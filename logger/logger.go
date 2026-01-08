package logger

import (
	"os"
	"sync"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var once sync.Once

func Init(debug bool) {
	once.Do(func() {
		console := zerolog.ConsoleWriter{Out: os.Stdout, NoColor: false}

		if debug {
			log.Logger = zerolog.New(console).With().Timestamp().Logger()
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
		} else {
			log.Logger = zerolog.New(console).With().Timestamp().Logger()
			zerolog.SetGlobalLevel(zerolog.InfoLevel)
		}
	})
}
