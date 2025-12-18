package main

import (
	"log"
	"math/rand"

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
	initialised bool
	cfg         *config
	grid        Grid
	wallImg     *ebiten.Image
	frontier    map[*Tile]struct{}
}

func (g *game) getRandomTile() *Tile {
	// choose random tile from frontier list
	randomIndex := rand.Intn(len(g.frontier))
	i := 0

	for tile := range g.frontier {
		if i == randomIndex {
			return tile
		}
		i++
	}

	return nil
}

type Dir struct {
	RowOffset int
	ColOffset int
}

func (g *game) findNeighbours(row, col int) []*Tile {
	dirs := []Dir{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}

	var neighbours []*Tile

	for _, dir := range dirs {
		newRow := row + dir.RowOffset
		newCol := col + dir.ColOffset

		if (newRow >= 0 && newRow < g.cfg.maxRows) && (newCol >= 0 && newCol < g.cfg.maxCols) {
			neighbours = append(neighbours, &g.grid[newRow][newCol])
		}
	}

	return neighbours
}

func (g *game) primsRecursive() {
	frontierTile := g.getRandomTile()
	delete(g.frontier, frontierTile)

	neighbours := g.findNeighbours(frontierTile.row, frontierTile.col)
	var visitedNeighbours []*Tile
	for _, n := range neighbours {
		if n.visited {
			visitedNeighbours = append(visitedNeighbours, n)
		} else if _, ok := g.frontier[n]; !ok {
			g.frontier[n] = struct{}{}
		}
	}

	if len(visitedNeighbours) == 0 {
		return
	}
	// choose random tile from visited neighbours
	randomIndex := rand.Intn(len(visitedNeighbours))
	visitedTile := visitedNeighbours[randomIndex]

	if visitedTile.col == frontierTile.col {
		if frontierTile.row < visitedTile.row {
			// visited is to the south
			frontierTile.wallS = false
			visitedTile.wallN = false
		} else if visitedTile.row < frontierTile.row {
			// visited to the north
			frontierTile.wallN = false
			visitedTile.wallS = false
		} else {
			log.Fatal(visitedTile, frontierTile)
		}
	} else if visitedTile.row == frontierTile.row {
		if frontierTile.col < visitedTile.col {
			// visited to west
			frontierTile.wallE = false
			visitedTile.wallW = false
		} else if visitedTile.col < frontierTile.col {
			// visited to east
			frontierTile.wallW = false
			visitedTile.wallE = false
		} else {
			log.Fatal(visitedTile, frontierTile)
		}
	} else {
		log.Fatal(visitedTile, frontierTile)
	}

	frontierTile.visited = true
}

func (g *game) Update() error {
	if !g.initialised {
		row := 0
		col := 0

		start := &g.grid[row][col]
		start.visited = true
		neighbours := g.findNeighbours(row, col)

		for _, n := range neighbours {
			g.frontier[n] = struct{}{}
		}

		g.initialised = true
	}

	if len(g.frontier) > 0 {
		g.primsRecursive()
	}

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
		false,
		&cfg,
		nil,
		nil,
		make(map[*Tile]struct{}),
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
