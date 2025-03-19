package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/go-sanitize/sanitize"
	"github.com/rs/zerolog/log"
)

type Config struct {
	Path       string
	Sanitation Sanitation
}

func NewConfig() Config {
	return Config{
		Path:       "/api/v1/",
		Sanitation: NewSanitation(),
	}
}

type Sanitation struct {
	Validator *validator.Validate
	Sanitizer *sanitize.Sanitizer
}

func NewSanitation() Sanitation {
	sanitizer, err := sanitize.New()
	if err != nil {
		log.Error().Msgf("Error creating sanitizer: %v", err)
		panic(err)
	}

	return Sanitation{
		Validator: validator.New(validator.WithRequiredStructEnabled()),
		Sanitizer: sanitizer,
	}
}

type Path struct {
	Base  string
	Users string
	Auth  string
}
