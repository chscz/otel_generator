package spanaction

type RenderAttribute interface{}
type Render struct{}

func NewRender(attrGenerator RenderAttribute) *Render {
	return &Render{}
}
