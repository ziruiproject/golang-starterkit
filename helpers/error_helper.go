package helpers

import (
	"github.com/rs/zerolog/log"
)

func PanicOnError(err error) {
	if err != nil {
		log.Error().Msgf("System error: %s", err)
		panic(err)
	}
}
