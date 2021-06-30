package decorator

// ColorSquare 有颜色的正方形
type ColorSquare struct {
	square IDraw
	color  string
}

func NewColorSquare(color string, square IDraw) *ColorSquare {
	return &ColorSquare{color: color, square: square}
}

// Draw Draw
func (c ColorSquare) Draw() string {
	return c.square.Draw() + ", color is " + c.color
}

// IDraw IDraw
type IDraw interface {
	Draw() string
}

// Square 正方形
type Square struct{}

// Draw Draw
func (s Square) Draw() string {
	return "this is a square"
}
