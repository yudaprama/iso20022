package model

// Choice between an ISODate or ISODateTime format or a date code.
type DateFormat33Choice struct {

	// Date expressed as a calendar date.
	Date *DateAndDateTimeChoice `xml:"Dt"`

	// Specifies the type of date.
	DateCode *DateCode20Choice `xml:"DtCd"`
}

func (d *DateFormat33Choice) AddDate() *DateAndDateTimeChoice {
	d.Date = new(DateAndDateTimeChoice)
	return d.Date
}

func (d *DateFormat33Choice) AddDateCode() *DateCode20Choice {
	d.DateCode = new(DateCode20Choice)
	return d.DateCode
}
