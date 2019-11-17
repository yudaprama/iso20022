package model

// Choice between an ISODate or ISODateTime format or a date code.
type DateFormat39Choice struct {

	// Date expressed as a calendar date.
	Date *DateAndDateTimeChoice `xml:"Dt"`

	// Specifies the type of date.
	DateCode *DateCode27Choice `xml:"DtCd"`
}

func (d *DateFormat39Choice) AddDate() *DateAndDateTimeChoice {
	d.Date = new(DateAndDateTimeChoice)
	return d.Date
}

func (d *DateFormat39Choice) AddDateCode() *DateCode27Choice {
	d.DateCode = new(DateCode27Choice)
	return d.DateCode
}
