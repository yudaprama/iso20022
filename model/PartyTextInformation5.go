package model

// Provides additional information regarding the party, eg, the contact unit or person responsible for the transaction identified in the message.
type PartyTextInformation5 struct {

	// Provides declaration details narrative relative to the party.
	DeclarationDetails *Max350Text `xml:"DclrtnDtls,omitempty"`

	// Provides additional information regarding the party, eg, the contact unit or person responsible for the transaction identified in the message.
	PartyContactDetails *Max140Text `xml:"PtyCtctDtls,omitempty"`
}

func (p *PartyTextInformation5) SetDeclarationDetails(value string) {
	p.DeclarationDetails = (*Max350Text)(&value)
}

func (p *PartyTextInformation5) SetPartyContactDetails(value string) {
	p.PartyContactDetails = (*Max140Text)(&value)
}
