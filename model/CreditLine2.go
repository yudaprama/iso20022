package model

// Set of elements used to provide details of the credit line.
type CreditLine2 struct {

	// Indicates whether or not the credit line is included in the balance.
	//
	// Usage: If not present, credit line is not included in the balance amount.
	Included *TrueFalseIndicator `xml:"Incl"`

	// Amount of money of the credit line.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt,omitempty"`
}

func (c *CreditLine2) SetIncluded(value string) {
	c.Included = (*TrueFalseIndicator)(&value)
}

func (c *CreditLine2) SetAmount(value, currency string) {
	c.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}
