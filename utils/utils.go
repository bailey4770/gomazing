// Package utils defines Tile type. Recommend to define Tile alias. Package contains tile utility functions
package utils

import (
	"math/rand"
)

type Tile struct {
	PosX  float64
	PosY  float64
	Row   int
	Col   int
	WallN bool
	WallE bool
	WallS bool
	WallW bool
}

type (
	Grid [][]*Tile
)

func GetRandomTile(tiles []*Tile) *Tile {
	if len(tiles) == 0 {
		return nil
	}

	// choose random tile from frontier list
	randomIndex := rand.Intn(len(tiles))
	return tiles[randomIndex]
}

func FindNeighbours(t *Tile, grid Grid, maxRows, maxCols int) []*Tile {
	dirs := [4]struct{ rowOffset, colOffset int }{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}

	neighbours := make([]*Tile, 0, 4)

	for _, dir := range dirs {
		newRow := t.Row + dir.rowOffset
		newCol := t.Col + dir.colOffset

		if (newRow >= 0 && newRow < maxRows) && (newCol >= 0 && newCol < maxCols) {
			neighbours = append(neighbours, grid[newRow][newCol])
		}
	}

	return neighbours
}

func RemoveWalls(tile1 *Tile, tile2 *Tile) {
	type wallPair struct {
		frontierWall *bool
		visitedWall  *bool
	}
	var walls wallPair

	if tile2.Col == tile1.Col {
		if tile1.Row < tile2.Row {
			// visited is to the south
			walls = wallPair{&tile1.WallS, &tile2.WallN}
		} else if tile2.Row < tile1.Row {
			// visited to the north
			walls = wallPair{&tile1.WallN, &tile2.WallS}
		}
	} else if tile2.Row == tile1.Row {
		if tile1.Col < tile2.Col {
			// visited to west
			walls = wallPair{&tile1.WallE, &tile2.WallW}
		} else if tile2.Col < tile1.Col {
			// visited to east
			walls = wallPair{&tile1.WallW, &tile2.WallE}
		}
	}

	*walls.frontierWall = false
	*walls.visitedWall = false
}
