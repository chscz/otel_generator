package attrspan

import (
	"math/rand"

	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.25.0"
)

func (sg *SpanAttrGenerator) GenerateRandomNetworkConnectionType() attribute.KeyValue {
	return sg.getWeightedRandomNetworkConnectionType()
}

type networkConnectionTypeChoice struct {
	NetworkConnectionType attribute.KeyValue
	Weight                int
}

func (sg *SpanAttrGenerator) getWeightedRandomNetworkConnectionType() attribute.KeyValue {
	totalWeight := 0
	for _, choice := range sg.NetworkConnectionType {
		totalWeight += choice.Weight
	}
	if totalWeight == 0 {
		if len(sg.NetworkConnectionType) == 0 {
			return attribute.KeyValue{}
		}
		totalWeight = len(sg.NetworkConnectionType)
	}

	r := rand.Intn(totalWeight)

	upto := 0
	for _, choice := range sg.NetworkConnectionType {
		if upto+choice.Weight > r {
			return choice.NetworkConnectionType
		}
		upto += choice.Weight
	}

	return sg.NetworkConnectionType[len(sg.NetworkConnectionType)-1].NetworkConnectionType
}

func setWeightedRandomNetworkConnectionType() []networkConnectionTypeChoice {
	return []networkConnectionTypeChoice{
		{NetworkConnectionType: semconv.NetworkConnectionTypeWifi, Weight: 1},
		{NetworkConnectionType: semconv.NetworkConnectionTypeWired, Weight: 1},
		{NetworkConnectionType: semconv.NetworkConnectionTypeCell, Weight: 1},
		{NetworkConnectionType: semconv.NetworkConnectionTypeUnavailable, Weight: 1},
		{NetworkConnectionType: semconv.NetworkConnectionTypeUnknown, Weight: 1},
	}
}
