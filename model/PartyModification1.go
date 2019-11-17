package model

// Specifies the type of change to party.
type PartyModification1 struct {

	// Specifies the type of change.
	ModificationCode *Modification1Code `xml:"ModCd,omitempty"`

	// Identification of the party.
	PartyIdentification *PartyIdentification40 `xml:"PtyId"`
}

func (p *PartyModification1) SetModificationCode(value string) {
	p.ModificationCode = (*Modification1Code)(&value)
}

func (p *PartyModification1) AddPartyIdentification() *PartyIdentification40 {
	p.PartyIdentification = new(PartyIdentification40)
	return p.PartyIdentification
}
