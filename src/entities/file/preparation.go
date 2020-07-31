package hipafile

import (
	"fmt"
)

func Prepare() error {
	for _, inputFile := range InputFiles {
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