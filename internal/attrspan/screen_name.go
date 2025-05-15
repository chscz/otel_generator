package attrspan

import (
	"math/rand"

	"go.opentelemetry.io/otel/attribute"
)

const SpanAttributeScreenNameKey = attribute.Key("screen.name")

type SpanAttributeScreenName struct {
	Android []string `yaml:"android"`
	IOS     []string `yaml:"ios"`
	Web     []string `yaml:"web"`
}

func (sg *SpanAttrGenerator) ScreenNameKey(val string) attribute.KeyValue {
	return SpanAttributeScreenNameKey.String(val)
}

func (sg *SpanAttrGenerator) ScreenNameRandomGenerate() attribute.KeyValue {
	screenNames := GenerateScreenNameMocks()
	screenName := screenNames[rand.Intn(len(screenNames))]
	return SpanAttributeScreenNameKey.String(screenName)
}

func GenerateScreenNameMocks() []string {
	return []string{
		"MainActivity",
		"test-screen-name-1",
		"test-screen-name-2",
	}
}
