package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Tile struct {
	img   *ebiten.Image
	posX  float64
	posY  float64
	wallN bool
	wallE bool
	wallS bool
	wallW bool
}

type (
	Grid [][]Tile
)

func createTile(g *game, posX, posY float64) Tile {
	var tile Tile
	tile.img = ebiten.NewImage(g.cfg.tileSize, g.cfg.tileSize)

	tile.img.Fill(color.RGBA{255, 255, 255, 255})

	tile.posX = posX
	tile.posY = posY
	tile.wallN = true
	tile.wallE = true
	tile.wallS = true
	tile.wallW = true
	return tile
}
