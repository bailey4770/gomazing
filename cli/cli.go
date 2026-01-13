// package cli reads command line args and parses them into config struct. Call GetConfig() to get struct.
package cli

import (
	"flag"
	"fmt"
	"image/color"

	"github.com/bailey4770/gomazing/generators/dfs"
	"github.com/bailey4770/gomazing/generators/kruskals"
	"github.com/bailey4770/gomazing/generators/prims"
	"github.com/bailey4770/gomazing/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

type Generator interface {
	Initialise(utils.Grid) error
	Iterate(utils.Grid) error
	IsComplete() bool
}

func GetGenerators() map[string]Generator {
	return map[string]Generator{
		"prims":    prims.GetMazeState(),
		"dfs":      dfs.GetMazeState(),
		"kruskals": kruskals.GetMazeState(),
	}
}

func getGeneratorNames(generators map[string]Generator) []string {
	var names []string
	for name := range generators {
		names = append(names, name)
	}
	return names
}

type Config struct {
	Generator     Generator
	WindowWidth   int
	WindowHeight  int
	TileSize      int
	WallThickness int
	MaxRows       int
	MaxCols       int
	Speed         int
	// size of *ebiten.Image == 8 == size of int
	WallImg   *ebiten.Image
	ShowStats bool
}

func GetConfig() Config {
	var generator string
	var windowWidth, windowHeight, tileSize, wallThickness, gameSpeed int
	var showStats bool

	generators := GetGenerators()
	generatorUsage := fmt.Sprintf("Input maze generation algorithm %v", getGeneratorNames(generators))
	flag.StringVar(&generator, "gen", "prims", generatorUsage)

	flag.IntVar(&windowWidth, "width", 640, "Input window width")
	flag.IntVar(&windowHeight, "height", 480, "Input window height")
	flag.IntVar(&tileSize, "tile", 20, "Input tile size")
	flag.IntVar(&wallThickness, "wall", 1, "Input cell wall thickness")
	flag.IntVar(&gameSpeed, "speed", 3, "Input game speed")

	flag.BoolVar(&showStats, "debug", false, "Show FPS and TPS info")
	flag.Parse()

	wallImg := ebiten.NewImage(1, 1)
	wallImg.Fill(color.White)

	return Config{
		Generator:     generators[generator],
		WindowWidth:   windowWidth,
		WindowHeight:  windowHeight,
		TileSize:      tileSize,
		WallThickness: wallThickness,
		MaxRows:       windowHeight / tileSize,
		MaxCols:       windowWidth / tileSize,
		Speed:         gameSpeed,
		WallImg:       wallImg,
		ShowStats:     showStats,
	}
}
