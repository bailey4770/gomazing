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

func removeWalls(frontierTile *Tile, visitedTile *Tile) {
	type wallPair struct {
		frontierWall *bool
		visitedWall  *bool
	}
	var walls wallPair

	if visitedTile.Col == frontierTile.Col {
		if frontierTile.Row < visitedTile.Row {
			// visited is to the south
			walls = wallPair{&frontierTile.WallS, &visitedTile.WallN}
		} else if visitedTile.Row < frontierTile.Row {
			// visited to the north
			walls = wallPair{&frontierTile.WallN, &visitedTile.WallS}
		}
	} else if visitedTile.Row == frontierTile.Row {
		if frontierTile.Col < visitedTile.Col {
			// visited to west
			walls = wallPair{&frontierTile.WallE, &visitedTile.WallW}
		} else if visitedTile.Col < frontierTile.Col {
			// visited to east
			walls = wallPair{&frontierTile.WallW, &visitedTile.WallE}
		}
	}

	*walls.frontierWall = false
	*walls.visitedWall = false
}

func Iterate(frontier map[*Tile]struct{}, grid Grid, maxRows, maxCols int) map[*Tile]struct{} {
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
		if n.Visited {
			visitedNeighbours = append(visitedNeighbours, n)
		} else if _, ok := frontier[n]; !ok {
			frontier[n] = struct{}{}
		}
	}

	if len(visitedNeighbours) == 0 {
		return frontier
	}

	// choose random tile from visited neighbours
	randomIndex := rand.Intn(len(visitedNeighbours))
	visitedTile := visitedNeighbours[randomIndex]

	removeWalls(frontierTile, visitedTile)
	frontierTile.Visited = true

	return frontier
}
