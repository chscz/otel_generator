package config

import (
	"fmt"
	"os"

	"otel-generator/internal/attrresource"
	"otel-generator/internal/attrspan"

	"github.com/caarlos0/env/v11"
	"gopkg.in/yaml.v2"
)

type Config struct {
	CollectorURL       string                          `yaml:"collector_url" env:"COLLECTOR_URL"`
	GoRoutineCount     int                             `yaml:"go_routine_count" env:"GO_ROUTINE_COUNT"`
	UserCount          int                             `yaml:"user_count" env:"USER_COUNT"`
	GenerateOption     GenerateOption                  `yaml:"generate" envPrefix:"GENERATE_OPTION_"`
	Services           []attrresource.Service          `yaml:"services"`
	SpanAttributes     attrspan.SpanAttributes         `yaml:"span_attribute"`
	ResourceAttributes attrresource.ResourceAttributes `yaml:"resource_attribute"`
}

func LoadConfig(configFilePath string) (*Config, error) {
	cfg := &Config{
		//CollectorURL:   "http://localhost:4318/v1/traces",
		GoRoutineCount: 0,
		UserCount:      0,
		GenerateOption: GenerateOption{
			UseDynamicInterval:         false,
			MinTraceIntervalSecond:     10,
			MaxTraceIntervalSecond:     60,
			MaxChildSpanCount:          15,
			MaxSpanDurationMilliSecond: 5432,
		},
		SpanAttributes: attrspan.SpanAttributes{
			ScreenNames: attrspan.SpanAttributeScreenName{
				Android: []string{
					"MainActivity",
				},
				IOS: []string{
					"ios-test-screen-name-0",
					"ios-test-screen-name-0",
				},
				Web: []string{
					"web-test-screen-name-0",
					"web-test-screen-name-9",
				},
			},
			HTTPURLs: []string{
				"www.google.com",
				"www.github.com",
			},
		},
	}
	b, err := os.ReadFile(configFilePath)
	if err != nil {
		return nil, err
	}

	if err = yaml.Unmarshal(b, &cfg); err != nil {
		return nil, err
	}

	if err = env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, cfg.validate()
}

func (c *Config) validate() error {
	if c.CollectorURL == "" || c.GoRoutineCount <= 0 || c.UserCount <= 0 {
		return fmt.Errorf("invalid config: collector_url, goroutine_count, user_count")
	}
	if len(c.Services) == 0 {
		return fmt.Errorf("invalid config: no services specified")
	}
	return c.GenerateOption.validate()
}
