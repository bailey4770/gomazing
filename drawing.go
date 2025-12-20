package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func drawTileWalls(screen *ebiten.Image, cfg Config, t *Tile) {
	tileSize := cfg.tileSize
	wallThickness := cfg.wallThickness

	// tile.posX and tile.posY are already pixel coordinates
	x, y := t.PosX, t.PosY

	// NORTH wall
	if t.WallN {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(float64(tileSize), float64(wallThickness))
		op.GeoM.Translate(x, y)
		screen.DrawImage(cfg.wallImg, op)
	}

	// SOUTH wall
	if t.WallS {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(float64(tileSize), float64(wallThickness))
		op.GeoM.Translate(x, y+float64(tileSize-wallThickness))
		screen.DrawImage(cfg.wallImg, op)
	}

	// WEST wall
	if t.WallW {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(float64(wallThickness), float64(tileSize))
		op.GeoM.Translate(x, y)
		screen.DrawImage(cfg.wallImg, op)
	}

	// EAST wall
	if t.WallE {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(float64(wallThickness), float64(tileSize))
		op.GeoM.Translate(x+float64(tileSize-wallThickness), y)
		screen.DrawImage(cfg.wallImg, op)
	}
}

func (g *game) Draw(screen *ebiten.Image) {
	for row := 0; row < g.cfg.maxRows; row++ {
		for col := 0; col < g.cfg.maxCols; col++ {
			tile := g.grid[row][col]

			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(tile.PosX, tile.PosY)

			drawTileWalls(screen, g.cfg, tile)
		}
	}

	// Display FPS and TPS
	fps := ebiten.ActualFPS()
	tps := ebiten.ActualTPS()
	msg := fmt.Sprintf("FPS: %.2f\nTPS: %.2f",
		fps, tps)
	ebitenutil.DebugPrintAt(screen, msg, 1, 1)
}

func (g *game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.cfg.windowWidth, g.cfg.windowHeight
}
