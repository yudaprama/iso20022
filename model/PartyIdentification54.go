package model

// Choice of identification of a party. The party can be identified by providing a BIC or a proprietary code.
// Optionally, the client account number can also be provided.
type PartyIdentification54 struct {

	// Code allocated to a financial or non-financial institution by the ISO 9362 Registration Authority, as described in ISO 9362 "Banking - Banking telecommunication messages - Business identifier code (BIC)".
	BIC *AnyBICIdentifier `xml:"BIC"`

	// Unique and unambiguous identifier, as assigned to a financial institution using a proprietary identification scheme.
	ProprietaryIdentification *GenericIdentification29 `xml:"PrtryId"`

	// Identification of a party with its name and address in free text.
	NameAndAddress *NameAndAddress13 `xml:"NmAndAdr,omitempty"`
}

func (p *PartyIdentification54) SetBIC(value string) {
	p.BIC = (*AnyBICIdentifier)(&value)
}

func (p *PartyIdentification54) AddProprietaryIdentification() *GenericIdentification29 {
	p.ProprietaryIdentification = new(GenericIdentification29)
	return p.ProprietaryIdentification
}

func (p *PartyIdentification54) AddNameAndAddress() *NameAndAddress13 {
	p.NameAndAddress = new(NameAndAddress13)
	return p.NameAndAddress
}
