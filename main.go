package main

import (
	"image/color"
	"log"
	"math/rand"

	"github.com/bailey4770/gomazing/generators/prims"
	"github.com/bailey4770/gomazing/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

type (
	Tile = utils.Tile
	Grid = utils.Grid
)

type Config struct {
	windowWidth  int
	windowHeight int
	tileSize     int
	maxRows      int
	maxCols      int
	wallImg      *ebiten.Image
}

type game struct {
	initialised bool
	cfg         Config
	grid        Grid
	frontier    map[*Tile]struct{}
}

func (g *game) initGrid() {
	// allocate row slices
	g.grid = make(Grid, g.cfg.maxRows)
	g.cfg.wallImg.Fill(color.White)

	for row := range g.grid {
		g.grid[row] = make([]*Tile, g.cfg.maxCols)
		posY := float64(row * g.cfg.tileSize)

		for col := range g.grid[row] {
			posX := float64(col * g.cfg.tileSize)
			g.grid[row][col] = createTile(g, posX, posY, row, col)
		}
	}
}

func (g *game) Update() error {
	if !g.initialised {
		randomRow := rand.Intn(len(g.grid))
		start := utils.GetRandomTile(g.grid[randomRow])

		start.Visited = true
		neighbours := utils.FindNeighbours(start, g.grid, g.cfg.maxRows, g.cfg.maxCols)

		for _, n := range neighbours {
			g.frontier[n] = struct{}{}
		}

		g.initialised = true
	}

	if len(g.frontier) > 0 {
		g.frontier = prims.Iterate(g.frontier, g.grid, g.cfg.maxRows, g.cfg.maxCols)
	}

	return nil
}

func main() {
	const windowWidth = 640
	const windowHeight = 480
	const tileSize = 20

	cfg := Config{
		windowWidth:  windowWidth,
		windowHeight: windowHeight,
		tileSize:     tileSize,
		maxRows:      windowHeight / tileSize,
		maxCols:      windowWidth / tileSize,
		wallImg:      ebiten.NewImage(1, 1),
	}
	// TODO since ebiten funcs are methods on game struct, create different game dependent on cli input.
	// game stuct will contain name of algo and callback func to algo to be called from within Update method
	game := &game{
		initialised: false,
		cfg:         cfg,
		grid:        nil,
		frontier:    make(map[*Tile]struct{}),
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
