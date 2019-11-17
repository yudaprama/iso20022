package model

// Provides additional information regarding the party, eg, the contact unit or person responsible for the transaction identified in the message.
type PartyTextInformation3 struct {

	// Provides declaration details narrative relative to the party.
	DeclarationDetails *RestrictedFINXMax350Text `xml:"DclrtnDtls,omitempty"`

	// Provides additional information regarding the party, for example, the contact unit or person responsible for the transaction identified in the message.
	PartyContactDetails *RestrictedFINXMax140Text `xml:"PtyCtctDtls,omitempty"`

	// Provides information required for the registration.
	RegistrationDetails *RestrictedFINXMax350Text `xml:"RegnDtls,omitempty"`
}

func (p *PartyTextInformation3) SetDeclarationDetails(value string) {
	p.DeclarationDetails = (*RestrictedFINXMax350Text)(&value)
}

func (p *PartyTextInformation3) SetPartyContactDetails(value string) {
	p.PartyContactDetails = (*RestrictedFINXMax140Text)(&value)
}

func (p *PartyTextInformation3) SetRegistrationDetails(value string) {
	p.RegistrationDetails = (*RestrictedFINXMax350Text)(&value)
}
