package game

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

func (d Direction) String() string {
	names := [...]string{
		"Up",
		"Down",
		"Left",
		"Right",
	}

	if d < Up || d > Right {
		return "UnknownDirection"
	}

	return names[d]
}
