package attrspan

import (
	"otel-generator/internal/attrresource"

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

func NewSpanAttrGenerator(serviceType attrresource.ServiceType, spanAttrConfig SpanAttributes, userCount int) *SpanAttrGenerator {
	return &SpanAttrGenerator{
		ServiceType:          serviceType,
		SpanTypes:            setWeightedRandomSpanType(),
		SessionID:            GenerateSessionIDMock(),
		ScreenNames:          getAttributeByServiceType[SpanAttributeScreenName](serviceType, spanAttrConfig.ScreenNames),
		UserIDs:              GenerateUserIDMocks(userCount),
		HTTPURLs:             spanAttrConfig.HTTPURLs,
		HTTPMethods:          setWeightedRandomHttpMethod(),
		HTTPStatusCodes:      setWeightedRandomHttpStatusCode(),
		ExceptionTypes:       getAttributeByServiceType[SpanAttributeExceptionType](serviceType, spanAttrConfig.ExceptionTypes),
		ExceptionMessages:    getAttributeByServiceType[SpanAttributeExceptionMessage](serviceType, spanAttrConfig.ExceptionMessages),
		ExceptionStackTraces: getAttributeByServiceType[SpanAttributeExceptionStackTrace](serviceType, spanAttrConfig.ExceptionStackTraces),
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

func (sg *SpanAttrGenerator) SetPopulateChildSpanAttributes(span trace.Span, spanType SpanAttrSpanType, attr InheritedSpanAttr) {
	span.SetAttributes(
		sg.SpanTypeKey(spanType),
		attr.UserID,
		attr.SessionID,
		attr.ScreenName,
		attr.ScreenType,
	)
}

func getAttributeByServiceType[T AttributeSourceByServiceType](serviceType attrresource.ServiceType, attr T) []string {
	return attr.GetAttributes(serviceType)
}
