package config

type Config struct {
	GoroutineCount int
	CollectorURL   string
}

func LoadConfig() (Config, error) {
	cfg := Config{
		GoroutineCount: 10,
		CollectorURL:   "http://localhost:4318/v1/traces",
	}
	return cfg, nil
}
