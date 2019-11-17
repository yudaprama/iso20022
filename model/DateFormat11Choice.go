package model

// Choice between an ISODate or ISODateTime format or a date code.
type DateFormat11Choice struct {

	// Date expressed as a calendar date.
	Date *DateAndDateTimeChoice `xml:"Dt"`

	// Specifies the type of date.
	DateCode *DateCode3Choice `xml:"DtCd"`
}

func (d *DateFormat11Choice) AddDate() *DateAndDateTimeChoice {
	d.Date = new(DateAndDateTimeChoice)
	return d.Date
}

func (d *DateFormat11Choice) AddDateCode() *DateCode3Choice {
	d.DateCode = new(DateCode3Choice)
	return d.DateCode
}
