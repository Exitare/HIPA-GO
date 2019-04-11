package main

type Cell struct {
	Name                 string
	TimeFrames           []TimeFrame
	Threshold            float32
	BaselineMean         float32
	NormalizedTimeFrames []TimeFrame
	TimeFrameMaximum     float32
	HighIntensityCounts  map[float64]int32
}

