package model

// Provides information about the corporate action option.
type CorporateActionOption120 struct {

	// Number identifying the available corporate action options.
	OptionNumber *OptionNumber1Choice `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption20Choice `xml:"OptnTp"`

	// Quantity of securities to which this instruction applies.
	InstructedQuantity *Quantity20Choice `xml:"InstdQty"`
}

func (c *CorporateActionOption120) AddOptionNumber() *OptionNumber1Choice {
	c.OptionNumber = new(OptionNumber1Choice)
	return c.OptionNumber
}

func (c *CorporateActionOption120) AddOptionType() *CorporateActionOption20Choice {
	c.OptionType = new(CorporateActionOption20Choice)
	return c.OptionType
}

func (c *CorporateActionOption120) AddInstructedQuantity() *Quantity20Choice {
	c.InstructedQuantity = new(Quantity20Choice)
	return c.InstructedQuantity
}
