package model

// Choice of formats for the identification of a party.
type PartyIdentification96Choice struct {

	// Identification of the party expressed as a BIC.
	AnyBIC *AnyBICIdentifier `xml:"AnyBIC"`

	// Unique and unambiguous identifier, as assigned to the party using a proprietary identification scheme.
	ProprietaryIdentification *GenericIdentification1 `xml:"PrtryId"`

	// Name and address of the party.
	NameAndAddress *NameAndAddress15 `xml:"NmAndAdr"`

	// Number assigned by a tax authority to an entity.
	TaxIdentificationNumber *Max35Text `xml:"TaxIdNb"`

	// Number assigned by a national registration authority to an entity.
	NationalRegistrationNumber *Max35Text `xml:"NtlRegnNb"`
}

func (p *PartyIdentification96Choice) SetAnyBIC(value string) {
	p.AnyBIC = (*AnyBICIdentifier)(&value)
}

func (p *PartyIdentification96Choice) AddProprietaryIdentification() *GenericIdentification1 {
	p.ProprietaryIdentification = new(GenericIdentification1)
	return p.ProprietaryIdentification
}

func (p *PartyIdentification96Choice) AddNameAndAddress() *NameAndAddress15 {
	p.NameAndAddress = new(NameAndAddress15)
	return p.NameAndAddress
}

func (p *PartyIdentification96Choice) SetTaxIdentificationNumber(value string) {
	p.TaxIdentificationNumber = (*Max35Text)(&value)
}

func (p *PartyIdentification96Choice) SetNationalRegistrationNumber(value string) {
	p.NationalRegistrationNumber = (*Max35Text)(&value)
}
