package model

// Party and account details.
type PartyIdentificationAndAccount16 struct {

	// Identification of the party.
	Identification *PartyIdentification10Choice `xml:"Id"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount13 `xml:"SfkpgAcct,omitempty"`

	// Unambiguous identification of the transaction for the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`
}

func (p *PartyIdentificationAndAccount16) AddIdentification() *PartyIdentification10Choice {
	p.Identification = new(PartyIdentification10Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount16) AddSafekeepingAccount() *SecuritiesAccount13 {
	p.SafekeepingAccount = new(SecuritiesAccount13)
	return p.SafekeepingAccount
}

func (p *PartyIdentificationAndAccount16) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}
