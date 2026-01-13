package main

import (
	"fmt"
	"log"

	"github.com/bailey4770/gomazing/cli"
	"github.com/bailey4770/gomazing/mazesave"
	"github.com/bailey4770/gomazing/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

type (
	Tile      = utils.Tile
	Grid      = utils.Grid
	Config    = cli.Config
	Generator = cli.Generator
)

type game struct {
	cfg       Config
	grid      Grid
	generator Generator
	complete  bool
}

func initGrid(cfg Config) Grid {
	// allocate row slices
	grid := make(Grid, cfg.MaxRows)

	for row := range grid {
		grid[row] = make([]*Tile, cfg.MaxCols)
		posY := float64(row * cfg.TileSize)

		for col := range grid[row] {
			posX := float64(col * cfg.TileSize)
			grid[row][col] = utils.CreateTile(posX, posY, row, col)
		}
	}

	return grid
}

func (g *game) Update() error {
	if g.generator == nil {
		return nil
	}

	for range g.cfg.Speed {
		if !g.generator.IsComplete() {
			err := g.generator.Iterate(g.grid)
			if err != nil {
				return err
			}
		} else if !g.complete {
			fmt.Println("maze complete")
			g.complete = true
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		if !g.generator.IsComplete() {
			log.Print("Error: wait until the maze has finished generating.")
		} else {
			log.Print("Saving maze...")

			fileName := "test.maze"

			err := mazesave.SaveMaze(g.grid, g.cfg.TileSize, fileName)
			if err != nil {
				return fmt.Errorf("could not save maze: %v", err)
			}
		}
	}

	return nil
}

func main() {
	// Set up ebiten game
	cfg := cli.GetConfig()
	grid := initGrid(cfg)
	game := &game{
		cfg:       cfg,
		grid:      grid,
		generator: cfg.Generator,
		complete:  false,
	}

	if game.generator != nil {
		// start generation game
		if err := game.generator.Initialise(grid); err != nil {
			log.Fatalf("Error: %v", err)
		}
	} else {
		err := mazesave.LoadMazeWalls(game.cfg.MazePath, game.grid)
		if err != nil {
			log.Fatal("Error:", err)
		}
	}

	ebiten.SetWindowSize(cfg.WindowWidth, cfg.WindowHeight)
	ebiten.SetWindowTitle("Gomazing")
	// Start game loop
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal("Error:", err)
	}
}
