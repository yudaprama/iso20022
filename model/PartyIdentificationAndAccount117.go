package model

// Party and account details.
type PartyIdentificationAndAccount117 struct {

	// Identification of the party.
	Identification *PartyIdentification71Choice `xml:"Id"`

	// Legal entity identification as an alternate identification for a party.
	LEI *LEIIdentifier `xml:"LEI,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount19 `xml:"SfkpgAcct,omitempty"`

	// Unambiguous identification of the transaction for the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`
}

func (p *PartyIdentificationAndAccount117) AddIdentification() *PartyIdentification71Choice {
	p.Identification = new(PartyIdentification71Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount117) SetLEI(value string) {
	p.LEI = (*LEIIdentifier)(&value)
}

func (p *PartyIdentificationAndAccount117) AddSafekeepingAccount() *SecuritiesAccount19 {
	p.SafekeepingAccount = new(SecuritiesAccount19)
	return p.SafekeepingAccount
}

func (p *PartyIdentificationAndAccount117) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}
