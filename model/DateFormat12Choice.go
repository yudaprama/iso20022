package model

// Specifies the value of a date.
type DateFormat12Choice struct {

	// Date expressed as an ISO Date.
	Date *DateAndDateTimeChoice `xml:"Dt"`

	// Date not specified, for example, the date is unknown.
	NotSpecifiedDate *DateType6Code `xml:"NotSpcfdDt"`
}

func (d *DateFormat12Choice) AddDate() *DateAndDateTimeChoice {
	d.Date = new(DateAndDateTimeChoice)
	return d.Date
}

func (d *DateFormat12Choice) SetNotSpecifiedDate(value string) {
	d.NotSpecifiedDate = (*DateType6Code)(&value)
}
