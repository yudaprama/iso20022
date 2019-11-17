package model

// Provides information about a corporate action election amendment request.
type CorporateActionElection2 struct {

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption1FormatChoice `xml:"OptnTp"`

	// Number identifying the available corporate action options.
	OptionNumber *Exact3NumericText `xml:"OptnNb"`

	// New instructed securities quantity after the amendment.
	NewInstructedQuantity *UnitOrFaceAmount1Choice `xml:"NewInstdQty"`

	// The reason for the amendment request.
	Reason *Max350Text `xml:"Rsn,omitempty"`
}

func (c *CorporateActionElection2) AddOptionType() *CorporateActionOption1FormatChoice {
	c.OptionType = new(CorporateActionOption1FormatChoice)
	return c.OptionType
}

func (c *CorporateActionElection2) SetOptionNumber(value string) {
	c.OptionNumber = (*Exact3NumericText)(&value)
}

func (c *CorporateActionElection2) AddNewInstructedQuantity() *UnitOrFaceAmount1Choice {
	c.NewInstructedQuantity = new(UnitOrFaceAmount1Choice)
	return c.NewInstructedQuantity
}

func (c *CorporateActionElection2) SetReason(value string) {
	c.Reason = (*Max350Text)(&value)
}
