package attrspan

import (
	"otel-generator/internal/attrresource"
	"otel-generator/internal/util"

	"go.opentelemetry.io/otel/attribute"
)

const SpanAttributeScreenNameKey = attribute.Key("screen.name")

type SpanAttributeScreenName struct {
	Android []string `yaml:"android"`
	IOS     []string `yaml:"ios"`
	Web     []string `yaml:"web"`
}

func (sn SpanAttributeScreenName) GetAttributes(serviceType attrresource.ServiceType) []string {
	switch serviceType {
	case attrresource.ServiceTypeIOS:
		return sn.IOS
	case attrresource.ServiceTypeAndroid:
		return sn.Android
	case attrresource.ServiceTypeWeb:
		return sn.Web
	default:
		return nil
	}
}

func (sg *SpanAttrGenerator) SetAttrScreenName(val string) attribute.KeyValue {
	return SpanAttributeScreenNameKey.String(val)
}

func (sg *SpanAttrGenerator) GenerateRandomScreenName() attribute.KeyValue {
	if len(sg.ScreenNames) == 0 {
		return attribute.KeyValue{}
	}

	pick, ok := util.PickRandomElementFromSlice[string](sg.ScreenNames)
	if !ok {
		return attribute.KeyValue{}
	}
	return sg.SetAttrScreenName(pick)
}
