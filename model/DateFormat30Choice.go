package model

// Choice between an ISODate format or a date code.
type DateFormat30Choice struct {

	// Date expressed as a calendar date.
	Date *ISODate `xml:"Dt"`

	// Specifies the type of date.
	DateCode *DateCode19Choice `xml:"DtCd"`
}

func (d *DateFormat30Choice) SetDate(value string) {
	d.Date = (*ISODate)(&value)
}

func (d *DateFormat30Choice) AddDateCode() *DateCode19Choice {
	d.DateCode = new(DateCode19Choice)
	return d.DateCode
}
