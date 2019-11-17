package model

// Provides the date.
type DateFormat15Choice struct {

	// Date expressed as a calendar date.
	Date *ISODate `xml:"Dt"`

	// Date is expressed using a code.
	DateCode *DateCode3Choice `xml:"DtCd"`
}

func (d *DateFormat15Choice) SetDate(value string) {
	d.Date = (*ISODate)(&value)
}

func (d *DateFormat15Choice) AddDateCode() *DateCode3Choice {
	d.DateCode = new(DateCode3Choice)
	return d.DateCode
}
