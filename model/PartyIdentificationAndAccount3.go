package model

// Party and account information.
type PartyIdentificationAndAccount3 struct {

	// Party that legally owns the account.
	PartyIdentification *PartyIdentification2Choice `xml:"PtyId"`

	// Identification of the account owned by the party.
	AccountIdentification *AccountIdentification1 `xml:"AcctId,omitempty"`
}

func (p *PartyIdentificationAndAccount3) AddPartyIdentification() *PartyIdentification2Choice {
	p.PartyIdentification = new(PartyIdentification2Choice)
	return p.PartyIdentification
}

func (p *PartyIdentificationAndAccount3) AddAccountIdentification() *AccountIdentification1 {
	p.AccountIdentification = new(AccountIdentification1)
	return p.AccountIdentification
}
