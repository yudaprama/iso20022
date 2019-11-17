package model

// Choice between an ISODate or ISODateTime format or a date code or a date code and a time.
type DateFormat38Choice struct {

	// Date expressed as a calendar date.
	Date *DateAndDateTimeChoice `xml:"Dt"`

	// Specifies  a date code and a time.
	DateCodeAndTime *DateCodeAndTimeFormat4 `xml:"DtCdAndTm"`

	// Specifies the type of date.
	DateCode *DateCode22Choice `xml:"DtCd"`
}

func (d *DateFormat38Choice) AddDate() *DateAndDateTimeChoice {
	d.Date = new(DateAndDateTimeChoice)
	return d.Date
}

func (d *DateFormat38Choice) AddDateCodeAndTime() *DateCodeAndTimeFormat4 {
	d.DateCodeAndTime = new(DateCodeAndTimeFormat4)
	return d.DateCodeAndTime
}

func (d *DateFormat38Choice) AddDateCode() *DateCode22Choice {
	d.DateCode = new(DateCode22Choice)
	return d.DateCode
}
