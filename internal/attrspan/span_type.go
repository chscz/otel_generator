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
	spanTypes := GenerateSpanTypeMocks()
	spanType := spanTypes[rand.Intn(len(spanTypes))]
	return SpanAttributeSpanTypeKey.String(string(spanType))
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

func GenerateSpanTypeMocks() []SpanAttrSpanType {

	return []SpanAttrSpanType{
		SpanAttrSpanTypeRender,
		SpanAttrSpanTypeXHR,
		SpanAttrSpanTypeCrash,
		SpanAttrSpanTypeEvent,
		SpanAttrSpanTypeANR,
		SpanAttrSpanTypeError,
		SpanAttrSpanTypeWebVitals,
		SpanAttrSpanTypeLog,
	}
}
