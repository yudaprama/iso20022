package model

// Choice between the formats to express a date.
type DateFormat3Choice struct {

	// Date expressed as a calendar date.
	Date *ISODate `xml:"Dt"`

	// Indicates that date is unknown.
	DateCode *DateType1Code `xml:"DtCd"`
}

func (d *DateFormat3Choice) SetDate(value string) {
	d.Date = (*ISODate)(&value)
}

func (d *DateFormat3Choice) SetDateCode(value string) {
	d.DateCode = (*DateType1Code)(&value)
}
