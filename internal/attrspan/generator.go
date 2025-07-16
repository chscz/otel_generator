package attrspan

import (
	"otel-generator/internal/attrresource"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type SpanAttrGenerator struct {
	ServiceType          attrresource.ServiceType
	SpanTypes            []spanTypeChoice
	SessionID            string
	UserIDs              []string
	ScreenNames          []string
	HTTPURLs             []string
	HTTPMethods          []httpMethodChoice
	HTTPStatusCodes      []httpStatusCodeChoice
	ExceptionTypes       []string
	ExceptionMessages    []string
	ExceptionStackTraces []string
}

func NewSpanAttrGenerator(serviceType attrresource.ServiceType, screenNames SpanAttributeScreenName, httpurls, exceptionTypes, exceptionMessages, exceptionStackTraces []string, userCount int) *SpanAttrGenerator {
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
		ServiceType:          serviceType,
		SpanTypes:            setWeightedRandomSpanType(),
		SessionID:            GenerateSessionIDMock(),
		ScreenNames:          sn,
		UserIDs:              GenerateUserIDMocks(userCount),
		HTTPURLs:             httpurls,
		HTTPMethods:          setWeightedRandomHttpMethod(),
		HTTPStatusCodes:      setWeightedRandomHttpStatusCode(),
		ExceptionTypes:       exceptionTypes,
		ExceptionMessages:    exceptionMessages,
		ExceptionStackTraces: exceptionStackTraces,
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
