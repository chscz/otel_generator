package config

import (
	"errors"
	"fmt"
)

const minAllowedInterval = 5

var (
	ErrInvalidInterval      = errors.New("invalid interval")
	ErrIntervalTooShort     = errors.New("interval must be at least 5 second")
	ErrInvalidIntervalRange = errors.New("invalid interval range: min interval cannot be greater than or equal to max interval")
)

type GenerateOption struct {
	UseDynamicInterval      bool `yaml:"use_dynamic_interval"`
	traceIntervalMinSeconds int  `yaml:"trace_interval_min_seconds"`
	traceIntervalMaxSeconds int  `yaml:"trace_interval_max_seconds"`
}

func (g *GenerateOption) validate() error {
	if g.traceIntervalMinSeconds <= 0 || g.traceIntervalMaxSeconds <= 0 {
		return fmt.Errorf("간격은 양수여야 합니다 (min: %s, max: %s)", g.traceIntervalMinSeconds, g.traceIntervalMaxSeconds)
	}
	if g.traceIntervalMinSeconds >= g.traceIntervalMaxSeconds {
		return fmt.Errorf(
			"최소 간격(%s)이 최대 간격(%s)보다 크거나 같을 수 없습니다",
			g.traceIntervalMinSeconds,
			g.traceIntervalMaxSeconds,
		)
	}
	if g.traceIntervalMinSeconds < minAllowedInterval {
		return fmt.Errorf(
			"최소 간격(%s)이 허용된 최솟값(%s)보다 작습니다",
			g.traceIntervalMinSeconds,
			minAllowedInterval,
		)
	}
	return nil
}
