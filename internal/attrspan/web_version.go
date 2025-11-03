package attrspan

import (
	"otel-generator/internal/util"

	"go.opentelemetry.io/otel/attribute"
)

const SpanAttributeWebVersionKey = attribute.Key("web.version")

func (sg *SpanAttrGenerator) SetAttrWebVersion(val string) attribute.KeyValue {
	return SpanAttributeWebVersionKey.String(val)
}

func (sg *SpanAttrGenerator) GenerateRandomWebVersion() attribute.KeyValue {
	if len(sg.WebVersion) == 0 {
		return attribute.KeyValue{}
	}

	pick, ok := util.PickRandomElementFromSlice[string](sg.WebVersion)
	if !ok {
		return attribute.KeyValue{}
	}
	return sg.SetAttrWebVersion(pick)
}
