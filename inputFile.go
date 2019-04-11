package main

import (
	"fmt"
	"io/ioutil"
	"log"
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
func ReadFile(inputFile *InputFile) {
	content, err := ioutil.ReadFile("testdata/hello")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("File contents: %s", content)
}

func (inputFile *InputFile) showContent() {
	fmt.Println(inputFile.Content)
}

func (inputFile *InputFile) getFileName() {
	splits := strings.Split(inputFile.Name, ".")
	inputFile.Name = splits[0]
}

func (inputFile *InputFile) countLines() {
	var count int

	lines := strings.Split(inputFile.Content, "\n")

	for count < len(lines) {
		count++
	}

	fmt.Println("Count is", count)
}

func (inputFile *InputFile) cellBuilder() {
	//lines := strings.Split(inputFile.Content, "\n")

}
