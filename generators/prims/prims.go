// Package prims handles one iteration of the prims mazze generation algorithm. Repeatedly call from ebiten's Update game method.
package prims

import (
	"math/rand"

	"github.com/bailey4770/gomazing/utils"
)

type (
	Tile = utils.Tile
	Grid = utils.Grid
)

func Iterate(frontier map[*Tile]struct{}, visited map[*Tile]struct{}, grid Grid, maxRows, maxCols int) (map[*Tile]struct{}, map[*Tile]struct{}) {
	// getRandomTile takes a slice for efficient random selection.
	// no good way of randomly selecting an element from a map in Go. Frontier remains a map for efficient contains checking and easy deletion from queue.
	var frontierSlice []*Tile
	for t := range frontier {
		frontierSlice = append(frontierSlice, t)
	}
	frontierTile := utils.GetRandomTile(frontierSlice)
	delete(frontier, frontierTile)

	neighbours := utils.FindNeighbours(frontierTile, grid, maxRows, maxCols)
	var visitedNeighbours []*Tile
	for _, n := range neighbours {
		if _, ok := visited[n]; ok {
			visitedNeighbours = append(visitedNeighbours, n)
		} else if _, ok := frontier[n]; !ok {
			frontier[n] = struct{}{}
		}
	}

	if len(visitedNeighbours) == 0 {
		return frontier, visited
	}

	// choose random tile from visited neighbours
	randomIndex := rand.Intn(len(visitedNeighbours))
	visitedTile := visitedNeighbours[randomIndex]

	utils.RemoveWalls(frontierTile, visitedTile)
	visited[frontierTile] = struct{}{}

	return frontier, visited
}
