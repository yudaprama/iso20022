package model

// Choice between an ISODate format or a date code.
type DateFormat41Choice struct {

	// Date expressed as a calendar date.
	Date *ISODate `xml:"Dt"`

	// Specifies the type of date.
	DateCode *DateCode22Choice `xml:"DtCd"`
}

func (d *DateFormat41Choice) SetDate(value string) {
	d.Date = (*ISODate)(&value)
}

func (d *DateFormat41Choice) AddDateCode() *DateCode22Choice {
	d.DateCode = new(DateCode22Choice)
	return d.DateCode
}
