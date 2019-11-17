package model

// Entity involved in an activity.
type PartyIdentification28 struct {

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max70Text `xml:"Nm"`

	// Unique and unambiguous identifier assigned to a party using a proprietary identification scheme.
	ProprietaryIdentification *GenericIdentification4 `xml:"PrtryId,omitempty"`
}

func (p *PartyIdentification28) SetName(value string) {
	p.Name = (*Max70Text)(&value)
}

func (p *PartyIdentification28) AddProprietaryIdentification() *GenericIdentification4 {
	p.ProprietaryIdentification = new(GenericIdentification4)
	return p.ProprietaryIdentification
}
