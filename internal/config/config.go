package config

import (
	"fmt"
	"time"

	"github.com/alex-bodnar/lib/config"
	"github.com/alex-bodnar/lib/database"
	"github.com/alex-bodnar/lib/errs"
	"github.com/alex-bodnar/lib/log"
	"github.com/go-playground/validator/v10"
)

const (
	// DefaultPath - default path for config.
	DefaultPath = "./cmd/config.yaml"
)

type (
	// Config defines the properties of the application configuration.
	Config struct {
		Logger   log.Config `yaml:"logger" validate:"required,dive,required"`
		Storage  Storage    `yaml:"storage" validate:"required,dive,required"`
		Delivery Delivery   `yaml:"delivery" validate:"required,dive,required"`
		// Extra    Extra      `yaml:"extra"`
	}

	// Storage defines database engines configuration
	Storage struct {
		Postgres database.Config `yaml:"postgres" validate:"required,dive,required"`
	}

	// Delivery defines API server configuration.
	Delivery struct {
		HTTPServer HTTPServer `yaml:"http-server" validate:"required,dive,required"`
	}

	// HTTPServer defines HTTP section of the API server configuration.
	HTTPServer struct {
		LogRequests        bool          `yaml:"log-requests"`
		ListenAddress      string        `yaml:"listen-address" validate:"required"`
		ReadTimeout        time.Duration `yaml:"read-timeout" validate:"required"`
		WriteTimeout       time.Duration `yaml:"write-timeout" validate:"required"`
		BodySizeLimitBytes int           `yaml:"body-size-limit" validate:"required"`
		GracefulTimeout    int           `yaml:"graceful-timeout" validate:"required"`
	}
)

// New loads and validates all configuration data, returns filled Cfg - configuration data model.
func New(appName, cfgFilePath string) (*Config, error) {
	cfg := new(Config)

	if cfgErr := cfg.loadFromFile(cfgFilePath); cfgErr != nil {
		return nil, fmt.Errorf("config loader: %s", cfgErr)
	}

	return cfg.valid()
}

// loadFromFile loads configuration from file.
func (c *Config) loadFromFile(configPath string) error {
	if err := config.LoadFromFile(configPath, c); err != nil {
		return err
	}

	return nil
}

// valid validates configuration data.
func (c *Config) valid() (*Config, error) {
	validate := validator.New()

	if err := validate.Struct(c); err != nil {
		// TODO: rewrite this error
		return nil, &errs.FieldsValidation{Errors: []string{err.Error()}}
	}

	return c, nil
}
