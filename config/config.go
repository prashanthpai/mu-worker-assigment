package config

import (
	validator "github.com/go-playground/validator/v10"

	"github.com/Unacademy/mu-worker/pkg/db"
	"github.com/Unacademy/mu-worker/pkg/queue"
)

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

func init() {
	validate = validator.New()
}

type Config struct {
	DB    *db.Config    `toml:"db" validate:"required"`
	Queue *queue.Config `toml:"queue" validate:"required"`
}

func NewConfig(path string) (*Config, error) {
	// load config from file

	// override config from environment variables

	// validate structure using validator

	return nil, nil
}
