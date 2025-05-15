package attrspan

import (
	"math/rand"

	"go.opentelemetry.io/otel/attribute"
)

const SpanAttributeSpanTypeKey = attribute.Key("span.type")

func (sg *SpanAttrGenerator) SpanTypeKey(val string) attribute.KeyValue {
	return SpanAttributeSpanTypeKey.String(val)
}

func (sg *SpanAttrGenerator) SpanTypeRandomGenerate() attribute.KeyValue {
	spanTypes := GenerateSpanTypeMocks()
	spanType := spanTypes[rand.Intn(len(spanTypes))]
	return SpanAttributeSpanTypeKey.String(string(spanType))
}

type SpanAttrSpanType string

const (
	SpanAttrSpanTypeRender    = "render"
	SpanAttrSpanTypeXHR       = "xhr"
	SpanAttrSpanTypeCrash     = "crash"
	SpanAttrSpanTypeEvent     = "event"
	SpanAttrSpanTypeANR       = "anr"
	SpanAttrSpanTypeError     = "error"
	SpanAttrSpanTypeWebVitals = "webvitals"
	SpanAttrSpanTypeLog       = "log"
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
