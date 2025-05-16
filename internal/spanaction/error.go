package spanaction

type ErrorAttribute interface{}

type Error struct{}

func NewError(attrGenerator ErrorAttribute) *Error {
	return &Error{}
}
