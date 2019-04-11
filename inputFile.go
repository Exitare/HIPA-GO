package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

//InputFile represents an uploaded file
type InputFile struct {
	Name                 string
	Content              string
	Foldername           string
	Lines                uint32
	PercentageLimit      float32
	Cells                []Cell
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

func (inputFile *InputFile) countLines() {
	inputFile.Lines = uint32(len(strings.Split(inputFile.Content, "\n")))
}

func (inputFile *InputFile) cellBuilder() {
	//lines := strings.Split(inputFile.Content, "\n")

}
