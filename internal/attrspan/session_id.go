package attrspan

import (
	"encoding/hex"

	"github.com/google/uuid"
	"go.opentelemetry.io/otel/attribute"
)

const (
	SpanAttributeSessionIDKey = attribute.Key("session.id")
)

func SessionIDKey(val string) attribute.KeyValue {
	return SpanAttributeSessionIDKey.String(val)
}

func GenerateSessionIDMock() string {
	id := uuid.New()
	return hex.EncodeToString(id[:])
}
