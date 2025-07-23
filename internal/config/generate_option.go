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
	UseDynamicInterval                bool `yaml:"use_dynamic_interval"`
	MinTraceIntervalSecond            int  `yaml:"min_trace_interval_second"`
	MaxTraceIntervalSecond            int  `yaml:"max_trace_interval_second"`
	MaxChildSpanCount                 int  `yaml:"max_child_span_count"`
	MaxSpanDurationMilliSecond        int  `yaml:"max_span_duration_millisecond"`
	UseSessionIDRefresh               bool `yaml:"use_session_id_refresh"`
	MinSessionIDRefreshIntervalMinute int  `yaml:"min_session_id_refresh_interval_minute"`
	MaxSessionIDRefreshIntervalMinute int  `yaml:"max_session_id_refresh_interval_minute"`
}

func (g *GenerateOption) validate() error {
	if g.MinTraceIntervalSecond <= 0 || g.MaxTraceIntervalSecond <= 0 {
		return fmt.Errorf("간격은 양수여야 합니다 (min: %d, max: %d)", g.MinTraceIntervalSecond, g.MaxTraceIntervalSecond)
	}
	if g.MinTraceIntervalSecond > g.MaxTraceIntervalSecond {
		return fmt.Errorf(
			"trace_interval 의 최소 간격(%d)이 최대 간격(%d)보다 클 수 없습니다",
			g.MinTraceIntervalSecond,
			g.MaxTraceIntervalSecond,
		)
	}
	if g.MinTraceIntervalSecond < minAllowedInterval {
		return fmt.Errorf(
			"최소 간격(%d)이 허용된 최솟값(%d)보다 작습니다",
			g.MinTraceIntervalSecond,
			minAllowedInterval,
		)
	}
	if g.MaxChildSpanCount < 0 {
		return fmt.Errorf("")
	}
	if g.MaxSpanDurationMilliSecond < 0 {
		return fmt.Errorf("")
	}
	if g.MinSessionIDRefreshIntervalMinute > g.MinSessionIDRefreshIntervalMinute {
		return fmt.Errorf(
			"session_id_refresh_interval 의 최소 간격(%d)이 최대 간격(%d)보다 클 수 없습니다",
			g.MinSessionIDRefreshIntervalMinute,
			g.MaxSessionIDRefreshIntervalMinute,
		)
	}
	return nil
}
