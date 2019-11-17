package model

// Choice between the formats to express a date and time.
type DateFormat2Choice struct {

	// Date expressed as a calendar date.
	Date *ISODateTime `xml:"Dt"`

	// Indicates that date and time are unknown.
	DateCode *DateType1Code `xml:"DtCd"`
}

func (d *DateFormat2Choice) SetDate(value string) {
	d.Date = (*ISODateTime)(&value)
}

func (d *DateFormat2Choice) SetDateCode(value string) {
	d.DateCode = (*DateType1Code)(&value)
}
