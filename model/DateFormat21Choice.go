package model

// Specifies the value of a date.
type DateFormat21Choice struct {

	// Date expressed as an ISO Date.
	Date *DateAndDateTimeChoice `xml:"Dt"`

	// Date not specified, for example, the date is unknown.
	NotSpecifiedDate *DateType8Code `xml:"NotSpcfdDt"`
}

func (d *DateFormat21Choice) AddDate() *DateAndDateTimeChoice {
	d.Date = new(DateAndDateTimeChoice)
	return d.Date
}

func (d *DateFormat21Choice) SetNotSpecifiedDate(value string) {
	d.NotSpecifiedDate = (*DateType8Code)(&value)
}
