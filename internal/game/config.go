package game

type config struct {
	Width  int
	Height int
	// pixels per second
	MaxMoveSpeed float64
}

var Config *config = &config{
	Width:  448,
	Height: 496,
	// https://gaming.stackexchange.com/questions/93692/how-fast-does-pac-man-move
	MaxMoveSpeed: 80,
}
