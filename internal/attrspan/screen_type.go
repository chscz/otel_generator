package attrspan

import (
	"math/rand"

	"otel-generator/internal/attrresource"

	"go.opentelemetry.io/otel/attribute"
)

const SpanAttributeScreenTypeKey = attribute.Key("screen.type")

func (sg *SpanAttrGenerator) SetAttrScreenType(val string) attribute.KeyValue {
	return SpanAttributeScreenTypeKey.String(val)
}

func (sg *SpanAttrGenerator) GenerateRandomScreenType() attribute.KeyValue {
	screenTypeMap := GenerateScreenTypeMocks()
	screenTypes, ok := screenTypeMap[sg.ServiceType]
	if !ok || len(screenTypes) == 0 {
		return attribute.KeyValue{}
	}
	screenType := screenTypes[rand.Intn(len(screenTypes))]
	return SpanAttributeScreenTypeKey.String(screenType)
}

type SpanAttrScreenTypeMap map[attrresource.ServiceType][]string
type SpanAttrScreenType string

const (
	SpanAttrScreenTypePage        = "page"
	SpanAttrScreenTypeView        = "view"
	SpanAttrScreenTypeActivity    = "activity"
	SpanAttrScreenTypeFragment    = "fragment"
	SpanAttrScreenTypeUnspecified = ""
)

func GenerateScreenTypeMocks() SpanAttrScreenTypeMap {
	return SpanAttrScreenTypeMap{
		attrresource.ServiceTypeAndroid: []string{
			SpanAttrScreenTypeActivity,
			SpanAttrScreenTypeFragment,
		},
		attrresource.ServiceTypeIOS: []string{
			SpanAttrScreenTypePage,
			SpanAttrScreenTypeView,
		},
		attrresource.ServiceTypeWeb: []string{},
	}
}
