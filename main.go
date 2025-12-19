package main

import (
	"image/color"
	"log"
	"math/rand"

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
	windowWidth  int
	windowHeight int
	tileSize     int
	maxRows      int
	maxCols      int
	wallImg      *ebiten.Image
	speed        int
}

type game struct {
	initialised bool
	prims       bool
	dfs         bool
	cfg         Config
	grid        Grid
	frontier    map[*Tile]struct{}
	stack       []*Tile
	visited     map[*Tile]struct{}
	currentTile *Tile
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

func (g *game) resetGrid() {
	for row := range g.grid {
		for col := range g.grid[row] {
			g.grid[row][col].WallN = true
			g.grid[row][col].WallE = true
			g.grid[row][col].WallS = true
			g.grid[row][col].WallW = true
		}
	}
}

func (g *game) Update() error {
	if g.prims {
		if !g.initialised {
			randomRow := rand.Intn(len(g.grid))
			start := utils.GetRandomTile(g.grid[randomRow])

			g.visited[start] = struct{}{}
			neighbours := utils.FindNeighbours(start, g.grid, g.cfg.maxRows, g.cfg.maxCols)

			for _, n := range neighbours {
				g.frontier[n] = struct{}{}
			}

			g.initialised = true
		}

		for range g.cfg.speed {
			if len(g.frontier) > 0 {
				g.frontier, g.visited = prims.Iterate(g.frontier, g.visited, g.grid, g.cfg.maxRows, g.cfg.maxCols)
			}
		}

		if len(g.frontier) == 0 {
			g.resetGrid()
			g.prims = false
			g.dfs = true
			g.initialised = false
		}

	}

	if g.dfs {
		if !g.initialised {
			randomRow := rand.Intn(len(g.grid))
			start := utils.GetRandomTile(g.grid[randomRow])

			g.currentTile = start
			g.visited = make(map[*Tile]struct{})
			g.initialised = true
		}

		for range g.cfg.speed {
			if len(g.visited) < len(g.grid)*len(g.grid[0]) {
				g.stack, g.visited, g.currentTile = dfs.Iterate(g.stack, g.visited, g.currentTile, g.grid, g.cfg.maxRows, g.cfg.maxCols)
			}
		}
	}

	return nil
}

func main() {
	const windowWidth = 640
	const windowHeight = 480
	const tileSize = 20
	const gameSpeed = 3

	cfg := Config{
		windowWidth:  windowWidth,
		windowHeight: windowHeight,
		tileSize:     tileSize,
		maxRows:      windowHeight / tileSize,
		maxCols:      windowWidth / tileSize,
		wallImg:      ebiten.NewImage(1, 1),
		speed:        gameSpeed,
	}
	// TODO since ebiten funcs are methods on game struct, create different game dependent on cli input.
	// game stuct will contain name of algo and callback func to algo to be called from within Update method
	game := &game{
		initialised: false,
		prims:       true,
		dfs:         false,
		cfg:         cfg,
		grid:        nil,
		frontier:    make(map[*Tile]struct{}),
		visited:     make(map[*Tile]struct{}),
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
