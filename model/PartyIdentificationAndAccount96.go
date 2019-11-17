package model

// Party and account information.
type PartyIdentificationAndAccount96 struct {

	// Identification of the party that legally owns the account.
	PartyIdentification *PartyIdentification64 `xml:"PtyId"`

	// Identification of the account.
	AccountIdentification *AccountIdentification26 `xml:"AcctId"`
}

func (p *PartyIdentificationAndAccount96) AddPartyIdentification() *PartyIdentification64 {
	p.PartyIdentification = new(PartyIdentification64)
	return p.PartyIdentification
}

func (p *PartyIdentificationAndAccount96) AddAccountIdentification() *AccountIdentification26 {
	p.AccountIdentification = new(AccountIdentification26)
	return p.AccountIdentification
}
