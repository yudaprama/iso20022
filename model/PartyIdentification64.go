package model

// Identification of a party.
type PartyIdentification64 struct {

	// Identification of the party expressed as a BIC.
	AnyBIC *AnyBICIdentifier `xml:"AnyBIC,omitempty"`

	// Unique and unambiguous identifier, as assigned to the party using a proprietary identification scheme.
	ProprietaryIdentification *GenericIdentification1 `xml:"PrtryId,omitempty"`

	// Name and address of the party.
	NameAndAddress *NameAndAddress5 `xml:"NmAndAdr,omitempty"`
}

func (p *PartyIdentification64) SetAnyBIC(value string) {
	p.AnyBIC = (*AnyBICIdentifier)(&value)
}

func (p *PartyIdentification64) AddProprietaryIdentification() *GenericIdentification1 {
	p.ProprietaryIdentification = new(GenericIdentification1)
	return p.ProprietaryIdentification
}

func (p *PartyIdentification64) AddNameAndAddress() *NameAndAddress5 {
	p.NameAndAddress = new(NameAndAddress5)
	return p.NameAndAddress
}
