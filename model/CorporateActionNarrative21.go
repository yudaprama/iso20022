package model

// Provides additional information such as the registration details.
type CorporateActionNarrative21 struct {

	// Provides additional information or specifies in more detail the content of a message. This field may only be used when the information to be transmitted, cannot be coded.
	AdditionalText []*Max350Text `xml:"AddtlTxt,omitempty"`

	// Provides information that can be ignored for automated processing; - reiteration of information that has been included within structured fields of this message, - or narrative information not needed for automatic processing.
	NarrativeVersion []*Max350Text `xml:"NrrtvVrsn,omitempty"`

	// Provides information required for the registration.
	RegistrationDetails []*Max350Text `xml:"RegnDtls,omitempty"`

	// Provides additional information regarding the party, for example, the contact unit or person responsible for the transaction identified in the message.
	PartyContactNarrative []*Max350Text `xml:"PtyCtctNrrtv,omitempty"`

	// Disclaimer relative to the information provided in the message. It may be ignored for automated processing. No information about the instruction itself is allowed here.
	Disclaimer []*Max350Text `xml:"Dsclmr,omitempty"`

	// Provides additional information on the basket or index underlying a security, for example a warrant.
	BasketOrIndexInformation []*Max350Text `xml:"BsktOrIndxInf,omitempty"`

	// Provides information required for the certification/breakdown.
	CertificationBreakdown []*Max350Text `xml:"CertfctnBrkdwn,omitempty"`
}

func (c *CorporateActionNarrative21) AddAdditionalText(value string) {
	c.AdditionalText = append(c.AdditionalText, (*Max350Text)(&value))
}

func (c *CorporateActionNarrative21) AddNarrativeVersion(value string) {
	c.NarrativeVersion = append(c.NarrativeVersion, (*Max350Text)(&value))
}

func (c *CorporateActionNarrative21) AddRegistrationDetails(value string) {
	c.RegistrationDetails = append(c.RegistrationDetails, (*Max350Text)(&value))
}

func (c *CorporateActionNarrative21) AddPartyContactNarrative(value string) {
	c.PartyContactNarrative = append(c.PartyContactNarrative, (*Max350Text)(&value))
}

func (c *CorporateActionNarrative21) AddDisclaimer(value string) {
	c.Disclaimer = append(c.Disclaimer, (*Max350Text)(&value))
}

func (c *CorporateActionNarrative21) AddBasketOrIndexInformation(value string) {
	c.BasketOrIndexInformation = append(c.BasketOrIndexInformation, (*Max350Text)(&value))
}

func (c *CorporateActionNarrative21) AddCertificationBreakdown(value string) {
	c.CertificationBreakdown = append(c.CertificationBreakdown, (*Max350Text)(&value))
}
