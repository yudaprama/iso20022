package model

// Provides additional information such as the taxation conditions.
type CorporateActionNarrative11 struct {

	// Provides additional information or specifies in more detail the content of a message.  This field may only be used when the information to be transmitted, cannot be coded.
	AdditionalText *UpdatedAdditionalInformation2 `xml:"AddtlTxt,omitempty"`

	// Provides information that can be ignored for automated processing; - reiteration of information that has been included within structured fields of this message, - or narrative information not needed for automatic processing.
	NarrativeVersion *UpdatedAdditionalInformation2 `xml:"NrrtvVrsn,omitempty"`

	// Provides conditional information related to the event, for example, an offer is subject to 50 percent acceptance, the offeror allows the securities holder to set some conditions.
	InformationConditions *UpdatedAdditionalInformation2 `xml:"InfConds,omitempty"`

	// Provides information conditions to the account owner that are to be complied with, for example, not open to US/Canadian residents, Qualified Institutional Buyers (QIB) or SIL (Sophisticated Investor Letter)  to be provided.
	InformationToComplyWith *UpdatedAdditionalInformation2 `xml:"InfToCmplyWth,omitempty"`

	// Provides taxation conditions that cannot be included within the structured fields of this message and has not been mentioned in the Service Level Agreement (SLA).
	TaxationConditions *UpdatedAdditionalInformation2 `xml:"TaxtnConds,omitempty"`

	// Provides a disclaimer relative to the information provided in the message. It may be ignored for automated processing.
	Disclaimer *UpdatedAdditionalInformation2 `xml:"Dsclmr,omitempty"`

	// Provides additional information regarding the party, for example, the contact unit or person responsible for the transaction identified in the message.
	PartyContactNarrative *UpdatedAdditionalInformation2 `xml:"PtyCtctNrrtv,omitempty"`

	// Provides declaration details narrative relative to the financial instrument, for example, beneficial ownership.
	DeclarationDetails *UpdatedAdditionalInformation2 `xml:"DclrtnDtls,omitempty"`

	// Provides information required for the registration.
	RegistrationDetails *UpdatedAdditionalInformation2 `xml:"RegnDtls,omitempty"`

	// Provides additional information on the basket or index underlying a security, for example a warrant.
	BasketOrIndexInformation *UpdatedAdditionalInformation2 `xml:"BsktOrIndxInf,omitempty"`
}

func (c *CorporateActionNarrative11) AddAdditionalText() *UpdatedAdditionalInformation2 {
	c.AdditionalText = new(UpdatedAdditionalInformation2)
	return c.AdditionalText
}

func (c *CorporateActionNarrative11) AddNarrativeVersion() *UpdatedAdditionalInformation2 {
	c.NarrativeVersion = new(UpdatedAdditionalInformation2)
	return c.NarrativeVersion
}

func (c *CorporateActionNarrative11) AddInformationConditions() *UpdatedAdditionalInformation2 {
	c.InformationConditions = new(UpdatedAdditionalInformation2)
	return c.InformationConditions
}

func (c *CorporateActionNarrative11) AddInformationToComplyWith() *UpdatedAdditionalInformation2 {
	c.InformationToComplyWith = new(UpdatedAdditionalInformation2)
	return c.InformationToComplyWith
}

func (c *CorporateActionNarrative11) AddTaxationConditions() *UpdatedAdditionalInformation2 {
	c.TaxationConditions = new(UpdatedAdditionalInformation2)
	return c.TaxationConditions
}

func (c *CorporateActionNarrative11) AddDisclaimer() *UpdatedAdditionalInformation2 {
	c.Disclaimer = new(UpdatedAdditionalInformation2)
	return c.Disclaimer
}

func (c *CorporateActionNarrative11) AddPartyContactNarrative() *UpdatedAdditionalInformation2 {
	c.PartyContactNarrative = new(UpdatedAdditionalInformation2)
	return c.PartyContactNarrative
}

func (c *CorporateActionNarrative11) AddDeclarationDetails() *UpdatedAdditionalInformation2 {
	c.DeclarationDetails = new(UpdatedAdditionalInformation2)
	return c.DeclarationDetails
}

func (c *CorporateActionNarrative11) AddRegistrationDetails() *UpdatedAdditionalInformation2 {
	c.RegistrationDetails = new(UpdatedAdditionalInformation2)
	return c.RegistrationDetails
}

func (c *CorporateActionNarrative11) AddBasketOrIndexInformation() *UpdatedAdditionalInformation2 {
	c.BasketOrIndexInformation = new(UpdatedAdditionalInformation2)
	return c.BasketOrIndexInformation
}
