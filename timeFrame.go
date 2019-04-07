package main

type TimeFrame struct {
	ID                      uint32
	Value                   float64
	IncludingMinute         uint32
	AboveBelowCellThreshold bool
}
