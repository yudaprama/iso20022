package model

// Time span defined by a start date and time, and an end date and time.
type DateTimePeriodDetails1 struct {

	// Date and time at which the range starts.
	FromDateTime *ISODateTime `xml:"FrDtTm"`

	// Date and time at which the range ends.
	ToDateTime *ISODateTime `xml:"ToDtTm,omitempty"`
}

func (d *DateTimePeriodDetails1) SetFromDateTime(value string) {
	d.FromDateTime = (*ISODateTime)(&value)
}

func (d *DateTimePeriodDetails1) SetToDateTime(value string) {
	d.ToDateTime = (*ISODateTime)(&value)
}
