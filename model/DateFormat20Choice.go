package model

// Choice between an ISODate or ISODateTime format or a date code or a date code and a time.
type DateFormat20Choice struct {

	// Date expressed as a calendar date.
	Date *DateAndDateTimeChoice `xml:"Dt"`

	// Specifies  a date code and a time.
	DateCodeAndTime *DateCodeAndTimeFormat1 `xml:"DtCdAndTm"`

	// Specifies the type of date.
	DateCode *DateCode11Choice `xml:"DtCd"`
}

func (d *DateFormat20Choice) AddDate() *DateAndDateTimeChoice {
	d.Date = new(DateAndDateTimeChoice)
	return d.Date
}

func (d *DateFormat20Choice) AddDateCodeAndTime() *DateCodeAndTimeFormat1 {
	d.DateCodeAndTime = new(DateCodeAndTimeFormat1)
	return d.DateCodeAndTime
}

func (d *DateFormat20Choice) AddDateCode() *DateCode11Choice {
	d.DateCode = new(DateCode11Choice)
	return d.DateCode
}
