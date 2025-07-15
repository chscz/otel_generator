package attrspan

import (
	"otel-generator/internal/attrresource"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type SpanAttrGenerator struct {
	ServiceType    attrresource.ServiceType
	SessionID      string
	UserID         []string
	ScreenName     []string
	HTTPURLs       []string
	HTTPMethods    []string
	HTTPStatusCode []statusCodeChoice
}

func NewSpanAttrGenerator(serviceType attrresource.ServiceType, screenNames SpanAttributeScreenName, httpurls, httpMethods []string, userCount int) *SpanAttrGenerator {
	var sn []string
	switch serviceType {
	case attrresource.ServiceTypeAndroid:
		sn = screenNames.Android
	case attrresource.ServiceTypeIOS:
		sn = screenNames.IOS
	case attrresource.ServiceTypeWeb:
		sn = screenNames.Web
	default:
	}

	return &SpanAttrGenerator{
		ServiceType:    serviceType,
		SessionID:      GenerateSessionIDMock(),
		ScreenName:     sn,
		HTTPURLs:       httpurls,
		HTTPMethods:    httpMethods,
		UserID:         GenerateUserIDMocks(userCount),
		HTTPStatusCode: setWeightedRandomHttpStatusCode(),
	}
}

func (sg *SpanAttrGenerator) SetPopulateParentSpanAttributes(span trace.Span, spanType SpanAttrSpanType, userID string) InheritedSpanAttr {
	attrSpanType := sg.SpanTypeKey(spanType)
	attrUserID := sg.UserIDKey(userID)
	attrSessionID := sg.SessionIDKey(sg.SessionID)
	attrScreenName := sg.ScreenNameRandomGenerate()
	attrScreenType := sg.ScreenTypeRandomGenerate()

	span.SetAttributes(
		attrSpanType,
		attrUserID,
		attrSessionID,
		attrScreenName,
		attrScreenType,
	)
	return InheritedSpanAttr{
		SpanType:   attrSpanType,
		UserID:     attrUserID,
		SessionID:  attrSessionID,
		ScreenName: attrScreenName,
		ScreenType: attrScreenType,
	}
}

func (sg *SpanAttrGenerator) SetPopulateChildSpanAttributes(span trace.Span, attr InheritedSpanAttr) {
	span.SetAttributes(
		attr.SpanType,
		attr.UserID,
		attr.SessionID,
		attr.ScreenName,
		attr.ScreenType,
	)
}

type InheritedSpanAttr struct {
	SessionID  attribute.KeyValue
	UserID     attribute.KeyValue
	SpanType   attribute.KeyValue
	ScreenName attribute.KeyValue
	ScreenType attribute.KeyValue
}
