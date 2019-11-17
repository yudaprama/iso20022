package model

// Choice between an ISODate format or a date code.
type DateFormat16Choice struct {

	// Date expressed as a calendar date.
	Date *ISODate `xml:"Dt"`

	// Specifies the type of date.
	DateCode *DateCode10Choice `xml:"DtCd"`
}

func (d *DateFormat16Choice) SetDate(value string) {
	d.Date = (*ISODate)(&value)
}

func (d *DateFormat16Choice) AddDateCode() *DateCode10Choice {
	d.DateCode = new(DateCode10Choice)
	return d.DateCode
}
