package config

import (
	"os"

	"otel-generator/internal/attrresource"
	"otel-generator/internal/attrspan"

	"gopkg.in/yaml.v2"
)

type Config struct {
	GoroutineCount int                    `yaml:"go_routine_count"`
	CollectorURL   string                 `yaml:"collector_url"`
	UserCount      int                    `yaml:"user_count"`
	Services       []attrresource.Service `yaml:"services"`
	SpanAttributes SpanAttributes         `yaml:"span_attribute"`
}

func LoadConfig() (*Config, error) {
	cfg := &Config{
		GoroutineCount: 0,
		UserCount:      0,
		//CollectorURL:   "http://localhost:4318/v1/traces",
		CollectorURL: "",
		SpanAttributes: SpanAttributes{
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
	b, err := os.ReadFile("./config.yaml")
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(b, &cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
