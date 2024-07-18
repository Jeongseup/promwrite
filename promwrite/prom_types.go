package promwrite

import "time"

// TSList is a slice of TimeSeries.
// type TSList []TimeSeries

// TimeSeries are made of labels and a datapoint.
// One timeseries has two components, one is label and the other is sample(data)
type TimeSeries struct {
	Labels []Label
	Sample Sample
}

// Label is a metric label.
type Label struct {
	Name  string
	Value string
}

type Sample struct {
	Time  time.Time
	Value float64
}
