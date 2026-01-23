package mocks

type Config struct {
	Host string
}

func MakeMockCfg(host string) *Config {
	return &Config{Host: host}
}

func (s *Config) GetHost() string {
	return s.Host
}

func (s *Config) GetDatabaseURL() string {
	return ""
}
