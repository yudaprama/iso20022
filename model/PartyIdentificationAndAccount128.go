package model

// Provides information about identification and account of the party .
type PartyIdentificationAndAccount128 struct {

	// Identification of a party.
	Identification *PartyIdentification104Choice `xml:"Id"`

	// Account where financial instruments are maintained.
	SafekeepingAccount *RestrictedFINXMax35Text `xml:"SfkpgAcct,omitempty"`

	// Reference meaningful to the party identified.
	ProcessingIdentification *RestrictedFINXMax16Text `xml:"PrcgId,omitempty"`

	// Alternate identification for a party.
	AlternateIdentification []*AlternatePartyIdentification9 `xml:"AltrnId,omitempty"`
}

func (p *PartyIdentificationAndAccount128) AddIdentification() *PartyIdentification104Choice {
	p.Identification = new(PartyIdentification104Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount128) SetSafekeepingAccount(value string) {
	p.SafekeepingAccount = (*RestrictedFINXMax35Text)(&value)
}

func (p *PartyIdentificationAndAccount128) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (p *PartyIdentificationAndAccount128) AddAlternateIdentification() *AlternatePartyIdentification9 {
	newValue := new(AlternatePartyIdentification9)
	p.AlternateIdentification = append(p.AlternateIdentification, newValue)
	return newValue
}
