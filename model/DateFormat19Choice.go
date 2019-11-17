package model

// Choice between an ISODate or ISODateTime format or a date code.
type DateFormat19Choice struct {

	// Date expressed as a calendar date.
	Date *DateAndDateTimeChoice `xml:"Dt"`

	// Specifies the type of date.
	DateCode *DateCode11Choice `xml:"DtCd"`
}

func (d *DateFormat19Choice) AddDate() *DateAndDateTimeChoice {
	d.Date = new(DateAndDateTimeChoice)
	return d.Date
}

func (d *DateFormat19Choice) AddDateCode() *DateCode11Choice {
	d.DateCode = new(DateCode11Choice)
	return d.DateCode
}
