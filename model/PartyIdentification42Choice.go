package model

// Choice of identification of a party.
type PartyIdentification42Choice struct {

	// Code allocated to a financial institution by the ISO 9362 Registration Authority, as described in ISO 9362 "Banking - Banking telecommunication messages - Business identifier code (BIC)".
	BICFI *BICFIIdentifier `xml:"BICFI"`

	// Name by which a party is known and which is usually used to identify that party.
	NameAndAddress *NameAndAddress5 `xml:"NmAndAdr"`

	// Unique and unambiguous identifier, as assigned to a financial institution using a proprietary identification scheme.
	ProprietaryIdentification *GenericIdentification19 `xml:"PrtryId"`
}

func (p *PartyIdentification42Choice) SetBICFI(value string) {
	p.BICFI = (*BICFIIdentifier)(&value)
}

func (p *PartyIdentification42Choice) AddNameAndAddress() *NameAndAddress5 {
	p.NameAndAddress = new(NameAndAddress5)
	return p.NameAndAddress
}

func (p *PartyIdentification42Choice) AddProprietaryIdentification() *GenericIdentification19 {
	p.ProprietaryIdentification = new(GenericIdentification19)
	return p.ProprietaryIdentification
}
