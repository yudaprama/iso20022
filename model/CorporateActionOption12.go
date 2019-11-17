package model

// Provides information about the corporate action option.
type CorporateActionOption12 struct {

	// Number identifying the available corporate action options.
	OptionNumber *OptionNumber1Choice `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption3Choice `xml:"OptnTp"`

	// Provides information about the securities movement linked to the corporate action option.
	SecuritiesMovementDetails []*SecuritiesOption5 `xml:"SctiesMvmntDtls,omitempty"`

	// Provides information about the cash movement related to a corporate action option.
	CashMovementDetails []*CashOption5 `xml:"CshMvmntDtls,omitempty"`
}

func (c *CorporateActionOption12) AddOptionNumber() *OptionNumber1Choice {
	c.OptionNumber = new(OptionNumber1Choice)
	return c.OptionNumber
}

func (c *CorporateActionOption12) AddOptionType() *CorporateActionOption3Choice {
	c.OptionType = new(CorporateActionOption3Choice)
	return c.OptionType
}

func (c *CorporateActionOption12) AddSecuritiesMovementDetails() *SecuritiesOption5 {
	newValue := new(SecuritiesOption5)
	c.SecuritiesMovementDetails = append(c.SecuritiesMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption12) AddCashMovementDetails() *CashOption5 {
	newValue := new(CashOption5)
	c.CashMovementDetails = append(c.CashMovementDetails, newValue)
	return newValue
}
