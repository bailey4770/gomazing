package mazesave

import (
	"log"
	"os"
	"testing"

	"github.com/bailey4770/gomazing/generators/prims"
	"github.com/bailey4770/gomazing/utils"
)

func TestSaveAndLoad(t *testing.T) {
	savedNumRows, savedNumCols, savedTileSize := 10, 10, 2
	savedGrid := initGrid(savedNumRows, savedNumCols, savedTileSize)
	mazeState := prims.GetMazeState()

	err := mazeState.Initialise(savedGrid)
	if err != nil {
		t.Fatal("could not initialise mazestate:", err)
	}

	for !mazeState.IsComplete() {
		err := mazeState.Iterate(savedGrid)
		if err != nil {
			t.Fatal("could not iterate maze state")
		}
	}

	filePath := "test.maze"
	defer func() {
		err := os.Remove(filePath)
		if err != nil {
			log.Fatalf("could not remove test file %s: %v", filePath, err)
		}
	}()

	err = SaveMaze(savedGrid, savedTileSize, filePath)
	if err != nil {
		t.Fatal("could not save maze:", err)
	}

	loadedNumRows, loadedNumCols, loadedTileSize, err := GetMazeDimensions(filePath)
	if err != nil {
		t.Fatalf("could not get maze dimensions: %v", err)
	}

	if savedNumRows != loadedNumRows {
		t.Fatalf("expected numRows to be %d but got %d", savedNumRows, loadedNumRows)
	}
	if savedNumCols != loadedNumCols {
		t.Fatalf("expected numCols to be %d but got %d", savedNumCols, loadedNumRows)
	}

	loadedGrid := initGrid(loadedNumRows, loadedNumCols, loadedTileSize)

	err = LoadMazeWalls(filePath, loadedGrid)
	if err != nil {
		t.Fatalf("could not load maze walls: %v", err)
	}

	for i, row := range loadedGrid {
		for j, loadedTile := range row {
			savedTile := savedGrid[i][j]

			if loadedTile.WallN != savedTile.WallN {
				t.Fatal("loaded wall does not match saved wall")
			}
			if loadedTile.WallE != savedTile.WallE {
				t.Fatal("loaded wall does not match saved wall")
			}
			if loadedTile.WallS != savedTile.WallS {
				t.Fatal("loaded wall does not match saved wall")
			}
			if loadedTile.WallW != savedTile.WallW {
				t.Fatal("loaded wall does not match saved wall")
			}
		}
	}
}

func initGrid(maxRows, maxCols, tileSize int) utils.Grid {
	// allocate row slices
	grid := make(utils.Grid, maxRows)

	for row := range grid {
		grid[row] = make([]*utils.Tile, maxCols)
		posY := float64(row * tileSize)

		for col := range grid[row] {
			posX := float64(col * tileSize)
			grid[row][col] = utils.CreateTile(posX, posY, row, col)
		}
	}

	return grid
}
