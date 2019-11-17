package model

// Choice of identification of a party.
type PartyIdentification10Choice struct {

	// Code allocated to a financial or non-financial institution by the ISO 9362 Registration Authority, as described in ISO 9362 "Banking - Banking telecommunication messages - Business identifier code (BIC)".
	BICOrBEI *AnyBICIdentifier `xml:"BICOrBEI"`

	// Unique and unambiguous identifier, as assigned to a financial institution using a proprietary identification scheme.
	ProprietaryIdentification *GenericIdentification19 `xml:"PrtryId"`

	// Name and address of a party.
	NameAndAddress *NameAndAddress5 `xml:"NmAndAdr"`
}

func (p *PartyIdentification10Choice) SetBICOrBEI(value string) {
	p.BICOrBEI = (*AnyBICIdentifier)(&value)
}

func (p *PartyIdentification10Choice) AddProprietaryIdentification() *GenericIdentification19 {
	p.ProprietaryIdentification = new(GenericIdentification19)
	return p.ProprietaryIdentification
}

func (p *PartyIdentification10Choice) AddNameAndAddress() *NameAndAddress5 {
	p.NameAndAddress = new(NameAndAddress5)
	return p.NameAndAddress
}
