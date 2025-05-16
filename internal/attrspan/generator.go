package attrspan

import "go.opentelemetry.io/otel/trace"

type SpanAttrGenerator struct {
	UserID         []string
	ScreenName     SpanAttributeScreenName
	HTTPURL        []string
	HTTPStatusCode []statusCodeChoice
}

func NewSpanAttrGenerator(screenNames SpanAttributeScreenName, httpurls []string, userCount int) *SpanAttrGenerator {
	return &SpanAttrGenerator{
		ScreenName:     screenNames,
		HTTPURL:        httpurls,
		UserID:         GenerateUserIDMocks(userCount),
		HTTPStatusCode: setWeightedRandomHttpStatusCode(),
	}
}

func (sg *SpanAttrGenerator) PopulateSpanAttributes(span trace.Span, userID string) SpanAttrSpanType {
	spanType := sg.SpanTypeRandomGenerate()
	span.SetAttributes(
		spanType,
		sg.UserIDKey(userID),
		sg.ScreenNameRandomGenerate(),
		sg.ScreenTypeRandomGenerate(),
	)
	return SpanAttrSpanType(spanType.Value.AsString())
}
