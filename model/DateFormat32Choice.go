package model

// Choice between an ISODate or ISODateTime format or a date code or a date code and a time.
type DateFormat32Choice struct {

	// Date expressed as a calendar date.
	Date *DateAndDateTimeChoice `xml:"Dt"`

	// Specifies  a date code and a time.
	DateCodeAndTime *DateCodeAndTimeFormat3 `xml:"DtCdAndTm"`

	// Specifies the type of date.
	DateCode *DateCode19Choice `xml:"DtCd"`
}

func (d *DateFormat32Choice) AddDate() *DateAndDateTimeChoice {
	d.Date = new(DateAndDateTimeChoice)
	return d.Date
}

func (d *DateFormat32Choice) AddDateCodeAndTime() *DateCodeAndTimeFormat3 {
	d.DateCodeAndTime = new(DateCodeAndTimeFormat3)
	return d.DateCodeAndTime
}

func (d *DateFormat32Choice) AddDateCode() *DateCode19Choice {
	d.DateCode = new(DateCode19Choice)
	return d.DateCode
}
