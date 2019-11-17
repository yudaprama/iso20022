package model

// Party involved in the settlement chain.
type PartyIdentificationAndAccount93 struct {

	// Party that legally owns the account.
	PartyIdentification *PartyIdentification2Choice `xml:"PtyId,omitempty"`

	// Identification of the account owned by the party.
	AccountIdentification *AccountIdentification1 `xml:"AcctId,omitempty"`

	// Place where settlement of the securities takes place.
	PlaceOfSettlement *PartyIdentification2Choice `xml:"PlcOfSttlm"`
}

func (p *PartyIdentificationAndAccount93) AddPartyIdentification() *PartyIdentification2Choice {
	p.PartyIdentification = new(PartyIdentification2Choice)
	return p.PartyIdentification
}

func (p *PartyIdentificationAndAccount93) AddAccountIdentification() *AccountIdentification1 {
	p.AccountIdentification = new(AccountIdentification1)
	return p.AccountIdentification
}

func (p *PartyIdentificationAndAccount93) AddPlaceOfSettlement() *PartyIdentification2Choice {
	p.PlaceOfSettlement = new(PartyIdentification2Choice)
	return p.PlaceOfSettlement
}
