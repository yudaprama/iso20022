package model

// Provides information about identification and account of the party .
type PartyIdentificationAndAccount122 struct {

	// Identification of a party.
	Identification *PartyIdentification71Choice `xml:"Id"`

	// Account where financial instruments are maintained.
	SafekeepingAccount *Max35Text `xml:"SfkpgAcct,omitempty"`

	// Reference meaningful to the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`

	// Alternate identification for a party.
	AlternateIdentification []*AlternatePartyIdentification7 `xml:"AltrnId,omitempty"`
}

func (p *PartyIdentificationAndAccount122) AddIdentification() *PartyIdentification71Choice {
	p.Identification = new(PartyIdentification71Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount122) SetSafekeepingAccount(value string) {
	p.SafekeepingAccount = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount122) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount122) AddAlternateIdentification() *AlternatePartyIdentification7 {
	newValue := new(AlternatePartyIdentification7)
	p.AlternateIdentification = append(p.AlternateIdentification, newValue)
	return newValue
}
