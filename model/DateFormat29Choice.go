package model

// Choice between an ISODate or ISODateTime format or a date code.
type DateFormat29Choice struct {

	// Date expressed as a calendar date.
	DateOrDateTime *DateAndDateTimeChoice `xml:"DtOrDtTm"`

	// Specifies the type of date.
	DateCode *DateType1Code `xml:"DtCd"`
}

func (d *DateFormat29Choice) AddDateOrDateTime() *DateAndDateTimeChoice {
	d.DateOrDateTime = new(DateAndDateTimeChoice)
	return d.DateOrDateTime
}

func (d *DateFormat29Choice) SetDateCode(value string) {
	d.DateCode = (*DateType1Code)(&value)
}
