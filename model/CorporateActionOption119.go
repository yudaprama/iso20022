package model

// Provides information about the corporate action option.
type CorporateActionOption119 struct {

	// Number identifying the available corporate action options.
	OptionNumber *OptionNumber1Choice `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption19Choice `xml:"OptnTp"`

	// Provides information about the securities movement linked to the corporate action option.
	SecuritiesMovementDetails []*SecuritiesOption53 `xml:"SctiesMvmntDtls,omitempty"`

	// Provides information about the cash movement related to a corporate action option.
	CashMovementDetails []*CashOption45 `xml:"CshMvmntDtls,omitempty"`
}

func (c *CorporateActionOption119) AddOptionNumber() *OptionNumber1Choice {
	c.OptionNumber = new(OptionNumber1Choice)
	return c.OptionNumber
}

func (c *CorporateActionOption119) AddOptionType() *CorporateActionOption19Choice {
	c.OptionType = new(CorporateActionOption19Choice)
	return c.OptionType
}

func (c *CorporateActionOption119) AddSecuritiesMovementDetails() *SecuritiesOption53 {
	newValue := new(SecuritiesOption53)
	c.SecuritiesMovementDetails = append(c.SecuritiesMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption119) AddCashMovementDetails() *CashOption45 {
	newValue := new(CashOption45)
	c.CashMovementDetails = append(c.CashMovementDetails, newValue)
	return newValue
}
