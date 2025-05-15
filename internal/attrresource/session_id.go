package attrresource

import (
	"encoding/hex"

	"github.com/google/uuid"
	"go.opentelemetry.io/otel/attribute"
)

const (
	ResourceAttributeSessionIDKey = attribute.Key("session.id")
)

func SessionIDKey(val string) attribute.KeyValue {
	return ResourceAttributeSessionIDKey.String(val)
}

func GenerateSessionIDMocks() string {
	id := uuid.New()
	return hex.EncodeToString(id[:])
}
