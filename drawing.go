package main

import "github.com/hajimehoshi/ebiten/v2"

func (g *game) drawTileWalls(screen *ebiten.Image, t *Tile) {
	tileSize := g.cfg.tileSize
	wallThickness := 1

	// tile.posX and tile.posY are already pixel coordinates
	x := t.posX
	y := t.posY

	// NORTH wall
	if t.wallN {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(float64(tileSize), float64(wallThickness))
		op.GeoM.Translate(x, y)
		screen.DrawImage(g.wallImg, op)
	}

	// SOUTH wall
	if t.wallS {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(float64(tileSize), float64(wallThickness))
		op.GeoM.Translate(x, y+float64(tileSize-wallThickness))
		screen.DrawImage(g.wallImg, op)
	}

	// WEST wall
	if t.wallW {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(float64(wallThickness), float64(tileSize))
		op.GeoM.Translate(x, y)
		screen.DrawImage(g.wallImg, op)
	}

	// EAST wall
	if t.wallE {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(float64(wallThickness), float64(tileSize))
		op.GeoM.Translate(x+float64(tileSize-wallThickness), y)
		screen.DrawImage(g.wallImg, op)
	}
}

func (g *game) Draw(screen *ebiten.Image) {
	for row := 0; row < g.cfg.maxRows; row++ {
		for col := 0; col < g.cfg.maxCols; col++ {
			tile := &g.grid[row][col]

			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(tile.posX, tile.posY)

			g.drawTileWalls(screen, tile)
		}
	}
}
