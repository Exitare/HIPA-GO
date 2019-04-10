package main

import (
	"fmt"
	"strings"
)

type InputFile struct {
	Name                 string
	Content              string
	Lines                uint32
	PercentageLimit      float32
	Cells                []Cell
	CellCount            uint64
	RowCount             uint64
	TimeFrameCount       uint64
	TotalDetectedMinutes uint32
	StimulationTimeFrame uint32
}

func (inputFile *InputFile) showContent() {
	fmt.Println(inputFile.Content)
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