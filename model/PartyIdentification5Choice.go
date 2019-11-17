package model

// Choice of identification of a party.
type PartyIdentification5Choice struct {

	// Code allocated to a financial or non-financial institution by the ISO 9362 Registration Authority, as described in ISO 9362 "Banking - Banking telecommunication messages - Business identifier code (BIC)".
	BICOrBEI *AnyBICIdentifier `xml:"BICOrBEI"`

	// Unique and unambiguous identifier, as assigned to a financial institution using a proprietary identification scheme.
	ProprietaryIdentification *GenericIdentification1 `xml:"PrtryId"`

	// Name by which a party is known and which is usually used to identify that party.
	NameAndAddress *NameAndAddress5 `xml:"NmAndAdr"`

	// Number assigned by a tax authority to an entity.
	TaxIdentificationNumber *Max35Text `xml:"TaxIdNb"`

	// Number assigned by a national registration authority to an entity.
	NationalRegistrationNumber *Max35Text `xml:"NtlRegnNb"`
}

func (p *PartyIdentification5Choice) SetBICOrBEI(value string) {
	p.BICOrBEI = (*AnyBICIdentifier)(&value)
}

func (p *PartyIdentification5Choice) AddProprietaryIdentification() *GenericIdentification1 {
	p.ProprietaryIdentification = new(GenericIdentification1)
	return p.ProprietaryIdentification
}

func (p *PartyIdentification5Choice) AddNameAndAddress() *NameAndAddress5 {
	p.NameAndAddress = new(NameAndAddress5)
	return p.NameAndAddress
}

func (p *PartyIdentification5Choice) SetTaxIdentificationNumber(value string) {
	p.TaxIdentificationNumber = (*Max35Text)(&value)
}

func (p *PartyIdentification5Choice) SetNationalRegistrationNumber(value string) {
	p.NationalRegistrationNumber = (*Max35Text)(&value)
}
