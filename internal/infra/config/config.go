package config

import "github.com/alecthomas/kong"

type Config struct {
	Debug       bool   `help:"Application mode" env:"DEBUG" required:"true"`
	Host        string `help:"Port to listen on"       env:"HOST"          default:":4000"`
	DatabaseURL string `help:"Database connection URL" env:"DATABASE_URL"  required:"true"`
}

func Make() *Config {
	cfg := &Config{}
	parser, err := kong.New(cfg)
	if err != nil {
		panic(err)
	}

	// Parse command-line flags, environment variables, and config file
	_, err = parser.Parse(nil)
	if err != nil {
		panic(err)
	}
	return cfg
}

func (s *Config) GetHost() string {
	return s.Host
}

func (s *Config) GetDatabaseURL() string {
	return s.DatabaseURL
}
