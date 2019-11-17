package model

// Provides information about the corporate action option.
type CorporateActionOption39 struct {

	// Number identifying the available corporate action options.
	OptionNumber *OptionNumber1Choice `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption11Choice `xml:"OptnTp"`

	// Provides information about the securities movement linked to the corporate action option.
	SecuritiesMovementDetails []*SecuritiesOption27 `xml:"SctiesMvmntDtls,omitempty"`

	// Provides information about the cash movement related to a corporate action option.
	CashMovementDetails []*CashOption19 `xml:"CshMvmntDtls,omitempty"`
}

func (c *CorporateActionOption39) AddOptionNumber() *OptionNumber1Choice {
	c.OptionNumber = new(OptionNumber1Choice)
	return c.OptionNumber
}

func (c *CorporateActionOption39) AddOptionType() *CorporateActionOption11Choice {
	c.OptionType = new(CorporateActionOption11Choice)
	return c.OptionType
}

func (c *CorporateActionOption39) AddSecuritiesMovementDetails() *SecuritiesOption27 {
	newValue := new(SecuritiesOption27)
	c.SecuritiesMovementDetails = append(c.SecuritiesMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption39) AddCashMovementDetails() *CashOption19 {
	newValue := new(CashOption19)
	c.CashMovementDetails = append(c.CashMovementDetails, newValue)
	return newValue
}
