package model

// Choice between an ISODate format or a date code.
type DateFormat5Choice struct {

	// Date expressed as a calendar date.
	Date *ISODate `xml:"Dt"`

	// Specifies the type of date.
	DateCode *DateCode2Choice `xml:"DtCd"`
}

func (d *DateFormat5Choice) SetDate(value string) {
	d.Date = (*ISODate)(&value)
}

func (d *DateFormat5Choice) AddDateCode() *DateCode2Choice {
	d.DateCode = new(DateCode2Choice)
	return d.DateCode
}
