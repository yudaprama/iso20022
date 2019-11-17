package model

// Provides additional information regarding the party, eg, the contact unit or person responsible for the transaction identified in the message.
type PartyTextInformation1 struct {

	// Provides declaration details narrative relative to the party.
	DeclarationDetails *Max350Text `xml:"DclrtnDtls,omitempty"`

	// Provides additional information regarding the party, for example, the contact unit or person responsible for the transaction identified in the message.
	PartyContactDetails *Max140Text `xml:"PtyCtctDtls,omitempty"`

	// Provides information required for the registration.
	RegistrationDetails *Max350Text `xml:"RegnDtls,omitempty"`
}

func (p *PartyTextInformation1) SetDeclarationDetails(value string) {
	p.DeclarationDetails = (*Max350Text)(&value)
}

func (p *PartyTextInformation1) SetPartyContactDetails(value string) {
	p.PartyContactDetails = (*Max140Text)(&value)
}

func (p *PartyTextInformation1) SetRegistrationDetails(value string) {
	p.RegistrationDetails = (*Max350Text)(&value)
}
