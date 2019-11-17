package model

// Particular time span specified between a start time and an end time. The time period cannot exceed 24 hours.
type TimePeriodDetails1 struct {

	// Time at which the time span starts.
	FromTime *ISOTime `xml:"FrTm"`

	// Time at which the time span ends.
	ToTime *ISOTime `xml:"ToTm,omitempty"`
}

func (t *TimePeriodDetails1) SetFromTime(value string) {
	t.FromTime = (*ISOTime)(&value)
}

func (t *TimePeriodDetails1) SetToTime(value string) {
	t.ToTime = (*ISOTime)(&value)
}
