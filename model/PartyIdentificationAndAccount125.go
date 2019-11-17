package model

// Party involved in the settlement chain.
type PartyIdentificationAndAccount125 struct {

	// Party that legally owns the account.
	PartyIdentification *PartyIdentification70Choice `xml:"PtyId,omitempty"`

	// Identification of the account owned by the party.
	AccountIdentification *Max35Text `xml:"AcctId,omitempty"`

	// Place where settlement of the securities takes place.
	PlaceOfSettlement *PartyIdentification70Choice `xml:"PlcOfSttlm"`
}

func (p *PartyIdentificationAndAccount125) AddPartyIdentification() *PartyIdentification70Choice {
	p.PartyIdentification = new(PartyIdentification70Choice)
	return p.PartyIdentification
}

func (p *PartyIdentificationAndAccount125) SetAccountIdentification(value string) {
	p.AccountIdentification = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount125) AddPlaceOfSettlement() *PartyIdentification70Choice {
	p.PlaceOfSettlement = new(PartyIdentification70Choice)
	return p.PlaceOfSettlement
}
