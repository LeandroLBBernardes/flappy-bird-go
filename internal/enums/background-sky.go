package enums

type BackgroundSky int

const (
	Day BackgroundSky = iota
	Night
)

func (s BackgroundSky) String() string {
	switch s {
	case Day:
		return "day"
	case Night:
		return "night"
	default:
		return "Unknown"
	}
}
