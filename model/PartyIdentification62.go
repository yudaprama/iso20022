package model

// Identification of a party.
type PartyIdentification62 struct {

	// Identification of the financial institution expressed as a BIC.
	BICFI *BICFIIdentifier `xml:"BICFI,omitempty"`

	// Unique and unambiguous identifier, as assigned to a financial institution using a proprietary identification scheme.
	ProprietaryIdentification *GenericIdentification1 `xml:"PrtryId,omitempty"`

	// Name and address of the party.
	NameAndAddress *NameAndAddress5 `xml:"NmAndAdr,omitempty"`
}

func (p *PartyIdentification62) SetBICFI(value string) {
	p.BICFI = (*BICFIIdentifier)(&value)
}

func (p *PartyIdentification62) AddProprietaryIdentification() *GenericIdentification1 {
	p.ProprietaryIdentification = new(GenericIdentification1)
	return p.ProprietaryIdentification
}

func (p *PartyIdentification62) AddNameAndAddress() *NameAndAddress5 {
	p.NameAndAddress = new(NameAndAddress5)
	return p.NameAndAddress
}
