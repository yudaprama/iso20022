package model

// Indicates when the amount of money will become available.
type CashAvailabilityDate1Choice struct {

	// Indicates the number of float days attached to the balance.
	NumberOfDays *Max15PlusSignedNumericText `xml:"NbOfDays"`

	// Identifies the actual availability date.
	ActualDate *ISODate `xml:"ActlDt"`
}

func (c *CashAvailabilityDate1Choice) SetNumberOfDays(value string) {
	c.NumberOfDays = (*Max15PlusSignedNumericText)(&value)
}

func (c *CashAvailabilityDate1Choice) SetActualDate(value string) {
	c.ActualDate = (*ISODate)(&value)
}
