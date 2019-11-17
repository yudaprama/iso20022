package model

// Provides additional information such as the information to comply with.
type CorporateActionNarrative8 struct {

	// Provides information conditions to the account owner that are to be complied with, for example, not open to US/Canadian residents, Qualified Institutional Buyers (QIB) or Sophisticated Investor Letter (SIL) to be provided.
	InformationToComplyWith []*Max350Text `xml:"InfToCmplyWth,omitempty"`

	// Provides additional information on the delivery details of the outturned (derived) securities. This narrative is only to be used in case the securities are not eligible at the agent/custodian, and may not be used for settlement instructions.
	DeliveryDetails []*Max350Text `xml:"DlvryDtls,omitempty"`

	// Provides additional details pertaining to foreign exchange instructions.
	ForeignExchangeInstructionsAdditionalInformation []*Max350Text `xml:"FXInstrsAddtlInf,omitempty"`

	// Disclaimer relative to the information provided in the message. It may be ignored for automated processing. No information about the instruction itself is allowed here.
	Disclaimer []*Max350Text `xml:"Dsclmr,omitempty"`

	// Provides additional details pertaining to the corporate action instruction.
	InstructionAdditionalInformation []*Max350Text `xml:"InstrAddtlInf,omitempty"`
}

func (c *CorporateActionNarrative8) AddInformationToComplyWith(value string) {
	c.InformationToComplyWith = append(c.InformationToComplyWith, (*Max350Text)(&value))
}

func (c *CorporateActionNarrative8) AddDeliveryDetails(value string) {
	c.DeliveryDetails = append(c.DeliveryDetails, (*Max350Text)(&value))
}

func (c *CorporateActionNarrative8) AddForeignExchangeInstructionsAdditionalInformation(value string) {
	c.ForeignExchangeInstructionsAdditionalInformation = append(c.ForeignExchangeInstructionsAdditionalInformation, (*Max350Text)(&value))
}

func (c *CorporateActionNarrative8) AddDisclaimer(value string) {
	c.Disclaimer = append(c.Disclaimer, (*Max350Text)(&value))
}

func (c *CorporateActionNarrative8) AddInstructionAdditionalInformation(value string) {
	c.InstructionAdditionalInformation = append(c.InstructionAdditionalInformation, (*Max350Text)(&value))
}
