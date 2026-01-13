// Package mazesave handles encoding/decoding maze grids into/from binary format for saving and loading to storage
package mazesave

import (
	"encoding/binary"
	"fmt"
	"log"
	"os"

	"github.com/bailey4770/gomazing/utils"
)

func SaveMaze(grid utils.Grid, tileSize int, fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("could not create file %s: %v", fileName, err)
	}

	numRows, numCols := len(grid), len(grid[0])
	if err := binary.Write(file, binary.LittleEndian, uint16(numRows)); err != nil {
		return fmt.Errorf("could not write numRows to file: %v", err)
	}
	if err := binary.Write(file, binary.LittleEndian, uint16(numCols)); err != nil {
		return fmt.Errorf("could not write numCols to file: %v", err)
	}

	if err := binary.Write(file, binary.LittleEndian, uint16(tileSize)); err != nil {
		return fmt.Errorf("could not write tileSize to file: %v", err)
	}

	var bitBuffer byte
	var bitPos uint

	writeBit := func(wall bool) error {
		if wall {
			bitBuffer |= (1 << bitPos)
		}
		bitPos++

		if bitPos == 8 {
			// we have filled up byte
			if _, err := file.Write([]byte{bitBuffer}); err != nil {
				return err
			}
			bitBuffer, bitPos = 0, 0
		}

		return nil
	}

	for i, row := range grid {
		for j, tile := range row {
			// Write east wall (except for final col)
			if j < numCols-1 {
				if err := writeBit(tile.WallE); err != nil {
					return fmt.Errorf("could not write tile WallE to buffer: %v", err)
				}
			}

			// Write south wall (except for final row)
			if i < numRows-1 {
				if err := writeBit(tile.WallS); err != nil {
					return fmt.Errorf("could not write tile WallS to buffer: %v", err)
				}
			}
		}
	}

	if bitPos > 0 {
		if _, err := file.Write([]byte{bitBuffer}); err != nil {
			return err
		}
	}

	return nil
}

func GetMazeDimensions(filepath string) (int, int, int, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("could not read from %s: %v", filepath, err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatalf("Error closing file: %v", err)
		}
	}()

	var numRows, numCols, tileSize uint16
	if err := binary.Read(file, binary.LittleEndian, &numRows); err != nil {
		return 0, 0, 0, fmt.Errorf("could not read numRows: %v", err)
	}
	if err := binary.Read(file, binary.LittleEndian, &numCols); err != nil {
		return 0, 0, 0, fmt.Errorf("could not read numCols: %v", err)
	}
	if err := binary.Read(file, binary.LittleEndian, &tileSize); err != nil {
		return 0, 0, 0, fmt.Errorf("could not read tileSize: %v", err)
	}

	return int(numRows), int(numCols), int(tileSize), nil
}

func LoadMazeWalls(filepath string, grid utils.Grid) error {
	file, err := os.Open(filepath)
	if err != nil {
		return fmt.Errorf("could not read from %s: %v", filepath, err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatalf("Error closing file: %v", err)
		}
	}()

	// Skip past firsrt 6 bytes (dimension info already read)
	var offset int64 = 6
	if _, err := file.Seek(offset, 0); err != nil {
		return fmt.Errorf("could not seek: %v", err)
	}

	numRows := len(grid)
	numCols := len(grid[0])

	var bitBuffer byte
	var bitPos uint = 8 // so we kick off with full bitBuffer
	readBits := func() (bool, error) {
		if bitPos == 8 {
			if err := binary.Read(file, binary.LittleEndian, &bitBuffer); err != nil {
				return false, fmt.Errorf("could not read from file: %v", err)
			}
			bitPos = 0
		}

		// shift byte to right by bitPos, then AND to see if bit in buffer is 1
		bit := (bitBuffer >> bitPos) & 1
		bitPos++
		return bit == 1, nil
	}

	for i, row := range grid {
		for j, tile := range row {
			// read next bit as bool for east wall (so long as we aren't in final col)
			if j < numCols-1 {
				wall, err := readBits()
				if err != nil {
					return err
				} else if !wall {
					utils.RemoveWalls(tile, row[j+1])
				}
			}

			// read next bit as bool for south wall (so long as we aren't in final row)
			if i < numRows-1 {
				wall, err := readBits()
				if err != nil {
					return err
				} else if !wall {
					utils.RemoveWalls(tile, grid[i+1][j])
				}
			}
		}
	}

	return nil
}
