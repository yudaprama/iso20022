package model

// Time span defined by a start date and time, and an end date and time.
type DateTimePeriodDetails struct {

	// Date and time at which the range starts.
	FromDateTime *ISODateTime `xml:"FrDtTm"`

	// Date and time at which the range ends.
	ToDateTime *ISODateTime `xml:"ToDtTm"`
}

func (d *DateTimePeriodDetails) SetFromDateTime(value string) {
	d.FromDateTime = (*ISODateTime)(&value)
}

func (d *DateTimePeriodDetails) SetToDateTime(value string) {
	d.ToDateTime = (*ISODateTime)(&value)
}
