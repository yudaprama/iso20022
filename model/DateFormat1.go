package model

// Format to express a date and a date mode.
type DateFormat1 struct {

	// Date at which the event occurs.
	Date *DateFormat3Choice `xml:"Dt"`

	// Specifies whether an event for which a date is provided occurs typically at the "beginning of day" or at the "end of day".
	DateMode *DateMode1Code `xml:"DtMd,omitempty"`
}

func (d *DateFormat1) AddDate() *DateFormat3Choice {
	d.Date = new(DateFormat3Choice)
	return d.Date
}

func (d *DateFormat1) SetDateMode(value string) {
	d.DateMode = (*DateMode1Code)(&value)
}
