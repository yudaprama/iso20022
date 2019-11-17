package model

// Party and account details.
type PartyIdentificationAndAccount95 struct {

	// Identification of the party.
	PartyIdentification *PartyIdentification71Choice `xml:"PtyId"`

	// Account to or from which a securities entry is made.
	AccountIdentification *SecuritiesAccount22 `xml:"AcctId,omitempty"`

	// Unambiguous identification of the transaction for the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`
}

func (p *PartyIdentificationAndAccount95) AddPartyIdentification() *PartyIdentification71Choice {
	p.PartyIdentification = new(PartyIdentification71Choice)
	return p.PartyIdentification
}

func (p *PartyIdentificationAndAccount95) AddAccountIdentification() *SecuritiesAccount22 {
	p.AccountIdentification = new(SecuritiesAccount22)
	return p.AccountIdentification
}

func (p *PartyIdentificationAndAccount95) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}
