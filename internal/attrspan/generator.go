package attrspan

import (
	"otel-generator/internal/attrresource"

	"go.opentelemetry.io/otel/trace"
)

type AttributeSourceByServiceType interface {
	GetAttributes(serviceType attrresource.ServiceType) []string
}

type SpanAttrGenerator struct {
	ServiceType           attrresource.ServiceType
	SpanTypes             []spanTypeChoice
	SessionID             string
	UserIDs               []string
	ScreenNames           []string
	HTTPURLs              []string
	HTTPMethods           []httpMethodChoice
	HTTPStatusCodes       []httpStatusCodeChoice
	ExceptionTypes        []string
	ExceptionMessages     []string
	ExceptionStackTraces  []string
	NetworkConnectionType []networkConnectionTypeChoice
}

func NewSpanAttrGenerator(serviceType attrresource.ServiceType, spanAttrConfig SpanAttributes, userCount int) *SpanAttrGenerator {
	return &SpanAttrGenerator{
		ServiceType:           serviceType,
		SpanTypes:             setWeightedRandomSpanType(),
		SessionID:             GenerateSessionIDMock(),
		ScreenNames:           getAttributeByServiceType[SpanAttributeScreenName](serviceType, spanAttrConfig.ScreenNames),
		UserIDs:               GenerateUserIDMocks(userCount),
		HTTPURLs:              spanAttrConfig.HTTPURLs,
		HTTPMethods:           setWeightedRandomHttpMethod(),
		HTTPStatusCodes:       setWeightedRandomHttpStatusCode(),
		ExceptionTypes:        getAttributeByServiceType[SpanAttributeExceptionType](serviceType, spanAttrConfig.ExceptionTypes),
		ExceptionMessages:     getAttributeByServiceType[SpanAttributeExceptionMessage](serviceType, spanAttrConfig.ExceptionMessages),
		ExceptionStackTraces:  getAttributeByServiceType[SpanAttributeExceptionStackTrace](serviceType, spanAttrConfig.ExceptionStackTraces),
		NetworkConnectionType: setWeightedRandomNetworkConnectionType(),
	}
}

func (sg *SpanAttrGenerator) SetPopulateParentSpanAttributes(span trace.Span, spanType SpanAttrSpanType, userID string) InheritedSpanAttr {
	attrSpanType := sg.SetAttrSpanType(spanType)
	attrUserID := sg.SetAttrUserID(userID)
	attrSessionID := sg.SetAttrSessionID(sg.SessionID)
	attrScreenName := sg.GenerateRandomScreenName()
	attrScreenType := sg.GenerateRandomScreenType()

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
		sg.SetAttrSpanType(spanType),
		attr.UserID,
		attr.SessionID,
		attr.ScreenName,
		attr.ScreenType,
	)
}

func getAttributeByServiceType[T AttributeSourceByServiceType](serviceType attrresource.ServiceType, attr T) []string {
	return attr.GetAttributes(serviceType)
}
