package model

// Party and account details.
type PartyIdentificationAndAccount43 struct {

	// Identification of the party.
	Identification *PartyIdentification43Choice `xml:"Id"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount13 `xml:"SfkpgAcct,omitempty"`

	// Unambiguous identification of the transaction for the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`
}

func (p *PartyIdentificationAndAccount43) AddIdentification() *PartyIdentification43Choice {
	p.Identification = new(PartyIdentification43Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount43) AddSafekeepingAccount() *SecuritiesAccount13 {
	p.SafekeepingAccount = new(SecuritiesAccount13)
	return p.SafekeepingAccount
}

func (p *PartyIdentificationAndAccount43) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}
