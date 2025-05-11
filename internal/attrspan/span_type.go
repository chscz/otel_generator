package attrspan

import "go.opentelemetry.io/otel/attribute"

const SpanSpanTypeKey = attribute.Key("span.type")

func SpanTypeKey(val string) attribute.KeyValue {
	return SpanSpanTypeKey.String(val)
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
