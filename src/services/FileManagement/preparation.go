package FileManagement

import (
	"fmt"
	"globals"
)

func Prepare() error {
	for _, inputFile := range globals.InputFiles {
		err := inputFile.ReadContent()
		if err != nil {
			return err
		}
		inputFile.CountTimeFrames()
		inputFile.CountCells()
		inputFile.CreateCells()
		inputFile.PopulateCells()
		fmt.Println(inputFile.CellCount)
	}

	return nil
}