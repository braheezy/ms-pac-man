package assets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadSprite(t *testing.T) {
	sprite := loadImage("100pts", spriteFS)

	// Ensure that the returned sprite is not nil
	assert.NotNil(t, sprite)

	// Ensure that the sprite's width and height are valid
	assert.Equal(t, SpriteSize, sprite.Bounds().Dx())
	assert.Equal(t, SpriteSize, sprite.Bounds().Dy())
}

func TestLoadTile(t *testing.T) {
	sprite := loadImage("Level 1/tiles/pellet", levelFS)

	// Ensure that the returned sprite is not nil
	assert.NotNil(t, sprite)

	// Ensure that the sprite's width and height are valid
	assert.Equal(t, TileSize, sprite.Bounds().Dx())
	assert.Equal(t, TileSize, sprite.Bounds().Dy())
}

func TestLoadLevelLayout(t *testing.T) {

	layoutText, err := loadLevelLayout("Level 1")

	// Ensure that the returned sprite is not nil
	assert.NoError(t, err)
	assert.NotNil(t, layoutText)

	assert.Equal(t, layoutText[0][0], '⌜')
	for _, row := range layoutText {
		assert.Equal(t, 14, len(row))
	}
}

func TestCreateLevelImage(t *testing.T) {

	level, tiles, err := LoadLevelImage("Level 1")

	assert.NoError(t, err)
	assert.NotNil(t, level)
	assert.Len(t, tiles[0], 28)
	assert.Len(t, tiles, 31)
	assert.Equal(t, TileTypeWall, tiles[0][0].Type, "First tile should be a wall")
	assert.Equal(t, TileTypeWall, tiles[0][27].Type, "Last tile should be a wall")
	assert.Equal(t, TileTypePellet, tiles[1][1].Type, "Confirm pellet tile")
	assert.Equal(t, TileTypeWall, tiles[30][27].Type, "Confirm random wall")
	assert.Equal(t, TileTypeBlank, tiles[6][1].Type, "Confirm random blank")
	assert.Equal(t, TileTypeGate, tiles[12][14].Type, "Gate is a gate")
	assert.Equal(t, TileTypeWall, tiles[16][12].Type, "Gate wall is a wall")

	for row := range tiles {
		for col, tile := range tiles[row] {
			assert.NotNil(t, tile, "tile at (%d, %d) should not be nil", row, col)
			assert.NotNil(t, tile.Type, "tile.Type at (%d, %d) should not be nil", row, col)
			assert.NotNil(t, tile.Sprite, "tile.Sprite at (%d, %d) should not be nil", row, col)
		}
	}
}
