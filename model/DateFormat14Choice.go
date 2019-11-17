package model

// Choice between an ISODate format or a date code.
type DateFormat14Choice struct {

	// Date expressed as a calendar date.
	Date *ISODate `xml:"Dt"`

	// Specifies the type of date.
	DateCode *DateCode9Choice `xml:"DtCd"`
}

func (d *DateFormat14Choice) SetDate(value string) {
	d.Date = (*ISODate)(&value)
}

func (d *DateFormat14Choice) AddDateCode() *DateCode9Choice {
	d.DateCode = new(DateCode9Choice)
	return d.DateCode
}
