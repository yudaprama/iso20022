package model

// Provides information about an amended election instruction.
type CorporateActionElection1 struct {

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption1FormatChoice `xml:"OptnTp"`

	// Number identifying the available corporate action options.
	OptionNumber *Exact3NumericText `xml:"OptnNb"`

	// Quantity of the securities that was instructed in the original election advice.
	OriginalInstructedQuantity *UnitOrFaceAmount1Choice `xml:"OrgnlInstdQty"`

	// Remaining instructed securities quantity after the amendment of the election.
	RemainingQuantity *UnitOrFaceAmount1Choice `xml:"RmngQty"`
}

func (c *CorporateActionElection1) AddOptionType() *CorporateActionOption1FormatChoice {
	c.OptionType = new(CorporateActionOption1FormatChoice)
	return c.OptionType
}

func (c *CorporateActionElection1) SetOptionNumber(value string) {
	c.OptionNumber = (*Exact3NumericText)(&value)
}

func (c *CorporateActionElection1) AddOriginalInstructedQuantity() *UnitOrFaceAmount1Choice {
	c.OriginalInstructedQuantity = new(UnitOrFaceAmount1Choice)
	return c.OriginalInstructedQuantity
}

func (c *CorporateActionElection1) AddRemainingQuantity() *UnitOrFaceAmount1Choice {
	c.RemainingQuantity = new(UnitOrFaceAmount1Choice)
	return c.RemainingQuantity
}
