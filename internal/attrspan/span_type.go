package attrspan

import (
	"math/rand"

	"go.opentelemetry.io/otel/attribute"
)

const SpanAttributeSpanTypeKey = attribute.Key("span.type")

func (sg *SpanAttrGenerator) SpanTypeKey(val SpanAttrSpanType) attribute.KeyValue {
	return SpanAttributeSpanTypeKey.String(string(val))
}

func (sg *SpanAttrGenerator) SpanTypeRandomGenerate() attribute.KeyValue {
	return sg.SpanTypeKey(sg.getWeightedRandomSpanType())
}

type SpanAttrSpanType string

const (
	SpanAttrSpanTypeRender    SpanAttrSpanType = "render"
	SpanAttrSpanTypeXHR       SpanAttrSpanType = "xhr"
	SpanAttrSpanTypeCrash     SpanAttrSpanType = "crash"
	SpanAttrSpanTypeEvent     SpanAttrSpanType = "event"
	SpanAttrSpanTypeANR       SpanAttrSpanType = "anr"
	SpanAttrSpanTypeError     SpanAttrSpanType = "error"
	SpanAttrSpanTypeWebVitals SpanAttrSpanType = "webvitals"
	SpanAttrSpanTypeLog       SpanAttrSpanType = "log"
)

type spanTypeChoice struct {
	spanType SpanAttrSpanType
	Weight   int
}

func (sg *SpanAttrGenerator) getWeightedRandomSpanType() SpanAttrSpanType {
	totalWeight := 0
	for _, choice := range sg.SpanTypes {
		totalWeight += choice.Weight
	}

	r := rand.Intn(totalWeight)

	upto := 0
	for _, choice := range sg.SpanTypes {
		if upto+choice.Weight > r {
			return choice.spanType
		}
		upto += choice.Weight
	}

	return sg.SpanTypes[len(sg.SpanTypes)-1].spanType
}

func setWeightedRandomSpanType() []spanTypeChoice {
	return []spanTypeChoice{
		{spanType: SpanAttrSpanTypeXHR, Weight: 100},
		//{spanType: SpanAttrSpanTypeRender, Weight: 70},
		//{spanType: SpanAttrSpanTypeLog, Weight: 5},
		//{spanType: SpanAttrSpanTypeEvent, Weight: 60},
		//{spanType: SpanAttrSpanTypeANR, Weight: 1},
		//{spanType: SpanAttrSpanTypeCrash, Weight: 1},
		//{spanType: SpanAttrSpanTypeError, Weight: 1},
		//{spanType: SpanAttrSpanTypeWebVitals, Weight: 5},
	}
}

func (s SpanAttrSpanType) IsErrorSpanType() bool {
	if s == SpanAttrSpanTypeError || s == SpanAttrSpanTypeCrash || s == SpanAttrSpanTypeANR {
		return true
	}
	return false
}
