package model

// Specifies  a date code and a time.
type DateCodeAndTimeFormat4 struct {

	// Specifies the type of date.
	DateCode *DateCode26Choice `xml:"DtCd"`

	// Specifies the time.
	Time *ISOTime `xml:"Tm"`
}

func (d *DateCodeAndTimeFormat4) AddDateCode() *DateCode26Choice {
	d.DateCode = new(DateCode26Choice)
	return d.DateCode
}

func (d *DateCodeAndTimeFormat4) SetTime(value string) {
	d.Time = (*ISOTime)(&value)
}
