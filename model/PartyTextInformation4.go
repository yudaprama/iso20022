package model

// Provides additional information regarding the party, for example, the contact unit or person responsible for the transaction identified in the message.
type PartyTextInformation4 struct {

	// Provides declaration details narrative relative to the party.
	DeclarationDetails *RestrictedFINXMax350Text `xml:"DclrtnDtls,omitempty"`

	// Provides additional information regarding the party, for example, the contact unit or person responsible for the transaction identified in the message.
	PartyContactDetails *RestrictedFINXMax140Text `xml:"PtyCtctDtls,omitempty"`
}

func (p *PartyTextInformation4) SetDeclarationDetails(value string) {
	p.DeclarationDetails = (*RestrictedFINXMax350Text)(&value)
}

func (p *PartyTextInformation4) SetPartyContactDetails(value string) {
	p.PartyContactDetails = (*RestrictedFINXMax140Text)(&value)
}
