// Package utils defines Tile type. Recommend to define Tile alias. Package contains tile utility functions
package utils

import (
	"math/rand"
)

type Tile struct {
	PosX    float64
	PosY    float64
	Row     int
	Col     int
	WallN   bool
	WallE   bool
	WallS   bool
	WallW   bool
	Visited bool
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
