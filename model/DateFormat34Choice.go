package model

// Choice between an ISODate or ISODateTime format or a date code.
type DateFormat34Choice struct {

	// Date expressed as a calendar date.
	Date *DateAndDateTimeChoice `xml:"Dt"`

	// Specifies the type of date.
	DateCode *DateCode22Choice `xml:"DtCd"`
}

func (d *DateFormat34Choice) AddDate() *DateAndDateTimeChoice {
	d.Date = new(DateAndDateTimeChoice)
	return d.Date
}

func (d *DateFormat34Choice) AddDateCode() *DateCode22Choice {
	d.DateCode = new(DateCode22Choice)
	return d.DateCode
}
