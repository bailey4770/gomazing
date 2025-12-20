// package cli reads command line args and parses them into config struct. Call GetConfig() to get struct.
package cli

import (
	"flag"
	"fmt"

	"github.com/bailey4770/gomazing/generators/dfs"
	"github.com/bailey4770/gomazing/generators/prims"
	"github.com/bailey4770/gomazing/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

type Generator interface {
	Initialise(utils.Grid)
	Iterate(utils.Grid)
	IsComplete() bool
}

func GetGenerators() map[string]Generator {
	return map[string]Generator{
		"prims": prims.GetMazeState(),
		"dfs":   dfs.GetMazeState(),
	}
}

func getGeneratorNames(generators map[string]Generator) []string {
	var names []string
	for k := range generators {
		names = append(names, k)
	}
	return names
}

type Config struct {
	WindowWidth   int
	WindowHeight  int
	TileSize      int
	WallThickness int
	MaxRows       int
	MaxCols       int
	Speed         int
	// size of *ebiten.Image == 8 == size of int
	WallImg   *ebiten.Image
	Generator Generator
}

func GetConfig() Config {
	var windowWidth, windowHeight, tileSize, wallThickness, gameSpeed int
	var generator string

	flag.IntVar(&windowWidth, "width", 640, "Input window width")
	flag.IntVar(&windowHeight, "height", 480, "Input window height")
	flag.IntVar(&tileSize, "tile", 20, "Input tile size")
	flag.IntVar(&wallThickness, "wall", 1, "Input cell wall thickness")
	flag.IntVar(&gameSpeed, "speed", 3, "Input game speed")

	generators := GetGenerators()
	generatorUsage := fmt.Sprintf("Input maze generation algorithm %v", getGeneratorNames(generators))
	flag.StringVar(&generator, "gen", "prims", generatorUsage)

	flag.Parse()

	return Config{
		WindowWidth:   windowWidth,
		WindowHeight:  windowHeight,
		TileSize:      tileSize,
		WallThickness: wallThickness,
		MaxRows:       windowHeight / tileSize,
		MaxCols:       windowWidth / tileSize,
		Speed:         gameSpeed,
		WallImg:       ebiten.NewImage(1, 1),
		Generator:     generators[generator],
	}
}
