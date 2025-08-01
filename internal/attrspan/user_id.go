package attrspan

import (
	"otel-generator/internal/util"

	"github.com/teris-io/shortid"
	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.30.0"
)

func (sg *SpanAttrGenerator) SetAttrUserID(val string) attribute.KeyValue {
	return semconv.UserID(val)
}

func (sg *SpanAttrGenerator) GenerateRandomUserID() attribute.KeyValue {
	return semconv.UserID(shortid.MustGenerate())
}

func (sg *SpanAttrGenerator) GetRandomUserID() string {
	userID, ok := util.PickRandomElementFromSlice[string](sg.UserIDs)
	if !ok {
		return ""
	}
	return userID
}

func GenerateUserIDMocks(count int) []string {
	userIDs := make([]string, count)
	for i := 0; i < count; i++ {
		userIDs[i] = shortid.MustGenerate()
	}
	return userIDs
}
