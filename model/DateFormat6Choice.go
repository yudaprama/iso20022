package model

// Choice between an ISODate or ISODateTime format or a date code.
type DateFormat6Choice struct {

	// Date expressed as a calendar date.
	Date *DateAndDateTimeChoice `xml:"Dt"`

	// Specifies the type of date.
	DateCode *DateCode2Choice `xml:"DtCd"`
}

func (d *DateFormat6Choice) AddDate() *DateAndDateTimeChoice {
	d.Date = new(DateAndDateTimeChoice)
	return d.Date
}

func (d *DateFormat6Choice) AddDateCode() *DateCode2Choice {
	d.DateCode = new(DateCode2Choice)
	return d.DateCode
}
