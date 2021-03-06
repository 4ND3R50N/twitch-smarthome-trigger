package light

type Colors int

const (
	Blue Colors = iota
	Orange
)

func (colors Colors) String() string {
	return [...]string{"Blue", "Orange"}[colors]
}
