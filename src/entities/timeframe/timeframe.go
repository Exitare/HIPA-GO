package entities

//TimeFrame represents a TimeFrame struct
type TimeFrame struct {
	ID                      uint32
	Value                   float64
	IncludingMinute         float64
	AboveBelowCellThreshold bool
}
