package light

type Color string

const (
	Blue    Color = "blue"
	Orange  Color = "orange"
	Unknown Color = "unknown"
)

func (color Color) ByString(stringColor string) Color {
	if Blue.String() == stringColor {
		return Blue
	}
	if Orange.String() == stringColor {
		return Orange
	}

	return Unknown
}

func (color Color) String() string {
	return string(color)
}
