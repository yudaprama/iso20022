package model

// Choice between an ISODate or ISODateTime format or a date code or a date code and a time.
type DateFormat7Choice struct {

	// Date expressed as a calendar date.
	Date *DateAndDateTimeChoice `xml:"Dt"`

	// Specifies  a date code and a time.
	DateCodeAndTime *DateCodeAndTimeFormat1 `xml:"DtCdAndTm"`

	// Specifies the type of date.
	DateCode *DateCode2Choice `xml:"DtCd"`
}

func (d *DateFormat7Choice) AddDate() *DateAndDateTimeChoice {
	d.Date = new(DateAndDateTimeChoice)
	return d.Date
}

func (d *DateFormat7Choice) AddDateCodeAndTime() *DateCodeAndTimeFormat1 {
	d.DateCodeAndTime = new(DateCodeAndTimeFormat1)
	return d.DateCodeAndTime
}

func (d *DateFormat7Choice) AddDateCode() *DateCode2Choice {
	d.DateCode = new(DateCode2Choice)
	return d.DateCode
}
