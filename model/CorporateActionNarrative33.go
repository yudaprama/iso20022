package model

// Provides additional information such as the information to comply with.
type CorporateActionNarrative33 struct {

	// Provides information conditions to the account owner that are to be complied with, for example, not open to US/Canadian residents, Qualified Institutional Buyers (QIB) or Sophisticated Investor Letter (SIL) to be provided.
	InformationToComplyWith []*RestrictedFINXMax350Text `xml:"InfToCmplyWth,omitempty"`

	// Provides additional information on the delivery details of the outturned (derived) securities. This narrative is only to be used in case the securities are not eligible at the agent/custodian, and may not be used for settlement instructions.
	DeliveryDetails []*RestrictedFINXMax350Text `xml:"DlvryDtls,omitempty"`

	// Provides additional details pertaining to foreign exchange instructions.
	ForeignExchangeInstructionsAdditionalInformation []*RestrictedFINXMax350Text `xml:"FXInstrsAddtlInf,omitempty"`

	// Provides additional details pertaining to the corporate action instruction.
	InstructionAdditionalInformation []*RestrictedFINXMax350Text `xml:"InstrAddtlInf,omitempty"`
}

func (c *CorporateActionNarrative33) AddInformationToComplyWith(value string) {
	c.InformationToComplyWith = append(c.InformationToComplyWith, (*RestrictedFINXMax350Text)(&value))
}

func (c *CorporateActionNarrative33) AddDeliveryDetails(value string) {
	c.DeliveryDetails = append(c.DeliveryDetails, (*RestrictedFINXMax350Text)(&value))
}

func (c *CorporateActionNarrative33) AddForeignExchangeInstructionsAdditionalInformation(value string) {
	c.ForeignExchangeInstructionsAdditionalInformation = append(c.ForeignExchangeInstructionsAdditionalInformation, (*RestrictedFINXMax350Text)(&value))
}

func (c *CorporateActionNarrative33) AddInstructionAdditionalInformation(value string) {
	c.InstructionAdditionalInformation = append(c.InstructionAdditionalInformation, (*RestrictedFINXMax350Text)(&value))
}
