package model

// Provides information about the corporate action option.
type CorporateActionOption24 struct {

	// Number identifying the available corporate action options.
	OptionNumber *OptionNumber1Choice `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption3Choice `xml:"OptnTp"`

	// Provides information about the securities movement linked to the corporate action option.
	SecuritiesMovementDetails []*SecuritiesOption19 `xml:"SctiesMvmntDtls,omitempty"`

	// Provides information about the cash movement related to a corporate action option.
	CashMovementDetails []*CashOption5 `xml:"CshMvmntDtls,omitempty"`
}

func (c *CorporateActionOption24) AddOptionNumber() *OptionNumber1Choice {
	c.OptionNumber = new(OptionNumber1Choice)
	return c.OptionNumber
}

func (c *CorporateActionOption24) AddOptionType() *CorporateActionOption3Choice {
	c.OptionType = new(CorporateActionOption3Choice)
	return c.OptionType
}

func (c *CorporateActionOption24) AddSecuritiesMovementDetails() *SecuritiesOption19 {
	newValue := new(SecuritiesOption19)
	c.SecuritiesMovementDetails = append(c.SecuritiesMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption24) AddCashMovementDetails() *CashOption5 {
	newValue := new(CashOption5)
	c.CashMovementDetails = append(c.CashMovementDetails, newValue)
	return newValue
}
