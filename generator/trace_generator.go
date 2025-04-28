package generator

import "otel-generator/domain"

type TraceGenerator struct {
	Resource map[int]domain.Service
}
