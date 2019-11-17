package model

// Provides information about the corporate action option.
type CorporateActionOption128 struct {

	// Number identifying the available corporate action options.
	OptionNumber *OptionNumber1Choice `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption29Choice `xml:"OptnTp"`

	// Quantity of securities to which this instruction applies.
	InstructedQuantity *Quantity40Choice `xml:"InstdQty"`
}

func (c *CorporateActionOption128) AddOptionNumber() *OptionNumber1Choice {
	c.OptionNumber = new(OptionNumber1Choice)
	return c.OptionNumber
}

func (c *CorporateActionOption128) AddOptionType() *CorporateActionOption29Choice {
	c.OptionType = new(CorporateActionOption29Choice)
	return c.OptionType
}

func (c *CorporateActionOption128) AddInstructedQuantity() *Quantity40Choice {
	c.InstructedQuantity = new(Quantity40Choice)
	return c.InstructedQuantity
}
