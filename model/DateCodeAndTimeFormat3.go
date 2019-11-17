package model

// Specifies  a date code and a time.
type DateCodeAndTimeFormat3 struct {

	// Specifies the type of date.
	DateCode *DateCode21Choice `xml:"DtCd"`

	// Specifies the time.
	Time *ISOTime `xml:"Tm"`
}

func (d *DateCodeAndTimeFormat3) AddDateCode() *DateCode21Choice {
	d.DateCode = new(DateCode21Choice)
	return d.DateCode
}

func (d *DateCodeAndTimeFormat3) SetTime(value string) {
	d.Time = (*ISOTime)(&value)
}
