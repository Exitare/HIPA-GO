package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

//InputFile represents an uploaded file
type InputFile struct {
	Name                 string
	Content              string
	Foldername           string
	PercentageLimit      float32
	Cells                []*Cell
	CellCount            uint64
	RowCount             uint64
	TimeFrameCount       uint64
	TotalDetectedMinutes uint32
	StimulationTimeFrame uint32
}

//ReadFile reads the content of a file
func (inputFile *InputFile) readContent() error {
	content, err := ioutil.ReadFile(workingDir + inputFile.Foldername + "/" + inputFile.Name + ".txt")
	if err != nil {
		return err
	}
	inputFile.Content = string(content)
	return nil
}

func (inputFile *InputFile) resolveName(uploadName string) {
	uploadName = strings.Join(strings.Fields(uploadName), "")
	splits := strings.Split(uploadName, ".")

	for i := 0; i < len(splits)-1; i++ {
		if i == 0 {
			inputFile.Name = splits[i]
		} else {
			inputFile.Name = inputFile.Name + "_" + splits[i]
		}
	}

	fmt.Println(inputFile.Name)
}

func (inputFile *InputFile) countTimeFrames() {
	inputFile.TimeFrameCount = uint64(len(strings.Split(inputFile.Content, "\n"))) - 1
}

func (inputFile *InputFile) countCells() {
	lines := strings.Split(inputFile.Content, "\n")
	for line := range lines {
		strings.Trim(lines[line], " ")
		fmt.Printf("Line is: %s \n", lines[line])
		fmt.Printf("Line lengt is %d", len(lines[line]))
		if len(lines[line]) != 0 {
			values := strings.Split(lines[line], "\t")
			count := 0
			for range values {
				count++
			}
			fmt.Printf("Lenght is: %d \n", count)
			inputFile.CellCount = uint64(len(values))
		}
	}
}

func (inputFile *InputFile) createCells() {
	var timeFrames []*TimeFrame
	var highIntensityCounts map[float64]int32
	for i := 0; i < int(inputFile.CellCount); i++ {
		cell := Cell{"", timeFrames, 0.0, 0.0, timeFrames, 0.0, highIntensityCounts}
		inputFile.Cells = append(inputFile.Cells, &cell)
	}
}

func (inputFile *InputFile) populateCells() {

	lines := strings.Split(inputFile.Content, "\n")

	for line := range lines {
		cellValueList := strings.Split(lines[line], "\t")

		for cell := 0; cell < int(inputFile.CellCount); cell++ {
			if line == 0 {
				inputFile.Cells[cell].Name = cellValueList[cell]
			} else {
				if timeframeValue, err := strconv.ParseFloat(cellValueList[cell], 32); err == nil {
					timeframe := TimeFrame{uint32(line), timeframeValue, (float64(line) * 3.9 / 60), false}
					inputFile.Cells[cell].TimeFrames = append(inputFile.Cells[cell].TimeFrames, &timeframe)
				}
			}
		}
	}
}
