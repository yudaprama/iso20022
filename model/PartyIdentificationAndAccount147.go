package model

// Party and account information.
type PartyIdentificationAndAccount147 struct {

	// Party that legally owns the account.
	PartyIdentification *PartyIdentification113 `xml:"PtyId"`

	// Identification of the account owned by the party.
	AccountIdentification *Max35Text `xml:"AcctId,omitempty"`
}

func (p *PartyIdentificationAndAccount147) AddPartyIdentification() *PartyIdentification113 {
	p.PartyIdentification = new(PartyIdentification113)
	return p.PartyIdentification
}

func (p *PartyIdentificationAndAccount147) SetAccountIdentification(value string) {
	p.AccountIdentification = (*Max35Text)(&value)
}
