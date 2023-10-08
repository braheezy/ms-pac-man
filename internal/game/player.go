package game

import (
	"math"

	"github.com/braheezy/ms-pacman/internal/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	// The current sprite to show
	image *ebiten.Image
	// The current movement direction
	currentDirection Direction
	// The input the user has requested
	requestedDirection Direction
	// The pixel location of the player, based on center of the sprite
	currentPixelPos PixelPos
	// Where on the grid the play is occupying
	currentWaypoint, nextWaypoint WaypointPos
	// How fast the player moves in pixels per frame
	moveSpeed                     float64
	waypointHeight, waypointWidth int

	grid [][]assets.Tile
}

func (p *Player) processWaypoints() {
	waypointDistance := math.Sqrt(math.Pow(p.currentWaypoint.Center().X-p.currentPixelPos.X, 2) + math.Pow(p.currentWaypoint.Center().Y-p.currentPixelPos.Y, 2))

	waypointIsCaptured := waypointDistance <= 1
	if waypointIsCaptured {
		p.currentWaypoint = p.nextWaypoint
		p.setNextWaypoint()
	} else {
		// continuing moving to waypoint
		p.currentPixelPos = p.currentWaypoint.Center()
	}
}

func (p *Player) setNextWaypoint() {
	nextX, nextY := p.currentWaypoint.X, p.currentWaypoint.Y
	switch p.requestedDirection {
	case Up:
		nextY = p.currentWaypoint.Y - 1
	case Right:
		nextX = p.currentWaypoint.X + 1
	case Down:
		nextY = p.currentWaypoint.Y + 1
	case Left:
		nextX = p.currentWaypoint.X - 1
	}

	if p.grid[nextY][nextX].Type == assets.TileTypeWall {
		return
	}

	if nextX < 1 {
		nextX = 1
	} else if nextX >= p.waypointWidth {
		nextX = p.waypointWidth
	}

	if nextY < 1 {
		nextY = 1
	} else if nextY >= p.waypointHeight {
		nextY = p.waypointHeight
	}

	p.nextWaypoint = WaypointPos{nextX, nextY}

}

func (p *Player) Update() {
	p.updateRequestedDirection()

	p.processWaypoints()
}

func (p *Player) updateRequestedDirection() {
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		p.requestedDirection = Right
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		p.requestedDirection = Left
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		p.requestedDirection = Up
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		p.requestedDirection = Down
	}
}
