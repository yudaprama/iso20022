package model

// Choice between various date time patterns.
type DateTimePeriodChoice struct {

	// Date and time at which the range starts.
	FromDateTime *ISODateTime `xml:"FrDtTm"`

	// Date and time at which the range ends.
	ToDateTime *ISODateTime `xml:"ToDtTm"`

	// Range of time between a start date and time and an end date and time.
	DateTimeRange *DateTimePeriodDetails `xml:"DtTmRg"`
}

func (d *DateTimePeriodChoice) SetFromDateTime(value string) {
	d.FromDateTime = (*ISODateTime)(&value)
}

func (d *DateTimePeriodChoice) SetToDateTime(value string) {
	d.ToDateTime = (*ISODateTime)(&value)
}

func (d *DateTimePeriodChoice) AddDateTimeRange() *DateTimePeriodDetails {
	d.DateTimeRange = new(DateTimePeriodDetails)
	return d.DateTimeRange
}
