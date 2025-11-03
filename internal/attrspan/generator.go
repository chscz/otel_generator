package attrspan

import (
	"context"
	"log"
	"math/rand"
	"time"

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
	WebVersion            []string

	needSessionRefresh                bool
	TimerSessionRefresh               *time.Timer
	minSessionIDRefreshIntervalMinute int
	maxSessionIDRefreshIntervalMinute int
}

func NewSpanAttrGenerator(
	serviceType attrresource.ServiceType,
	spanAttrConfig SpanAttributes,
	userCount, minSessionIDRefreshIntervalMinute, maxSessionIDRefreshIntervalMinute int,
) *SpanAttrGenerator {
	return &SpanAttrGenerator{
		ServiceType:                       serviceType,
		SpanTypes:                         setWeightedRandomSpanType(),
		SessionID:                         GenerateSessionIDMock(),
		ScreenNames:                       getAttributeByServiceType[SpanAttributeScreenName](serviceType, spanAttrConfig.ScreenNames),
		UserIDs:                           GenerateUserIDMocks(userCount),
		HTTPURLs:                          spanAttrConfig.HTTPURLs,
		HTTPMethods:                       setWeightedRandomHttpMethod(),
		HTTPStatusCodes:                   setWeightedRandomHttpStatusCode(),
		ExceptionTypes:                    getAttributeByServiceType[SpanAttributeExceptionType](serviceType, spanAttrConfig.ExceptionTypes),
		ExceptionMessages:                 getAttributeByServiceType[SpanAttributeExceptionMessage](serviceType, spanAttrConfig.ExceptionMessages),
		ExceptionStackTraces:              getAttributeByServiceType[SpanAttributeExceptionStackTrace](serviceType, spanAttrConfig.ExceptionStackTraces),
		NetworkConnectionType:             setWeightedRandomNetworkConnectionType(),
		WebVersion:                        spanAttrConfig.WebVersions,
		needSessionRefresh:                false,
		TimerSessionRefresh:               nil,
		minSessionIDRefreshIntervalMinute: minSessionIDRefreshIntervalMinute,
		maxSessionIDRefreshIntervalMinute: maxSessionIDRefreshIntervalMinute,
	}
}

func (sg *SpanAttrGenerator) SetPopulateParentSpanAttributes(span trace.Span, spanType SpanAttrSpanType, userID string) InheritedSpanAttr {
	attrSpanType := sg.SetAttrSpanType(spanType)
	attrUserID := sg.SetAttrUserID(userID)
	attrScreenName := sg.GenerateRandomScreenName()
	attrScreenType := sg.GenerateRandomScreenType()
	attrNetworkConnectionType := sg.GenerateRandomNetworkConnectionType()
	attrWebVersion := sg.GenerateRandomWebVersion()

	if sg.needSessionRefresh {
		beforeSessionID := sg.SessionID
		sg.SessionID = GenerateSessionIDMock()
		sg.needSessionRefresh = false
		log.Printf("session.id refresh::userID::%s   ###  before::%s -> after::%s", attrUserID.Value.AsString(), beforeSessionID, sg.SessionID)
	}
	attrSessionID := sg.SetAttrSessionID(sg.SessionID)

	span.SetAttributes(
		attrSpanType,
		attrUserID,
		attrSessionID,
		attrScreenName,
		attrScreenType,
		attrNetworkConnectionType,
		attrWebVersion,
	)
	return InheritedSpanAttr{
		SpanType:              attrSpanType,
		UserID:                attrUserID,
		SessionID:             attrSessionID,
		ScreenName:            attrScreenName,
		ScreenType:            attrScreenType,
		NetworkConnectionType: attrNetworkConnectionType,
		WebVersion:            attrWebVersion,
	}
}

func (sg *SpanAttrGenerator) SetPopulateChildSpanAttributes(span trace.Span, spanType SpanAttrSpanType, attr InheritedSpanAttr) {
	span.SetAttributes(
		sg.SetAttrSpanType(spanType),
		attr.UserID,
		attr.SessionID,
		attr.ScreenName,
		attr.ScreenType,
		attr.NetworkConnectionType,
		attr.WebVersion,
	)
}

func getAttributeByServiceType[T AttributeSourceByServiceType](serviceType attrresource.ServiceType, attr T) []string {
	return attr.GetAttributes(serviceType)
}

func (sg *SpanAttrGenerator) StartSessionIDRefreshTimer(ctx context.Context) {
	localRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	interval := localRand.Intn(sg.maxSessionIDRefreshIntervalMinute-sg.minSessionIDRefreshIntervalMinute+1) + sg.minSessionIDRefreshIntervalMinute
	sg.TimerSessionRefresh = time.NewTimer(time.Duration(interval) * time.Minute)
	for {
		select {
		case <-ctx.Done():
			return
		case <-sg.TimerSessionRefresh.C:
			sg.needSessionRefresh = true
			nextInterval := rand.Intn(sg.maxSessionIDRefreshIntervalMinute - sg.minSessionIDRefreshIntervalMinute)
			sg.TimerSessionRefresh.Reset(time.Duration(nextInterval) * time.Minute)
		}
	}
}
