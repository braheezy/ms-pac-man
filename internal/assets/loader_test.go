package assets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadSprite(t *testing.T) {
	sprite, err := LoadImage("100pts", spriteFS)
	if err != nil {
		t.Fatalf("Error loading sprite: %v", err)
	}

	// Ensure that the returned sprite is not nil
	assert.NotNil(t, sprite)

	// Ensure that the sprite's width and height are valid
	assert.Equal(t, SpriteSize, sprite.Bounds().Dx())
	assert.Equal(t, SpriteSize, sprite.Bounds().Dy())
}

func TestLoadTile(t *testing.T) {
	sprite, err := LoadImage("level1/tiles/pellet", levelFS)
	if err != nil {
		t.Fatalf("Error loading sprite: %v", err)
	}

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

	level, err := CreateLevelImage("level1")

	// Ensure that the returned sprite is not nil
	assert.NoError(t, err)
	assert.NotNil(t, level)

}
