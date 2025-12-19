// Package dfs runs one iteration of the dfs maze generation algorithm
package dfs

import (
	"github.com/bailey4770/gomazing/utils"
)

type (
	Tile = utils.Tile
	Grid = utils.Grid
)

func Iterate(stack []*Tile, visited map[*Tile]struct{}, curr *Tile, grid Grid, maxRows, maxCols int) ([]*Tile, map[*Tile]struct{}, *Tile) {
	neighbours := utils.FindNeighbours(curr, grid, maxRows, maxCols)
	var unvisitedNeighbours []*Tile
	for _, n := range neighbours {
		if _, ok := visited[n]; !ok {
			unvisitedNeighbours = append(unvisitedNeighbours, n)
		}
	}

	if len(unvisitedNeighbours) > 0 {
		randUnvisited := utils.GetRandomTile(unvisitedNeighbours)
		stack = append(stack, curr)
		utils.RemoveWalls(curr, randUnvisited)

		visited[randUnvisited] = struct{}{}
		curr = randUnvisited
		return stack, visited, curr

	} else if len(stack) > 0 {
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		return stack, visited, curr
	}

	return stack, visited, nil
}
