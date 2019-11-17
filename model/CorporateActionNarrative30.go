package model

// Provides additional information such as the registration details.
type CorporateActionNarrative30 struct {

	// Provides information required for the registration.
	RegistrationDetails []*Max350Text `xml:"RegnDtls,omitempty"`

	// Provides additional information regarding the party, for example, the contact unit or person responsible for the transaction identified in the message.
	PartyContactNarrative []*Max350Text `xml:"PtyCtctNrrtv,omitempty"`

	// Provides information required for the certification/breakdown.
	CertificationBreakdown []*Max350Text `xml:"CertfctnBrkdwn,omitempty"`
}

func (c *CorporateActionNarrative30) AddRegistrationDetails(value string) {
	c.RegistrationDetails = append(c.RegistrationDetails, (*Max350Text)(&value))
}

func (c *CorporateActionNarrative30) AddPartyContactNarrative(value string) {
	c.PartyContactNarrative = append(c.PartyContactNarrative, (*Max350Text)(&value))
}

func (c *CorporateActionNarrative30) AddCertificationBreakdown(value string) {
	c.CertificationBreakdown = append(c.CertificationBreakdown, (*Max350Text)(&value))
}
