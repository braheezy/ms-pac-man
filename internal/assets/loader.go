package assets

import (
	"embed"
	"io/fs"
	"log"
	"path"
	"strings"

	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//go:embed sprites/*
var spriteFS embed.FS

//go:embed levels/*
var levelFS embed.FS

func LoadSprite(name string) *ebiten.Image {
	return loadImage(name, spriteFS)
}

func loadImage(path string, filesystem fs.FS) (img *ebiten.Image) {
	var err error
	switch filesystem {
	case levelFS:
		img, _, err = ebitenutil.NewImageFromFileSystem(filesystem, "levels/"+path+".png")
	case spriteFS:
		img, _, err = ebitenutil.NewImageFromFileSystem(filesystem, "sprites/"+path+".png")
	}

	if err != nil {
		log.Fatalln(err)
	}
	return img
}

var tileLookup = map[rune]string{
	'⌜': "W_ULcorner_thin",
	'⌞': "W_DLcorner_thin",
	'-': "W_U",
	'_': "W_D",
	'⎡': "W_ULcorner_thick",
	'⎤': "W_URcorner_thick",
	'⎸': "W_L",
	'⎹': "W_R",
	'•': "pellet",
	']': "W_Rthick",
	'[': "W_Lthick",
	'⦾': "power_pellet",
	'⌌': "W_ULcorner_nub",
	'⌍': "W_URcorner_nub",
	'⌎': "W_DLcorner_nub",
	'⌏': "W_DRcorner_nub",
	'⎽': "W_Dthick",
	'⎺': "W_Uthick",
	'⎩': "W_DLcorner_inner",
	'⎭': "W_DRcorner_inner",
	'⎧': "W_ULcorner_inner",
	' ': "blank",
	'▓': "W_fill",
	'「': "W_ULcorner_pen",
	'﹂': "W_DLcorner_pen",
	'>': "W_L_pengate",
	'*': "gate",
}

const (
	TileSize   = 16
	SpriteSize = 32
)

type TileType int

const (
	TileTypeBlank TileType = iota
	TileTypeWall
	TileTypeGate
	TileTypePellet
	TileTypePowerPellet
	TileTypePlayer
)

func (t TileType) String() string {
	names := [...]string{
		"TileTypeBlank",
		"TileTypeWall",
		"TileTypeGate",
		"TileTypePellet",
		"TileTypePowerPellet",
		"TileTypePlayer",
	}

	if t < TileTypeBlank || t > TileTypePlayer {
		return "UnknownTileType"
	}

	return names[t]
}

type Tile struct {
	Type   TileType
	Sprite *ebiten.Image
}

func (t Tile) String() string {
	return tileLookup[rune(t.Type)]
}

func LoadLevelImage(levelName string) (*ebiten.Image, [][]Tile, error) {
	layout, err := loadLevelLayout(levelName)
	if err != nil {
		return nil, nil, err
	}

	levelHeight := len(layout)
	levelWidth := len(layout[0])

	fullImage := ebiten.NewImage(levelWidth*TileSize*2, levelHeight*TileSize)
	tiles := make([][]Tile, levelHeight)
	for i := range tiles {
		tiles[i] = make([]Tile, levelWidth*2)
	}

	for rowIdx, row := range layout {
		for colIdx, char := range row {
			tileName, ok := tileLookup[char]
			if !ok {
				continue
			}

			tileSprite := loadImage(path.Join(levelName, "tiles", tileName), levelFS)
			if err != nil {
				return nil, nil, err
			}

			// Calculate the position to draw the tile
			tileX := colIdx * TileSize
			tileY := rowIdx * TileSize
			op := &ebiten.DrawImageOptions{}

			var tileType TileType
			switch tileName {
			case "blank":
				tileType = TileTypeBlank
			case "gate":
				tileType = TileTypeGate
			case "pellet":
				tileType = TileTypePellet
			case "power_pellet":
				tileType = TileTypePowerPellet
			default:
				tileType = TileTypeWall
			}

			// Draw the tile on the full image
			op.GeoM.Translate(float64(tileX), float64(tileY))
			fullImage.DrawImage(tileSprite, op)

			tiles[rowIdx][colIdx] = Tile{Type: tileType, Sprite: tileSprite}

			// Create the mirrored sprite
			mirroredTileSprite := ebiten.NewImage(TileSize, TileSize)

			// Calculate the position to draw the mirrored tile (along the Y axis)
			mirroredX := (levelWidth - colIdx) * TileSize
			mirroredY := rowIdx * TileSize

			// Draw the mirrored tile on the full image
			mirroredOp := &ebiten.DrawImageOptions{}
			// Mirroring along the Y axis
			mirroredOp.GeoM.Scale(-1, 1)
			// First draw the mirrored sprite
			mirroredTileSprite.DrawImage(tileSprite, mirroredOp)
			// Then draw it on the level image
			mirroredOp.GeoM.Translate(float64(mirroredX+levelWidth*TileSize), float64(mirroredY))
			fullImage.DrawImage(tileSprite, mirroredOp)

			tiles[rowIdx][levelWidth*2-colIdx-1] = Tile{Type: tileType, Sprite: mirroredTileSprite}
		}
	}

	return fullImage, tiles, nil
}
func loadLevelLayout(levelName string) ([][]rune, error) {
	// Read layout from file system
	layoutText, err := fs.ReadFile(levelFS, path.Join("levels", levelName, "layout.txt"))
	if err != nil {
		return nil, err
	}

	// Parse layout into 2D array of characters
	layoutLines := strings.Split(string(layoutText), "\n")
	layout := make([][]rune, len(layoutLines))
	for i, line := range layoutLines {
		layout[i] = []rune(line)
	}

	return layout, nil
}
