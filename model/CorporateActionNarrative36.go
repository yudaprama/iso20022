package model

// Provides additional information such as the taxation conditions.
type CorporateActionNarrative36 struct {

	// Provides additional information or specifies in more detail the content of a message. This field may only be used when the information to be transmitted, cannot be coded.
	AdditionalText *UpdatedAdditionalInformation6 `xml:"AddtlTxt,omitempty"`

	// Provides information that can be ignored for automated processing; - reiteration of information that has been included within structured fields, - or narrative information not needed for automatic processing.
	NarrativeVersion *UpdatedAdditionalInformation6 `xml:"NrrtvVrsn,omitempty"`

	// Provides conditional information related to the event, for example, an offer is subject to 50 percent acceptance, the offeror allows the securities holder to set some conditions.
	InformationConditions *UpdatedAdditionalInformation5 `xml:"InfConds,omitempty"`

	// Provides information conditions to the account owner that are to be complied with, for example, not open to US/Canadian residents, Qualified Institutional Buyers (QIB) or SIL (Sophisticated Investor Letter)  to be provided.
	InformationToComplyWith *UpdatedAdditionalInformation5 `xml:"InfToCmplyWth,omitempty"`

	// Provides restriction(s) on securities.
	SecurityRestriction *UpdatedAdditionalInformation5 `xml:"SctyRstrctn,omitempty"`

	// Provides taxation conditions that cannot be included within the structured fields of this message  and has not been mentioned in the Service Level Agreement (SLA).
	TaxationConditions *UpdatedAdditionalInformation5 `xml:"TaxtnConds,omitempty"`

	// Provides a disclaimer relative to the information provided in the message. It may be ignored for automated processing.
	Disclaimer *UpdatedAdditionalInformation5 `xml:"Dsclmr,omitempty"`

	// Provides additional information about the type of certification/breakdown required.
	CertificationBreakdown *UpdatedAdditionalInformation5 `xml:"CertfctnBrkdwn,omitempty"`
}

func (c *CorporateActionNarrative36) AddAdditionalText() *UpdatedAdditionalInformation6 {
	c.AdditionalText = new(UpdatedAdditionalInformation6)
	return c.AdditionalText
}

func (c *CorporateActionNarrative36) AddNarrativeVersion() *UpdatedAdditionalInformation6 {
	c.NarrativeVersion = new(UpdatedAdditionalInformation6)
	return c.NarrativeVersion
}

func (c *CorporateActionNarrative36) AddInformationConditions() *UpdatedAdditionalInformation5 {
	c.InformationConditions = new(UpdatedAdditionalInformation5)
	return c.InformationConditions
}

func (c *CorporateActionNarrative36) AddInformationToComplyWith() *UpdatedAdditionalInformation5 {
	c.InformationToComplyWith = new(UpdatedAdditionalInformation5)
	return c.InformationToComplyWith
}

func (c *CorporateActionNarrative36) AddSecurityRestriction() *UpdatedAdditionalInformation5 {
	c.SecurityRestriction = new(UpdatedAdditionalInformation5)
	return c.SecurityRestriction
}

func (c *CorporateActionNarrative36) AddTaxationConditions() *UpdatedAdditionalInformation5 {
	c.TaxationConditions = new(UpdatedAdditionalInformation5)
	return c.TaxationConditions
}

func (c *CorporateActionNarrative36) AddDisclaimer() *UpdatedAdditionalInformation5 {
	c.Disclaimer = new(UpdatedAdditionalInformation5)
	return c.Disclaimer
}

func (c *CorporateActionNarrative36) AddCertificationBreakdown() *UpdatedAdditionalInformation5 {
	c.CertificationBreakdown = new(UpdatedAdditionalInformation5)
	return c.CertificationBreakdown
}
