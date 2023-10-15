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
	previousWaypoint, currentWaypoint, nextWaypoint WaypointPos
	// How fast the player moves in pixels per frame
	moveSpeed                     float64
	waypointHeight, waypointWidth int
	turning                       bool

	grid [][]assets.Tile
}

func (p *Player) processWaypoints() {
	if p.currentDirection != p.requestedDirection {
		// recalculate waypoints on direction change request
		if (p.currentDirection == Up && p.requestedDirection == Down) ||
			(p.currentDirection == Down && p.requestedDirection == Up) ||
			(p.currentDirection == Left && p.requestedDirection == Right) ||
			(p.currentDirection == Right && p.requestedDirection == Left) {
			p.currentWaypoint = p.previousWaypoint
			p.nextWaypoint = p.currentWaypoint
			p.currentDirection = p.requestedDirection
		}
	}
	// Handle progress towards current waypoint
	waypointDistance := math.Sqrt(math.Pow(p.currentWaypoint.Center().X-p.currentPixelPos.X, 2) + math.Pow(p.currentWaypoint.Center().Y-p.currentPixelPos.Y, 2))

	waypointIsCaptured := waypointDistance <= 1

	if waypointIsCaptured {
		p.previousWaypoint = p.currentWaypoint
		p.currentWaypoint = p.nextWaypoint
		p.setNextWaypoint()

		if p.currentWaypoint == p.nextWaypoint {
			return
		}
	}

	p.updateTurnStatus()

	// continuing moving to waypoint
	if p.turning {
		// TODO:
		//  - Handle curved waypoint paths i.e. turns

		// p.currentPixelPos = p.currentWaypoint.Center()
	} else {
		if waypointDistance == 0 {
			return
		}
		// Calculate the vector from the player to the waypoint
		directionVector := PixelPos{p.currentWaypoint.Center().X - p.currentPixelPos.X, p.currentWaypoint.Center().Y - p.currentPixelPos.Y}

		// Calculate the magnitude of the direction vector
		magnitude := math.Sqrt(directionVector.X*directionVector.X + directionVector.Y*directionVector.Y)

		// Normalize the direction vector (make it a unit vector)
		directionVector.X /= magnitude
		directionVector.Y /= magnitude

		// Calculate the new player position
		p.currentPixelPos = PixelPos{
			X: p.currentPixelPos.X + directionVector.X*p.moveSpeed,
			Y: p.currentPixelPos.Y + directionVector.Y*p.moveSpeed,
		}
	}

	// p.currentPixelPos = p.currentWaypoint.Center()
}

func (p *Player) setNextWaypoint() {
	nextX, nextY := p.currentWaypoint.X, p.currentWaypoint.Y
	newDirection := p.requestedDirection
	switch newDirection {
	case Up:
		nextY = nextY - 1
	case Right:
		nextX = nextX + 1
	case Down:
		nextY = nextY + 1
	case Left:
		nextX = nextX - 1
	}

	if p.grid[nextY][nextX].Type == assets.TileTypeWall {
		nextX, nextY = p.currentWaypoint.X, p.currentWaypoint.Y
		newDirection = p.currentDirection
		switch newDirection {
		case Up:
			nextY = nextY - 1
		case Right:
			nextX = nextX + 1
		case Down:
			nextY = nextY + 1
		case Left:
			nextX = nextX - 1
		}

		if p.grid[nextY][nextX].Type == assets.TileTypeWall {
			return
		}
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

	p.currentDirection = newDirection
	p.nextWaypoint = WaypointPos{nextX, nextY}
}

func (p *Player) Update() {
	p.updateRequestedDirection()

	p.processWaypoints()
}

func (p *Player) updateTurnStatus() {
	// Calculate vectors from the player to the current waypoint and from the current waypoint to the next waypoint
	vector1 := PixelPos{p.currentWaypoint.Center().X - p.currentPixelPos.X, p.currentWaypoint.Center().Y - p.currentPixelPos.Y}
	vector2 := PixelPos{p.nextWaypoint.Center().X - p.currentWaypoint.Center().X, p.nextWaypoint.Center().Y - p.currentWaypoint.Center().Y}

	// Calculate the angle between the two vectors
	dotProduct := vector1.X*vector2.X + vector1.Y*vector2.Y
	magnitude1 := math.Sqrt(vector1.X*vector1.X + vector1.Y*vector1.Y)
	magnitude2 := math.Sqrt(vector2.X*vector2.X + vector2.Y*vector2.Y)

	// Calculate the cosine of the angle
	cosTheta := dotProduct / (magnitude1 * magnitude2)

	// Calculate the angle in radians
	theta := math.Acos(cosTheta)

	// Convert the angle to degrees
	angleDegrees := theta * (180.0 / math.Pi)

	// Check if the angle is not approximately 180 degrees
	p.turning = angleDegrees != 0 && math.Abs(angleDegrees-180.0) > epsilon
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
