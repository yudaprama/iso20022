package model

// Choice between a date or a date and time format.
type DateAndDateTime1Choice struct {

	// Numeric representation of the day of the month and year.
	Date *ISODate `xml:"Dt"`

	// Numeric representation of time of the day and the day of the month and year.
	DateTime *ISODateTime `xml:"DtTm"`
}

func (d *DateAndDateTime1Choice) SetDate(value string) {
	d.Date = (*ISODate)(&value)
}

func (d *DateAndDateTime1Choice) SetDateTime(value string) {
	d.DateTime = (*ISODateTime)(&value)
}
