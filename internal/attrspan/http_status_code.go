package attrspan

import (
	"math/rand"

	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.25.0"
)

func (sg *SpanAttrGenerator) SetAttrHTTPStatusCode(val int) attribute.KeyValue {
	return semconv.HTTPStatusCode(val)
}

func (sg *SpanAttrGenerator) GenerateRandomHTTPStatusCode() attribute.KeyValue {
	return semconv.HTTPStatusCode(sg.getWeightedRandomHttpStatusCode())
}

type httpStatusCodeChoice struct {
	Code   int
	Weight int
}

func (sg *SpanAttrGenerator) getWeightedRandomHttpStatusCode() int {
	totalWeight := 0
	for _, choice := range sg.HTTPStatusCodes {
		totalWeight += choice.Weight
	}
	if totalWeight == 0 {
		if len(sg.HTTPStatusCodes) == 0 {
			return 0
		}
		totalWeight = len(sg.HTTPMethods)
	}
	r := rand.Intn(totalWeight)

	upto := 0
	for _, choice := range sg.HTTPStatusCodes {
		if upto+choice.Weight > r {
			return choice.Code
		}
		upto += choice.Weight
	}

	return sg.HTTPStatusCodes[len(sg.HTTPStatusCodes)-1].Code
}

func setWeightedRandomHttpStatusCode() []httpStatusCodeChoice {
	return []httpStatusCodeChoice{
		{Code: 200, Weight: 60}, // OK
		{Code: 201, Weight: 5},  // Created
		{Code: 204, Weight: 5},  // No Content
		{Code: 302, Weight: 3},  // Found
		{Code: 400, Weight: 5},  // Bad Request
		{Code: 401, Weight: 5},  // Unauthorized
		{Code: 403, Weight: 2},  // Forbidden
		{Code: 404, Weight: 5},  // Not Found
		{Code: 500, Weight: 7},  // Internal Server Error
		{Code: 503, Weight: 3},  // Service Unavailable
	}
}
