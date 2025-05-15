package attrspan

type SpanAttrGenerator struct {
	ScreenName SpanAttributeScreenName
	HTTPURL    []string
	UserID     []string
}

func NewSpanAttrGenerator(screenNames SpanAttributeScreenName, httpurls []string, userCount int) *SpanAttrGenerator {
	return &SpanAttrGenerator{
		ScreenName: screenNames,
		HTTPURL:    httpurls,
		UserID:     GenerateUserIDMocks(userCount),
	}
}
