package attrspan

import (
	"math/rand"

	"otel-generator/internal/attrresource"

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

func (sg *SpanAttrGenerator) ScreenNameKey(val string) attribute.KeyValue {
	return SpanAttributeScreenNameKey.String(val)
}

func (sg *SpanAttrGenerator) ScreenNameRandomGenerate() attribute.KeyValue {
	return sg.ScreenNameKey(sg.ScreenNames[rand.Intn(len(sg.ScreenNames))])
}
