package attr

import (
	"encoding/hex"

	"github.com/google/uuid"
	"go.opentelemetry.io/otel/attribute"
)

const (
	ResourceSessionIDKey = attribute.Key("session.id")
)

func SessionIDKey(val string) attribute.KeyValue {
	return ResourceSessionIDKey.String(val)
}

func GenerateSessionIDMocks(n int) []string {
	sessionIDs := make([]string, n)
	for i := 0; i < n; i++ {
		id := uuid.New()
		sessionIDs[i] = hex.EncodeToString(id[:])
	}
	return sessionIDs
}
