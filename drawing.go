package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Tile struct {
	img     *ebiten.Image
	posX    float64
	posY    float64
	row     int
	col     int
	wallN   bool
	wallE   bool
	wallS   bool
	wallW   bool
	visited bool
}

type (
	Grid [][]Tile
)

func createTile(g *game, posX, posY float64, row, col int) Tile {
	var tile Tile
	tile.img = ebiten.NewImage(g.cfg.tileSize, g.cfg.tileSize)

	tile.img.Fill(color.RGBA{255, 255, 255, 255})

	tile.posX = posX
	tile.posY = posY
	tile.row = row
	tile.col = col
	tile.wallN = true
	tile.wallE = true
	tile.wallS = true
	tile.wallW = true
	tile.visited = false
	return tile
}

func (g *game) initGrid() {
	// allocate row slices
	g.grid = make(Grid, g.cfg.maxRows)

	for row := range g.grid {
		g.grid[row] = make([]Tile, g.cfg.maxCols)
		posY := float64(row * g.cfg.tileSize)

		for col := range g.grid[row] {
			posX := float64(col * g.cfg.tileSize)
			g.grid[row][col] = createTile(g, posX, posY, row, col)
		}

		g.wallImg = ebiten.NewImage(1, 1)
		g.wallImg.Fill(color.White)
	}
}

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
