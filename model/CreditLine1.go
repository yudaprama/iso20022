package model

// Transaction has an origin and a destination in the same country and is made in the currency of that country.
type CreditLine1 struct {

	// Indicates whether the credit line is included or not.
	//
	// Usage : if not present, credit line is not included in the balance amount.
	Included *TrueFalseIndicator `xml:"Incl"`

	// Amount of money of the credit line.
	Amount *CurrencyAndAmount `xml:"Amt,omitempty"`
}

func (c *CreditLine1) SetIncluded(value string) {
	c.Included = (*TrueFalseIndicator)(&value)
}

func (c *CreditLine1) SetAmount(value, currency string) {
	c.Amount = NewCurrencyAndAmount(value, currency)
}
