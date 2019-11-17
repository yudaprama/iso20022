package model

// Choice between a date or a code.
type ClosingDate1Choice struct {

	// Closing date is defined as a choice between a date or a date and time format.
	Date *DateAndDateTimeChoice `xml:"Dt"`

	// Closing date is defined using a code or data source scheme.
	Code *Date2Choice `xml:"Cd"`
}

func (c *ClosingDate1Choice) AddDate() *DateAndDateTimeChoice {
	c.Date = new(DateAndDateTimeChoice)
	return c.Date
}

func (c *ClosingDate1Choice) AddCode() *Date2Choice {
	c.Code = new(Date2Choice)
	return c.Code
}
