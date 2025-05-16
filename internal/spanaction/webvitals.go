package spanaction

type WebVitalsAttribute interface{}
type WebVitals struct{}

func NewWebVitals(attrGenerator WebVitalsAttribute) *WebVitals {
	return &WebVitals{}
}
