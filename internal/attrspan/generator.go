package attrspan

import (
	"otel-generator/internal/attrresource"
	"otel-generator/internal/generator"

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
		SessionID:      GenerateSessionIDMocks(),
		ScreenName:     sn,
		HTTPURLs:       httpurls,
		HTTPMethods:    httpMethods,
		UserID:         GenerateUserIDMocks(userCount),
		HTTPStatusCode: setWeightedRandomHttpStatusCode(),
	}
}

func (sg *SpanAttrGenerator) SetPopulateSpanAttributes(span trace.Span, spanType SpanAttrSpanType, userID string) generator.InheritedSpanAttr {
	attrSpanType := sg.SpanTypeKey(spanType)
	attrUserID := sg.UserIDKey(userID)
	attrScreenName := sg.ScreenNameRandomGenerate()
	attrScreenType := sg.ScreenTypeRandomGenerate()
	span.SetAttributes(
		attrSpanType,
		attrUserID,
		attrScreenName,
		attrScreenType,
	)
	return generator.InheritedSpanAttr{
		UserID:     attrUserID,
		SpanType:   attrSpanType,
		ScreenName: attrScreenName,
		ScreenType: attrScreenType,
	}
}
