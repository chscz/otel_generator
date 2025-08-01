package attrspan

import (
	"math/rand"
	"net/http"

	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.25.0"
)

const SpanHTTPMethodKey = attribute.Key("http.method")

func (sg *SpanAttrGenerator) SetAttrHTTPMethod(val string) attribute.KeyValue {
	return semconv.HTTPMethod(val)
}

func (sg *SpanAttrGenerator) GenerateRandomHTTPMethod() attribute.KeyValue {
	return sg.SetAttrHTTPMethod(sg.getWeightedRandomHttpMethod())
}

type httpMethodChoice struct {
	Method string
	Weight int
}

func (sg *SpanAttrGenerator) getWeightedRandomHttpMethod() string {
	totalWeight := 0
	for _, choice := range sg.HTTPMethods {
		totalWeight += choice.Weight
	}
	if totalWeight == 0 {
		if len(sg.HTTPMethods) == 0 {
			return ""
		}
		totalWeight = len(sg.HTTPMethods)
	}

	r := rand.Intn(totalWeight)

	upto := 0
	for _, choice := range sg.HTTPMethods {
		if upto+choice.Weight > r {
			return choice.Method
		}
		upto += choice.Weight
	}

	return sg.HTTPMethods[len(sg.HTTPMethods)-1].Method
}

func setWeightedRandomHttpMethod() []httpMethodChoice {
	return []httpMethodChoice{
		{Method: http.MethodGet, Weight: 50},
		{Method: http.MethodPost, Weight: 50},
		{Method: http.MethodPut, Weight: 5},
		{Method: http.MethodPatch, Weight: 5},
		{Method: http.MethodDelete, Weight: 5},
		{Method: http.MethodOptions, Weight: 5},
		{Method: http.MethodHead, Weight: 1},
		{Method: http.MethodTrace, Weight: 1},
		{Method: http.MethodConnect, Weight: 1},
	}
}
