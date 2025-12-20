// Package dfs runs one iteration of the dfs maze generation algorithm
package dfs

import (
	"math/rand"

	"github.com/bailey4770/gomazing/utils"
)

type (
	Tile = utils.Tile
	Grid = utils.Grid
)

type mazeState struct {
	stack   []*Tile
	visited map[*Tile]struct{}
	curr    *Tile
	maxRows int
	maxCols int
}

func Initialise(grid Grid) *mazeState {
	randomRow := rand.Intn(len(grid))
	start := utils.GetRandomTile(grid[randomRow])

	return &mazeState{
		visited: make(map[*Tile]struct{}),
		curr:    start,
		maxRows: len(grid),
		maxCols: len(grid[0]),
	}
}

func (m *mazeState) Iterate(grid Grid) {
	neighbours := utils.FindNeighbours(m.curr, grid, m.maxRows, m.maxCols)
	var unvisitedNeighbours []*Tile
	for _, n := range neighbours {
		if _, ok := m.visited[n]; !ok {
			unvisitedNeighbours = append(unvisitedNeighbours, n)
		}
	}

	if len(unvisitedNeighbours) > 0 {
		randUnvisited := utils.GetRandomTile(unvisitedNeighbours)
		m.stack = append(m.stack, m.curr)
		utils.RemoveWalls(m.curr, randUnvisited)

		m.visited[randUnvisited] = struct{}{}
		m.curr = randUnvisited
	} else if len(m.stack) > 0 {
		m.curr = m.stack[len(m.stack)-1]
		m.stack = m.stack[:len(m.stack)-1]
	}
}

func (m *mazeState) IsComplete() bool {
	return len(m.visited) >= m.maxRows*m.maxCols
}
