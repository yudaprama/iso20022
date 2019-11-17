package model

// Specifies  a date code and a time.
type DateCodeAndTimeFormat1 struct {

	// Specifies the type of date.
	DateCode *DateCode4Choice `xml:"DtCd"`

	// Specifies the time.
	Time *ISOTime `xml:"Tm"`
}

func (d *DateCodeAndTimeFormat1) AddDateCode() *DateCode4Choice {
	d.DateCode = new(DateCode4Choice)
	return d.DateCode
}

func (d *DateCodeAndTimeFormat1) SetTime(value string) {
	d.Time = (*ISOTime)(&value)
}
