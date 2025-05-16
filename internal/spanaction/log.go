package spanaction

type LogAttribute interface{}
type Log struct{}

func NewLog(attrGenerator LogAttribute) *Log {
	return &Log{}
}
