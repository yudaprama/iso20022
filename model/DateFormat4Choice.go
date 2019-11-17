package model

// Specifies the value of a date.
type DateFormat4Choice struct {

	// Date expressed as an ISO Date.
	Date *DateAndDateTimeChoice `xml:"Dt"`

	// The date is not specified, eg, the date is unknown.
	NotSpecifiedDate *DateType6Code `xml:"NotSpcfdDt"`

	// Proprietary scheme to specify a date.
	Proprietary *GenericIdentification13 `xml:"Prtry"`
}

func (d *DateFormat4Choice) AddDate() *DateAndDateTimeChoice {
	d.Date = new(DateAndDateTimeChoice)
	return d.Date
}

func (d *DateFormat4Choice) SetNotSpecifiedDate(value string) {
	d.NotSpecifiedDate = (*DateType6Code)(&value)
}

func (d *DateFormat4Choice) AddProprietary() *GenericIdentification13 {
	d.Proprietary = new(GenericIdentification13)
	return d.Proprietary
}
