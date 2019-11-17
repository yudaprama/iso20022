package model

// Choice between a date or a date and time format for a period.
type DateOrDateTimePeriodChoice struct {

	// Period expressed with dates.
	Date *DatePeriodDetails `xml:"Dt"`

	// Period expressed a dates and times.
	DateTime *DateTimePeriodDetails `xml:"DtTm"`
}

func (d *DateOrDateTimePeriodChoice) AddDate() *DatePeriodDetails {
	d.Date = new(DatePeriodDetails)
	return d.Date
}

func (d *DateOrDateTimePeriodChoice) AddDateTime() *DateTimePeriodDetails {
	d.DateTime = new(DateTimePeriodDetails)
	return d.DateTime
}
