package enums

type Color int

const (
	Red Color = iota
	Yellow
	Blue
)

func (c Color) String() string {
	switch c {
	case Red:
		return "red"
	case Yellow:
		return "yellow"
	case Blue:
		return "blue"
	default:
		return "Unknown"
	}
}
