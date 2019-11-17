package model

// Provides additional information such as the taxation conditions.
type CorporateActionNarrative27 struct {

	// Provides additional information or specifies in more detail the content of a message.  This field may only be used when the information to be transmitted, cannot be coded.
	AdditionalText *UpdatedAdditionalInformation8 `xml:"AddtlTxt,omitempty"`

	// Provides information that can be ignored for automated processing; - reiteration of information that has been included within structured fields, - or narrative information not needed for automatic processing.
	NarrativeVersion *UpdatedAdditionalInformation8 `xml:"NrrtvVrsn,omitempty"`

	// Provides conditional information related to the event, for example, an offer is subject to 50 percent acceptance, the offeror allows the securities holder to set some conditions.
	InformationConditions *UpdatedAdditionalInformation8 `xml:"InfConds,omitempty"`

	// Provides information conditions to the account owner that are to be complied with, for example, not open to US/Canadian residents, Qualified Institutional Buyers (QIB) or SIL (Sophisticated Investor Letter)  to be provided.
	InformationToComplyWith *UpdatedAdditionalInformation8 `xml:"InfToCmplyWth,omitempty"`

	// Provides taxation conditions that cannot be included within the structured fields of this message and has not been mentioned in the Service Level Agreement (SLA).
	TaxationConditions *UpdatedAdditionalInformation8 `xml:"TaxtnConds,omitempty"`

	// Provides a disclaimer relative to the information provided in the message. It may be ignored for automated processing.
	Disclaimer *UpdatedAdditionalInformation8 `xml:"Dsclmr,omitempty"`

	// Provides additional information regarding the party, for example, the contact unit or person responsible for the transaction identified in the message.
	PartyContactNarrative *UpdatedAdditionalInformation8 `xml:"PtyCtctNrrtv,omitempty"`

	// Provides information required for the registration.
	RegistrationDetails *UpdatedAdditionalInformation8 `xml:"RegnDtls,omitempty"`

	// Provides additional information on the basket or index underlying a security, for example a warrant.
	BasketOrIndexInformation *UpdatedAdditionalInformation8 `xml:"BsktOrIndxInf,omitempty"`

	// Provides additional information about the type of certification/breakdown required.
	CertificationBreakdown *UpdatedAdditionalInformation8 `xml:"CertfctnBrkdwn,omitempty"`
}

func (c *CorporateActionNarrative27) AddAdditionalText() *UpdatedAdditionalInformation8 {
	c.AdditionalText = new(UpdatedAdditionalInformation8)
	return c.AdditionalText
}

func (c *CorporateActionNarrative27) AddNarrativeVersion() *UpdatedAdditionalInformation8 {
	c.NarrativeVersion = new(UpdatedAdditionalInformation8)
	return c.NarrativeVersion
}

func (c *CorporateActionNarrative27) AddInformationConditions() *UpdatedAdditionalInformation8 {
	c.InformationConditions = new(UpdatedAdditionalInformation8)
	return c.InformationConditions
}

func (c *CorporateActionNarrative27) AddInformationToComplyWith() *UpdatedAdditionalInformation8 {
	c.InformationToComplyWith = new(UpdatedAdditionalInformation8)
	return c.InformationToComplyWith
}

func (c *CorporateActionNarrative27) AddTaxationConditions() *UpdatedAdditionalInformation8 {
	c.TaxationConditions = new(UpdatedAdditionalInformation8)
	return c.TaxationConditions
}

func (c *CorporateActionNarrative27) AddDisclaimer() *UpdatedAdditionalInformation8 {
	c.Disclaimer = new(UpdatedAdditionalInformation8)
	return c.Disclaimer
}

func (c *CorporateActionNarrative27) AddPartyContactNarrative() *UpdatedAdditionalInformation8 {
	c.PartyContactNarrative = new(UpdatedAdditionalInformation8)
	return c.PartyContactNarrative
}

func (c *CorporateActionNarrative27) AddRegistrationDetails() *UpdatedAdditionalInformation8 {
	c.RegistrationDetails = new(UpdatedAdditionalInformation8)
	return c.RegistrationDetails
}

func (c *CorporateActionNarrative27) AddBasketOrIndexInformation() *UpdatedAdditionalInformation8 {
	c.BasketOrIndexInformation = new(UpdatedAdditionalInformation8)
	return c.BasketOrIndexInformation
}

func (c *CorporateActionNarrative27) AddCertificationBreakdown() *UpdatedAdditionalInformation8 {
	c.CertificationBreakdown = new(UpdatedAdditionalInformation8)
	return c.CertificationBreakdown
}
