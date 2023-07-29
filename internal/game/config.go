package game

type config struct {
	Width  int
	Height int
}

var Config *config

func init() {
	Config = &config{
		Width:  448,
		Height: 496,
	}
}
