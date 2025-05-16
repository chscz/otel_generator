package spanaction

type CrashAttribute interface{}

type Crash struct{}

func NewCrash(attrGenerator CrashAttribute) *Crash {
	return &Crash{}
}
