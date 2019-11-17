package model

// Provides additional information such as the taxation conditions.
type CorporateActionNarrative41 struct {

	// Provides additional information or specifies in more detail the content of a message.  This field may only be used when the information to be transmitted, cannot be coded.
	AdditionalText *UpdatedAdditionalInformation10 `xml:"AddtlTxt,omitempty"`

	// Provides information that can be ignored for automated processing; - reiteration of information that has been included within structured fields, - or narrative information not needed for automatic processing.
	NarrativeVersion *UpdatedAdditionalInformation10 `xml:"NrrtvVrsn,omitempty"`

	// Provides conditional information related to the event, for example, an offer is subject to 50 percent acceptance, the offeror allows the securities holder to set some conditions.
	InformationConditions *UpdatedAdditionalInformation10 `xml:"InfConds,omitempty"`

	// Provides information conditions to the account owner that are to be complied with, for example, not open to US/Canadian residents, Qualified Institutional Buyers (QIB) or SIL (Sophisticated Investor Letter)  to be provided.
	InformationToComplyWith *UpdatedAdditionalInformation10 `xml:"InfToCmplyWth,omitempty"`

	// Provides taxation conditions that cannot be included within the structured fields of this message and has not been mentioned in the Service Level Agreement (SLA).
	TaxationConditions *UpdatedAdditionalInformation10 `xml:"TaxtnConds,omitempty"`

	// Provides a disclaimer relative to the information provided in the message. It may be ignored for automated processing.
	Disclaimer *UpdatedAdditionalInformation10 `xml:"Dsclmr,omitempty"`

	// Provides additional information regarding the party, for example, the contact unit or person responsible for the transaction identified in the message.
	PartyContactNarrative *UpdatedAdditionalInformation10 `xml:"PtyCtctNrrtv,omitempty"`

	// Provides information required for the registration.
	RegistrationDetails *UpdatedAdditionalInformation10 `xml:"RegnDtls,omitempty"`

	// Provides additional information on the basket or index underlying a security, for example a warrant.
	BasketOrIndexInformation *UpdatedAdditionalInformation10 `xml:"BsktOrIndxInf,omitempty"`

	// Provides additional information about the type of certification/breakdown required.
	CertificationBreakdown *UpdatedAdditionalInformation10 `xml:"CertfctnBrkdwn,omitempty"`
}

func (c *CorporateActionNarrative41) AddAdditionalText() *UpdatedAdditionalInformation10 {
	c.AdditionalText = new(UpdatedAdditionalInformation10)
	return c.AdditionalText
}

func (c *CorporateActionNarrative41) AddNarrativeVersion() *UpdatedAdditionalInformation10 {
	c.NarrativeVersion = new(UpdatedAdditionalInformation10)
	return c.NarrativeVersion
}

func (c *CorporateActionNarrative41) AddInformationConditions() *UpdatedAdditionalInformation10 {
	c.InformationConditions = new(UpdatedAdditionalInformation10)
	return c.InformationConditions
}

func (c *CorporateActionNarrative41) AddInformationToComplyWith() *UpdatedAdditionalInformation10 {
	c.InformationToComplyWith = new(UpdatedAdditionalInformation10)
	return c.InformationToComplyWith
}

func (c *CorporateActionNarrative41) AddTaxationConditions() *UpdatedAdditionalInformation10 {
	c.TaxationConditions = new(UpdatedAdditionalInformation10)
	return c.TaxationConditions
}

func (c *CorporateActionNarrative41) AddDisclaimer() *UpdatedAdditionalInformation10 {
	c.Disclaimer = new(UpdatedAdditionalInformation10)
	return c.Disclaimer
}

func (c *CorporateActionNarrative41) AddPartyContactNarrative() *UpdatedAdditionalInformation10 {
	c.PartyContactNarrative = new(UpdatedAdditionalInformation10)
	return c.PartyContactNarrative
}

func (c *CorporateActionNarrative41) AddRegistrationDetails() *UpdatedAdditionalInformation10 {
	c.RegistrationDetails = new(UpdatedAdditionalInformation10)
	return c.RegistrationDetails
}

func (c *CorporateActionNarrative41) AddBasketOrIndexInformation() *UpdatedAdditionalInformation10 {
	c.BasketOrIndexInformation = new(UpdatedAdditionalInformation10)
	return c.BasketOrIndexInformation
}

func (c *CorporateActionNarrative41) AddCertificationBreakdown() *UpdatedAdditionalInformation10 {
	c.CertificationBreakdown = new(UpdatedAdditionalInformation10)
	return c.CertificationBreakdown
}
