package model

// Amount of money due to a party as compensation for a service.
type Commission19 struct {

	// Commission expressed as an amount of money.
	Amount *ImpliedCurrencyAndAmount `xml:"Amt"`

	// Additional information about the type of commission.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (c *Commission19) SetAmount(value, currency string) {
	c.Amount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *Commission19) SetAdditionalInformation(value string) {
	c.AdditionalInformation = (*Max350Text)(&value)
}
