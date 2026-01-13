// package cli reads command line args and parses them into config struct. Call GetConfig() to get struct.
package cli

import (
	"flag"
	"fmt"
	"image/color"
	"log"

	"github.com/bailey4770/gomazing/generators/dfs"
	"github.com/bailey4770/gomazing/generators/kruskals"
	"github.com/bailey4770/gomazing/generators/prims"
	"github.com/bailey4770/gomazing/mazesave"
	"github.com/bailey4770/gomazing/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

type Generator interface {
	Initialise(utils.Grid) error
	Iterate(utils.Grid) error
	IsComplete() bool
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
	MazePath  string
}

func GetConfig() Config {
	var generatorName, mazePath string
	var numRows, numCols, tileSize, wallThickness, gameSpeed int
	var showStats bool

	generators := GetGenerators()
	generatorUsage := fmt.Sprintf("Mutually exclusive with load. Input maze generation algorithm %v", getGeneratorNames(generators))
	flag.StringVar(&generatorName, "gen", "prims", generatorUsage)
	flag.StringVar(&mazePath, "load", "", "Mutually exclusive with gen. Load a saved maze from file")
	// TODO: usage for load prints available mazes to load

	flag.IntVar(&numRows, "rows", 24, "Input number of rows")
	flag.IntVar(&numCols, "cols", 32, "Input number of cols")
	flag.IntVar(&tileSize, "tile", 20, "Input desired size of each tile")
	flag.IntVar(&wallThickness, "wall", 1, "Input cell wall thickness")
	flag.IntVar(&gameSpeed, "speed", 3, "Input game speed")

	flag.BoolVar(&showStats, "debug", false, "Show FPS and TPS info")
	flag.Parse()

	wallImg := ebiten.NewImage(1, 1)
	wallImg.Fill(color.White)

	loadFlagged := checkFlags(mazePath)

	var generator Generator
	if !loadFlagged {
		generator = generators[generatorName]
	} else {
		generator = nil

		var err error
		numRows, numCols, tileSize, err = mazesave.GetMazeDimensions(mazePath)
		if err != nil {
			log.Fatalf("could not load maze dimensions from file: %v", err)
		}
	}

	windowHeight, windowWidth := getWindowDimensions(numRows, numCols, tileSize)

	return Config{
		Generator:     generator,
		WindowWidth:   windowWidth,
		WindowHeight:  windowHeight,
		TileSize:      windowHeight / numRows,
		WallThickness: wallThickness,
		MaxRows:       numRows,
		MaxCols:       numCols,
		Speed:         gameSpeed,
		WallImg:       wallImg,
		ShowStats:     showStats,
		MazePath:      mazePath,
	}
}

func getWindowDimensions(numRows, numCols, tileSize int) (int, int) {
	return numRows * tileSize, numCols * tileSize
}

func checkFlags(mazePath string) bool {
	loadFlagged := false
	genFlagged := false

	flag.Visit(func(f *flag.Flag) {
		switch f.Name {
		case "load":
			if mazePath == "" {
				log.Fatal("Error: must provide maze file name")
			}
			loadFlagged = true

		case "gen":
			genFlagged = true
		}
	})

	if loadFlagged && genFlagged {
		log.Fatal("Error: cannot gen and load a maze. Commands are mutually exclusive.")
	}

	return loadFlagged
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
