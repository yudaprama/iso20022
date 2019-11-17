package model

// Identification of a person, or a non-financial institution.
type PartyIdentification25 struct {

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max70Text `xml:"Nm"`

	// Unique and unambiguous identifier, as assigned to a party using a proprietary identification scheme.
	ProprietaryIdentification *GenericIdentification4 `xml:"PrtryId,omitempty"`

	// Identification of a non-financial institution.
	BEI *BEIIdentifier `xml:"BEI,omitempty"`
}

func (p *PartyIdentification25) SetName(value string) {
	p.Name = (*Max70Text)(&value)
}

func (p *PartyIdentification25) AddProprietaryIdentification() *GenericIdentification4 {
	p.ProprietaryIdentification = new(GenericIdentification4)
	return p.ProprietaryIdentification
}

func (p *PartyIdentification25) SetBEI(value string) {
	p.BEI = (*BEIIdentifier)(&value)
}
