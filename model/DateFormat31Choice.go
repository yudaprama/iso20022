package model

// Choice between an ISODate or ISODateTime format or a date code.
type DateFormat31Choice struct {

	// Date expressed as a calendar date.
	Date *DateAndDateTimeChoice `xml:"Dt"`

	// Specifies the type of date.
	DateCode *DateCode19Choice `xml:"DtCd"`
}

func (d *DateFormat31Choice) AddDate() *DateAndDateTimeChoice {
	d.Date = new(DateAndDateTimeChoice)
	return d.Date
}

func (d *DateFormat31Choice) AddDateCode() *DateCode19Choice {
	d.DateCode = new(DateCode19Choice)
	return d.DateCode
}
