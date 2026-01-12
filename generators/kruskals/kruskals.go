// Package kruskals handles one iteration of the kruskals mazze generation algorithm. Add GetMazeState() func to cli to include in program
package kruskals

import (
	"errors"
	"math/rand/v2"

	"github.com/bailey4770/gomazing/utils"
)

type (
	Tile = utils.Tile
	Grid = utils.Grid
)

type wall struct {
	tile1 *Tile
	tile2 *Tile
}

type mazeState struct {
	tileSets       *utils.UnionFind
	unionCount     int
	unionsRequired int
	walls          []wall
	wallIdx        int
}

func GetMazeState() *mazeState {
	return &mazeState{
		tileSets:   utils.NewUnionFind(),
		wallIdx:    0,
		unionCount: 0,
	}
}

func (m *mazeState) Initialise(grid Grid) error {
	for i, row := range grid {
		for j := range row {

			tile := row[j]
			// Add east wall if not on right edge
			if j < len(row)-1 {
				wallE := wall{tile1: tile, tile2: row[j+1]}
				m.walls = append(m.walls, wallE)
			}

			// Add south wall if not on bottom edge
			if i < len(grid)-1 {
				wallS := wall{tile1: tile, tile2: grid[i+1][j]}
				m.walls = append(m.walls, wallS)
			}
		}
	}

	rand.Shuffle(len(m.walls), func(i, j int) {
		m.walls[i], m.walls[j] = m.walls[j], m.walls[i]
	})

	m.unionsRequired = len(grid)*len(grid[0]) - 1

	return nil
}

func (m *mazeState) Iterate(grid Grid) error {
	if m.wallIdx >= len(m.walls) {
		return errors.New("wall index out of wall slice range. Must be error in IsComplete func")
	}

	currWall := m.walls[m.wallIdx]
	m.wallIdx++

	tile1 := currWall.tile1
	tile2 := currWall.tile2

	if !m.tileSets.AreConnected(tile1, tile2) {
		utils.RemoveWalls(tile1, tile2)
		m.tileSets.Union(tile1, tile2)
		m.unionCount++
	}

	return nil
}

func (m *mazeState) IsComplete() bool {
	return m.unionCount == m.unionsRequired
}
