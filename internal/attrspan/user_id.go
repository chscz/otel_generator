package attrspan

import (
	"math/rand"

	"github.com/teris-io/shortid"
	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.30.0"
)

func (sg *SpanAttrGenerator) UserIDKey(val string) attribute.KeyValue {
	return semconv.UserID(val)
}

func (sg *SpanAttrGenerator) UserIDRandomGenerate() attribute.KeyValue {
	return semconv.UserID(shortid.MustGenerate())
}

func (sg *SpanAttrGenerator) GetRandomUserID() string {
	return sg.UserID[rand.Intn(len(sg.UserID))]
}

func GenerateUserIDMocks(count int) []string {
	userIDs := make([]string, count)
	for i := 0; i < count; i++ {
		userIDs[i] = shortid.MustGenerate()
	}
	return userIDs
}
