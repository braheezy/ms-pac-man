package assets

import (
	"embed"
	"io/fs"
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

func LoadImage(path string, filesystem fs.FS) (img *ebiten.Image, err error) {
	switch filesystem {
	case levelFS:
		img, _, err = ebitenutil.NewImageFromFileSystem(filesystem, "levels/"+path+".png")
	case spriteFS:
		img, _, err = ebitenutil.NewImageFromFileSystem(filesystem, "sprites/"+path+".png")
	}

	if err != nil {
		return nil, err
	}
	return img, nil
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

const TileSize = 16
const SpriteSize = 32

func CreateLevelImage(levelName string) (*ebiten.Image, error) {
	layout, err := loadLevelLayout(levelName)
	if err != nil {
		return nil, err
	}

	levelHeight := len(layout)
	levelWidth := len(layout[0])

	fullImage := ebiten.NewImage(levelWidth*TileSize*2, levelHeight*TileSize)

	for rowIdx, row := range layout {
		for colIdx, char := range row {
			tileName, ok := tileLookup[char]
			if !ok {
				continue
			}

			tileSprite, err := LoadImage(path.Join(levelName, "tiles", tileName), levelFS)
			if err != nil {
				return nil, err
			}

			// Calculate the position to draw the tile
			tileX := colIdx * TileSize
			tileY := rowIdx * TileSize

			// Draw the tile on the full image
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(tileX), float64(tileY))
			// fmt.Printf("Adding tile %s at (%d,%d)\n", tileName, tileX, tileY)
			fullImage.DrawImage(tileSprite, op)

			// Calculate the position to draw the mirrored tile (along the Y axis)
			mirroredX := (levelWidth - colIdx) * TileSize
			mirroredY := rowIdx * TileSize

			// Draw the mirrored tile on the full image
			mirroredOp := &ebiten.DrawImageOptions{}
			mirroredOp.GeoM.Scale(-1, 1) // Mirroring along the Y axis
			mirroredOp.GeoM.Translate(float64(mirroredX+levelWidth*TileSize), float64(mirroredY))
			// fmt.Printf("Adding tile %s at (%d,%d)\n", tileName, mirroredX+levelWidth*TileSize, mirroredY)
			fullImage.DrawImage(tileSprite, mirroredOp)
		}
	}
	return fullImage, nil
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