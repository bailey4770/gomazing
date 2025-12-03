package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type config struct {
	windowWidth  int
	windowHeight int
	tileSize     int
	maxRows      int
	maxCols      int
}

type Tile struct {
	img  *ebiten.Image
	posX float64
	posY float64
}

type (
	Grid [][]Tile
)

func createTile(g *game, posX, posY float64, alt bool) Tile {
	var tile Tile
	tile.img = ebiten.NewImage(g.cfg.tileSize, g.cfg.tileSize)

	if !alt {
		tile.img.Fill(color.RGBA{0, 0, 0, 255})
	} else {
		tile.img.Fill(color.RGBA{255, 255, 255, 255})
	}

	tile.posX = posX
	tile.posY = posY
	return tile
}

type game struct {
	cfg  *config
	grid Grid
}

func (g *game) initGrid() {
	// allocate row slices
	g.grid = make(Grid, g.cfg.maxRows)

	for row := range g.grid {
		g.grid[row] = make([]Tile, g.cfg.maxCols)
		posY := float64(row * g.cfg.tileSize)

		altCol := (row % 2) == 0

		for col := range g.grid[row] {
			posX := float64(col * g.cfg.tileSize)
			g.grid[row][col] = createTile(g, posX, posY, altCol)

			altCol = !altCol
		}
	}
}

func (g *game) Update() error {
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	for row := 0; row < g.cfg.maxRows; row++ {
		for col := 0; col < g.cfg.maxCols; col++ {
			tile := g.grid[row][col]

			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(tile.posX, tile.posY)

			screen.DrawImage(tile.img, op)
		}
	}
}

func (g *game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.cfg.windowWidth, g.cfg.windowHeight
}

func main() {
	const windowWidth = 640
	const windowHeight = 480
	const tileSize = 20

	cfg := config{
		windowWidth,
		windowHeight,
		tileSize,
		windowHeight / tileSize,
		windowWidth / tileSize,
	}
	game := &game{
		&cfg,
		nil,
	}
	// create grid of tiles
	game.initGrid()

	ebiten.SetWindowSize(cfg.windowWidth, cfg.windowHeight)
	ebiten.SetWindowTitle("Maze generation")

	// Start game loop
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
