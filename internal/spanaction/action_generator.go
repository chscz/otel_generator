package spanaction

type ActionGenerator struct {
	Anr       *Anr
	Crash     *Crash
	Error     *Error
	Event     *Event
	Log       *Log
	Render    *Render
	WebVitals *WebVitals
	XHR       *XHR
}

func NewActionGenerator() *ActionGenerator {
	return &ActionGenerator{
		//Anr: NewAnr(attrspan.NewSpanAttrGenerator()),
	}
}
