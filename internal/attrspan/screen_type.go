package attrspan

import (
	"math/rand"

	"go.opentelemetry.io/otel/attribute"
)

const SpanAttributeScreenTypeKey = attribute.Key("screen.type")

func (sg *SpanAttrGenerator) ScreenTypeKey(val string) attribute.KeyValue {
	return SpanAttributeScreenTypeKey.String(val)
}

func (sg *SpanAttrGenerator) ScreenTypeRandomGenerate() attribute.KeyValue {
	screenTypes := GenerateScreenTypeMocks()
	screenType := screenTypes[rand.Intn(len(screenTypes))]
	return SpanAttributeScreenTypeKey.String(string(screenType))
}

type SpanAttrScreenType string

const (
	SpanAttrScreenTypePage        = "page"
	SpanAttrScreenTypeView        = "view"
	SpanAttrScreenTypeUnspecified = ""
)

func GenerateScreenTypeMocks() []SpanAttrScreenType {
	return []SpanAttrScreenType{
		SpanAttrScreenTypePage,
		SpanAttrScreenTypeView,
		//SpanAttrScreenTypeUnspecified,
	}
}
