package model

// Provides information about the corporate action option.
type CorporateActionOption126 struct {

	// Number identifying the available corporate action options.
	OptionNumber *OptionNumber1Choice `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption26Choice `xml:"OptnTp"`

	// Provides information about the securities movement linked to the corporate action option.
	SecuritiesMovementDetails []*SecuritiesOption58 `xml:"SctiesMvmntDtls,omitempty"`

	// Provides information about the cash movement related to a corporate action option.
	CashMovementDetails []*CashOption48 `xml:"CshMvmntDtls,omitempty"`
}

func (c *CorporateActionOption126) AddOptionNumber() *OptionNumber1Choice {
	c.OptionNumber = new(OptionNumber1Choice)
	return c.OptionNumber
}

func (c *CorporateActionOption126) AddOptionType() *CorporateActionOption26Choice {
	c.OptionType = new(CorporateActionOption26Choice)
	return c.OptionType
}

func (c *CorporateActionOption126) AddSecuritiesMovementDetails() *SecuritiesOption58 {
	newValue := new(SecuritiesOption58)
	c.SecuritiesMovementDetails = append(c.SecuritiesMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption126) AddCashMovementDetails() *CashOption48 {
	newValue := new(CashOption48)
	c.CashMovementDetails = append(c.CashMovementDetails, newValue)
	return newValue
}
