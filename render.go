package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func drawTileWalls(screen *ebiten.Image, cfg Config, t *Tile) {
	tileSize := cfg.TileSize
	wallThickness := cfg.WallThickness

	// tile.posX and tile.posY are already pixel coordinates
	x, y := t.PosX, t.PosY

	// NORTH wall
	if t.WallN {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(float64(tileSize), float64(wallThickness))
		op.GeoM.Translate(x, y)
		screen.DrawImage(cfg.WallImg, op)
	}

	// SOUTH wall
	if t.WallS {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(float64(tileSize), float64(wallThickness))
		op.GeoM.Translate(x, y+float64(tileSize-wallThickness))
		screen.DrawImage(cfg.WallImg, op)
	}

	// WEST wall
	if t.WallW {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(float64(wallThickness), float64(tileSize))
		op.GeoM.Translate(x, y)
		screen.DrawImage(cfg.WallImg, op)
	}

	// EAST wall
	if t.WallE {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(float64(wallThickness), float64(tileSize))
		op.GeoM.Translate(x+float64(tileSize-wallThickness), y)
		screen.DrawImage(cfg.WallImg, op)
	}
}

func (g *game) Draw(screen *ebiten.Image) {
	for row := 0; row < g.cfg.MaxRows; row++ {
		for col := 0; col < g.cfg.MaxCols; col++ {
			tile := g.grid[row][col]

			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(tile.PosX, tile.PosY)

			drawTileWalls(screen, g.cfg, tile)
		}
	}

	if g.cfg.ShowStats {
		// Display FPS and TPS
		fps := ebiten.ActualFPS()
		tps := ebiten.ActualTPS()
		msg := fmt.Sprintf("FPS: %.2f\nTPS: %.2f",
			fps, tps)
		ebitenutil.DebugPrintAt(screen, msg, 1, 1)
	}
}

func (g *game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
