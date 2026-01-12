// Package dfs runs one iteration of the dfs maze generation algorithm. Add GetMazeState() func to cli to include in program
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

func GetMazeState() *mazeState {
	return &mazeState{
		visited: make(map[*Tile]struct{}),
	}
}

func (m *mazeState) Initialise(grid Grid) error {
	randomRow := rand.Intn(len(grid))
	start, _, err := utils.GetRandomTile(grid[randomRow])
	if err != nil {
		return err
	}

	m.curr = start
	m.maxRows = len(grid)
	m.maxCols = len(grid[0])

	return nil
}

func (m *mazeState) Iterate(grid Grid) error {
	neighbours := utils.FindNeighbours(m.curr, grid, m.maxRows, m.maxCols)
	var unvisitedNeighbours []*Tile
	for _, n := range neighbours {
		if _, ok := m.visited[n]; !ok {
			unvisitedNeighbours = append(unvisitedNeighbours, n)
		}
	}

	if len(unvisitedNeighbours) > 0 {
		randUnvisited, _, err := utils.GetRandomTile(unvisitedNeighbours)
		if err != nil {
			return err
		}

		m.stack = append(m.stack, m.curr)
		utils.RemoveWalls(m.curr, randUnvisited)

		m.visited[randUnvisited] = struct{}{}
		m.curr = randUnvisited
	} else if len(m.stack) > 0 {
		m.curr = m.stack[len(m.stack)-1]
		m.stack = m.stack[:len(m.stack)-1]
	}

	return nil
}

func (m *mazeState) IsComplete() bool {
	return len(m.visited) >= m.maxRows*m.maxCols
}
