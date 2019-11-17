package model

// Provides additional information about the CA event.
type CorporateActionNarrative2 struct {

	// Provides conditional information related to the event, eg, an offer is subject to 50% acceptance, the offeror allows the securities holder to set some conditions.
	InformationConditions *Max350Text `xml:"InfConds,omitempty"`

	// Provides information conditions to the account owner that are to be complied with, eg, not open to US/Canadian residents, QIB or SIL to be provided.
	InformationToComplyWith *Max350Text `xml:"InfToCmplyWth,omitempty"`

	// Provides taxation conditions that cannot be included within the structured fields of this message and has not been mentioned in the SLA.
	TaxationConditions *Max350Text `xml:"TaxtnConds,omitempty"`

	// Provides declaration details narrative relative to the financial instrument, eg, beneficial ownership.
	DeclarationDetails *Max350Text `xml:"DclrtnDtls,omitempty"`

	// Provides information required for the registration.
	RegistrationDetails *Max350Text `xml:"RegnDtls,omitempty"`

	// Provides additional information or specifies in more detail the content of a message.
	AdditionalText *Max350Text `xml:"AddtlTxt,omitempty"`
}

func (c *CorporateActionNarrative2) SetInformationConditions(value string) {
	c.InformationConditions = (*Max350Text)(&value)
}

func (c *CorporateActionNarrative2) SetInformationToComplyWith(value string) {
	c.InformationToComplyWith = (*Max350Text)(&value)
}

func (c *CorporateActionNarrative2) SetTaxationConditions(value string) {
	c.TaxationConditions = (*Max350Text)(&value)
}

func (c *CorporateActionNarrative2) SetDeclarationDetails(value string) {
	c.DeclarationDetails = (*Max350Text)(&value)
}

func (c *CorporateActionNarrative2) SetRegistrationDetails(value string) {
	c.RegistrationDetails = (*Max350Text)(&value)
}

func (c *CorporateActionNarrative2) SetAdditionalText(value string) {
	c.AdditionalText = (*Max350Text)(&value)
}
