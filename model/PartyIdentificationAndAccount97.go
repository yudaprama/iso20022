package model

// Party and account information.
type PartyIdentificationAndAccount97 struct {

	// Identification of the party that legally owns the account.
	PartyIdentification *PartyIdentification62 `xml:"PtyId"`

	// Identification of the account.
	AccountIdentification *AccountIdentification26 `xml:"AcctId,omitempty"`
}

func (p *PartyIdentificationAndAccount97) AddPartyIdentification() *PartyIdentification62 {
	p.PartyIdentification = new(PartyIdentification62)
	return p.PartyIdentification
}

func (p *PartyIdentificationAndAccount97) AddAccountIdentification() *AccountIdentification26 {
	p.AccountIdentification = new(AccountIdentification26)
	return p.AccountIdentification
}
