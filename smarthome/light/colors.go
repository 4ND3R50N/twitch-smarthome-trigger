package light

type Color string

const (
	Blue   Color = "blue"
	Orange       = "orange"
)

func (color Color) String() string {
	return string(color)
}
