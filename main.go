package main

import (
	"image/color"
	"log"

	"github.com/bailey4770/gomazing/generators/dfs"
	"github.com/bailey4770/gomazing/generators/prims"
	"github.com/bailey4770/gomazing/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

type (
	Tile = utils.Tile
	Grid = utils.Grid
)

type Config struct {
	tileSize      int
	wallThickness int
	maxRows       int
	maxCols       int
	wallImg       *ebiten.Image
	speed         int
}

type Generator interface {
	Iterate(Grid)
	IsComplete() bool
}

func getGenerators(grid Grid) map[string]Generator {
	return map[string]Generator{
		"prims": prims.Initialise(grid),
		"dfs":   dfs.Initialise(grid),
	}
}

type game struct {
	cfg       Config
	grid      Grid
	generator Generator
}

func initGrid(cfg Config) Grid {
	// allocate row slices
	grid := make(Grid, cfg.maxRows)
	cfg.wallImg.Fill(color.White)

	for row := range grid {
		grid[row] = make([]*Tile, cfg.maxCols)
		posY := float64(row * cfg.tileSize)

		for col := range grid[row] {
			posX := float64(col * cfg.tileSize)
			grid[row][col] = utils.CreateTile(posX, posY, row, col)
		}
	}

	return grid
}

func (g *game) Update() error {
	for range g.cfg.speed {
		if !g.generator.IsComplete() {
			g.generator.Iterate(g.grid)
		}
	}

	return nil
}

func main() {
	// Set up ebiten game
	const windowWidth = 640
	const windowHeight = 480
	const tileSize = 20
	const gameSpeed = 3
	const wallThickness = 1

	cfg := Config{
		tileSize:      tileSize,
		wallThickness: wallThickness,
		maxRows:       windowHeight / tileSize,
		maxCols:       windowWidth / tileSize,
		wallImg:       ebiten.NewImage(1, 1),
		speed:         gameSpeed,
	}
	grid := initGrid(cfg)
	generator := "dfs"
	game := &game{
		cfg:       cfg,
		grid:      grid,
		generator: getGenerators(grid)[generator],
	}

	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("Maze generation")

	// Start game loop
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
