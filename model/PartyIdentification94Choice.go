package model

// Choice of identification of a party.
type PartyIdentification94Choice struct {

	// Code allocated to a financial institution by the ISO 9362 Registration Authority, as described in ISO 9362 "Banking - Banking telecommunication messages - Business identifier code (BIC)".
	BICFI *BICFIIdentifier `xml:"BICFI"`

	// Name by which a party is known and which is usually used to identify that party.
	NameAndAddress *NameAndAddress5 `xml:"NmAndAdr"`

	// Unique and unambiguous identifier, as assigned to a financial institution using a proprietary identification scheme.
	ProprietaryIdentification *GenericIdentification36 `xml:"PrtryId"`
}

func (p *PartyIdentification94Choice) SetBICFI(value string) {
	p.BICFI = (*BICFIIdentifier)(&value)
}

func (p *PartyIdentification94Choice) AddNameAndAddress() *NameAndAddress5 {
	p.NameAndAddress = new(NameAndAddress5)
	return p.NameAndAddress
}

func (p *PartyIdentification94Choice) AddProprietaryIdentification() *GenericIdentification36 {
	p.ProprietaryIdentification = new(GenericIdentification36)
	return p.ProprietaryIdentification
}
