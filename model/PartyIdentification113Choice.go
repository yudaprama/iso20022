package model

// Choice of identification of a party.
type PartyIdentification113Choice struct {

	// Code allocated to a financial institution by the ISO 9362 Registration Authority, as described in ISO 9362 "Banking - Banking telecommunication messages - Business identifier code (BIC)".
	BICFI *BICFIIdentifier `xml:"BICFI"`

	// Name by which a party is known and which is usually used to identify that party.
	NameAndAddress *NameAndAddress12 `xml:"NmAndAdr"`

	// Unique and unambiguous identifier, as assigned to a financial institution using a proprietary identification scheme.
	ProprietaryIdentification *GenericIdentification84 `xml:"PrtryId"`
}

func (p *PartyIdentification113Choice) SetBICFI(value string) {
	p.BICFI = (*BICFIIdentifier)(&value)
}

func (p *PartyIdentification113Choice) AddNameAndAddress() *NameAndAddress12 {
	p.NameAndAddress = new(NameAndAddress12)
	return p.NameAndAddress
}

func (p *PartyIdentification113Choice) AddProprietaryIdentification() *GenericIdentification84 {
	p.ProprietaryIdentification = new(GenericIdentification84)
	return p.ProprietaryIdentification
}
