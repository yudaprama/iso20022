package model

// Provides additional information such as the registration details.
type CorporateActionNarrative34 struct {

	// Provides information required for the registration.
	RegistrationDetails []*RestrictedFINXMax350Text `xml:"RegnDtls,omitempty"`

	// Provides additional information regarding the party, for example, the contact unit or person responsible for the transaction identified in the message.
	PartyContactNarrative []*RestrictedFINXMax350Text `xml:"PtyCtctNrrtv,omitempty"`

	// Provides information required for the certification/breakdown.
	CertificationBreakdown []*RestrictedFINXMax350Text `xml:"CertfctnBrkdwn,omitempty"`
}

func (c *CorporateActionNarrative34) AddRegistrationDetails(value string) {
	c.RegistrationDetails = append(c.RegistrationDetails, (*RestrictedFINXMax350Text)(&value))
}

func (c *CorporateActionNarrative34) AddPartyContactNarrative(value string) {
	c.PartyContactNarrative = append(c.PartyContactNarrative, (*RestrictedFINXMax350Text)(&value))
}

func (c *CorporateActionNarrative34) AddCertificationBreakdown(value string) {
	c.CertificationBreakdown = append(c.CertificationBreakdown, (*RestrictedFINXMax350Text)(&value))
}
