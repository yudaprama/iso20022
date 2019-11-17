package model

// Time span defined by a start date and time, and an end date and time.
type DateTimePeriod1 struct {

	// Date and time at which the period starts.
	FromDateTime *ISODateTime `xml:"FrDtTm"`

	// Date and time at which the period ends.
	ToDateTime *ISODateTime `xml:"ToDtTm"`
}

func (d *DateTimePeriod1) SetFromDateTime(value string) {
	d.FromDateTime = (*ISODateTime)(&value)
}

func (d *DateTimePeriod1) SetToDateTime(value string) {
	d.ToDateTime = (*ISODateTime)(&value)
}
