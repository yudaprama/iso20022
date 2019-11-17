package model

// Specification of the cash account.
type CashAccountType1 struct {

	// Structured format.
	Structured *FundCashAccount1Code `xml:"Strd"`

	// Additional information about the type of tax.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (c *CashAccountType1) SetStructured(value string) {
	c.Structured = (*FundCashAccount1Code)(&value)
}

func (c *CashAccountType1) SetAdditionalInformation(value string) {
	c.AdditionalInformation = (*Max350Text)(&value)
}
