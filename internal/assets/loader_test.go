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
	sprite := loadImage("level1/tiles/pellet", levelFS)

	// Ensure that the returned sprite is not nil
	assert.NotNil(t, sprite)

	// Ensure that the sprite's width and height are valid
	assert.Equal(t, TileSize, sprite.Bounds().Dx())
	assert.Equal(t, TileSize, sprite.Bounds().Dy())
}

func TestLoadLevelLayout(t *testing.T) {

	layoutText, err := loadLevelLayout("level1")

	// Ensure that the returned sprite is not nil
	assert.NoError(t, err)
	assert.NotNil(t, layoutText)

	assert.Equal(t, layoutText[0][0], '⌜')
}

func TestCreateLevelImage(t *testing.T) {

	level, tiles, err := LoadLevelImage("level1")

	assert.NoError(t, err)
	assert.NotNil(t, level)
	assert.Len(t, tiles[0], 28)
	assert.Len(t, tiles, 31)
	assert.Equal(t, tiles[0][0], TileTypeWall)
	assert.Equal(t, tiles[0][27], TileTypeWall)
	assert.Equal(t, tiles[1][1], TileTypePellet)
	assert.Equal(t, tiles[30][27], TileTypeWall)
	assert.Equal(t, TileTypePlayer, tiles[13][23])

}
