package model

// Choice between a date or a date and time format.
type DateAndDateTimeChoice struct {

	// Specified date.
	Date *ISODate `xml:"Dt"`

	// Specified date and time.
	DateTime *ISODateTime `xml:"DtTm"`
}

func (d *DateAndDateTimeChoice) SetDate(value string) {
	d.Date = (*ISODate)(&value)
}

func (d *DateAndDateTimeChoice) SetDateTime(value string) {
	d.DateTime = (*ISODateTime)(&value)
}
