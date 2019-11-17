package model

// Provides status of the deactivation instruction.
type CorporateActionDeactivationInstructionStatus1 struct {

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption1FormatChoice `xml:"OptnTp,omitempty"`

	// Number identifying the available corporate action options.
	OptionNumber *Exact3NumericText `xml:"OptnNb,omitempty"`

	// Provides information about the processing status of the instruction.
	ProcessedStatus *CorporateActionDeactivationInstructionProcessingStatus1 `xml:"PrcdSts"`

	// Provides information about the rejection status.
	RejectedStatus *CorporateActionDeactivationInstructionRejectionStatus1 `xml:"RjctdSts"`
}

func (c *CorporateActionDeactivationInstructionStatus1) AddOptionType() *CorporateActionOption1FormatChoice {
	c.OptionType = new(CorporateActionOption1FormatChoice)
	return c.OptionType
}

func (c *CorporateActionDeactivationInstructionStatus1) SetOptionNumber(value string) {
	c.OptionNumber = (*Exact3NumericText)(&value)
}

func (c *CorporateActionDeactivationInstructionStatus1) AddProcessedStatus() *CorporateActionDeactivationInstructionProcessingStatus1 {
	c.ProcessedStatus = new(CorporateActionDeactivationInstructionProcessingStatus1)
	return c.ProcessedStatus
}

func (c *CorporateActionDeactivationInstructionStatus1) AddRejectedStatus() *CorporateActionDeactivationInstructionRejectionStatus1 {
	c.RejectedStatus = new(CorporateActionDeactivationInstructionRejectionStatus1)
	return c.RejectedStatus
}
