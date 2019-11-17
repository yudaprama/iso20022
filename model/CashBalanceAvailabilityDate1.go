package model

// Indicates when the amount of money will become available.
type CashBalanceAvailabilityDate1 struct {

	// Indicates the number of float days attached to the balance.
	NumberOfDays *Max15PlusSignedNumericText `xml:"NbOfDays"`

	// Identifies the actual availability date.
	ActualDate *ISODate `xml:"ActlDt"`
}

func (c *CashBalanceAvailabilityDate1) SetNumberOfDays(value string) {
	c.NumberOfDays = (*Max15PlusSignedNumericText)(&value)
}

func (c *CashBalanceAvailabilityDate1) SetActualDate(value string) {
	c.ActualDate = (*ISODate)(&value)
}
