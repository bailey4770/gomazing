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

type game struct {
	cfg     *config
	grid    Grid
	wallImg *ebiten.Image
}

func (g *game) initGrid() {
	// allocate row slices
	g.grid = make(Grid, g.cfg.maxRows)

	for row := range g.grid {
		g.grid[row] = make([]Tile, g.cfg.maxCols)
		posY := float64(row * g.cfg.tileSize)

		for col := range g.grid[row] {
			posX := float64(col * g.cfg.tileSize)
			g.grid[row][col] = createTile(g, posX, posY)
		}

		g.wallImg = ebiten.NewImage(1, 1)
		g.wallImg.Fill(color.White)
	}
}

func (g *game) Update() error {
	return nil
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
