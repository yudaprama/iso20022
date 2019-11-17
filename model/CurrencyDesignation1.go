package model

// Information about the designation of a currency.
type CurrencyDesignation1 struct {

	// Specifies whether the currency is settled offshore or onshore.
	CurrencyDesignation *CurrencyDesignation1Code `xml:"CcyDsgnt,omitempty"`

	// Offshore location of the currency.
	Location *CountryCode `xml:"Lctn,omitempty"`

	// Additional information about the off-shore currency.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (c *CurrencyDesignation1) SetCurrencyDesignation(value string) {
	c.CurrencyDesignation = (*CurrencyDesignation1Code)(&value)
}

func (c *CurrencyDesignation1) SetLocation(value string) {
	c.Location = (*CountryCode)(&value)
}

func (c *CurrencyDesignation1) SetAdditionalInformation(value string) {
	c.AdditionalInformation = (*Max350Text)(&value)
}
