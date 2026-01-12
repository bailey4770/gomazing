// Package prims handles one iteration of the prims maze generation algorithm. Add GetMazeState() func to cli to include in program
package prims

import (
	"errors"
	"math/rand"

	"github.com/bailey4770/gomazing/utils"
)

type (
	Tile = utils.Tile
	Grid = utils.Grid
)

type mazeState struct {
	frontier map[*Tile]struct{}
	visited  map[*Tile]struct{}
	maxRows  int
	maxCols  int
}

func GetMazeState() *mazeState {
	return &mazeState{
		frontier: make(map[*Tile]struct{}),
		visited:  make(map[*Tile]struct{}),
	}
}

func (m *mazeState) Initialise(grid Grid) error {
	m.maxRows = len(grid)
	m.maxCols = len(grid[0])

	randomRow := rand.Intn(len(grid))
	start, _, err := utils.GetRandomTile(grid[randomRow])
	if err != nil {
		return err
	}
	m.visited[start] = struct{}{}

	neighbours := utils.FindNeighbours(start, grid, m.maxRows, m.maxCols)
	for _, n := range neighbours {
		m.frontier[n] = struct{}{}
	}

	return nil
}

func (m *mazeState) Iterate(grid Grid) error {
	// getRandomTile takes a slice for efficient random selection.
	// no good way of randomly selecting an element from a map in Go. Frontier remains a map for efficient contains checking and easy deletion from queue.
	var frontierSlice []*Tile
	for t := range m.frontier {
		frontierSlice = append(frontierSlice, t)
	}
	frontierTile, _, err := utils.GetRandomTile(frontierSlice)
	if err != nil {
		return err
	}
	delete(m.frontier, frontierTile)

	neighbours := utils.FindNeighbours(frontierTile, grid, m.maxRows, m.maxCols)
	var visitedNeighbours []*Tile
	for _, n := range neighbours {
		if _, ok := m.visited[n]; ok {
			visitedNeighbours = append(visitedNeighbours, n)
		} else if _, ok := m.frontier[n]; !ok {
			m.frontier[n] = struct{}{}
		}
	}

	if len(visitedNeighbours) == 0 {
		return errors.New("there were no visited neighbours")
	}

	// choose random tile from visited neighbours
	randomIndex := rand.Intn(len(visitedNeighbours))
	visitedTile := visitedNeighbours[randomIndex]

	utils.RemoveWalls(frontierTile, visitedTile)
	m.visited[frontierTile] = struct{}{}

	return nil
}

func (m *mazeState) IsComplete() bool {
	return len(m.frontier) <= 0
}
