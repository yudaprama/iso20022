package model

// Unique and unambiguous way to identify an organisation.
type OrganisationIdentification2 struct {

	// Code allocated to a financial institution by the ISO 9362 Registration Authority as described in ISO 9362 "Banking - Banking telecommunication messages - Business identifier code (BIC)".
	BIC *BICIdentifier `xml:"BIC,omitempty"`

	// International Business Entity Identifier to uniquely identify business entities playing a role in the lifecycle of and events related to a financial instrument. (tentative - to be confirmed)
	IBEI *IBEIIdentifier `xml:"IBEI,omitempty"`

	// Code allocated to a non-financial institution by the ISO 9362 Registration Authority as described in ISO 9362 "Banking - Banking telecommunication messages - Business identifier code (BIC)".
	BEI *BEIIdentifier `xml:"BEI,omitempty"`

	// Global Location Number. A non-significant reference number used to identify legal entities, functional entities, or physical entities according to the European Association for Numbering (EAN) numbering scheme rules.The number is used to retrieve detailed information that is linked to it.
	EANGLN *EANGLNIdentifier `xml:"EANGLN,omitempty"`

	// (United States) Clearing House Interbank Payments System (CHIPS) Universal Identification (UID) - identifies entities that own accounts at CHIPS participating financial institutions, through which CHIPS payments are effected. The CHIPS UID is assigned by the New York Clearing House.
	CHIPSUniversalIdentification *CHIPSUniversalIdentifier `xml:"USCHU,omitempty"`

	// Data Universal Numbering System. A unique identification number provided by Dun & Bradstreet to identify an organization.
	DUNS *DunsIdentifier `xml:"DUNS,omitempty"`

	// Unique and unambiguous assignment made by a specific bank to identify a relationship as defined between the bank and its client.
	BankPartyIdentification *Max35Text `xml:"BkPtyId,omitempty"`

	// Number assigned by a tax authority to an entity.
	TaxIdentificationNumber *Max35Text `xml:"TaxIdNb,omitempty"`

	// Unique and unambiguous identifier for an organisation that is allocated by an institution.
	ProprietaryIdentification *GenericIdentification3 `xml:"PrtryId,omitempty"`
}

func (o *OrganisationIdentification2) SetBIC(value string) {
	o.BIC = (*BICIdentifier)(&value)
}

func (o *OrganisationIdentification2) SetIBEI(value string) {
	o.IBEI = (*IBEIIdentifier)(&value)
}

func (o *OrganisationIdentification2) SetBEI(value string) {
	o.BEI = (*BEIIdentifier)(&value)
}

func (o *OrganisationIdentification2) SetEANGLN(value string) {
	o.EANGLN = (*EANGLNIdentifier)(&value)
}

func (o *OrganisationIdentification2) SetCHIPSUniversalIdentification(value string) {
	o.CHIPSUniversalIdentification = (*CHIPSUniversalIdentifier)(&value)
}

func (o *OrganisationIdentification2) SetDUNS(value string) {
	o.DUNS = (*DunsIdentifier)(&value)
}

func (o *OrganisationIdentification2) SetBankPartyIdentification(value string) {
	o.BankPartyIdentification = (*Max35Text)(&value)
}

func (o *OrganisationIdentification2) SetTaxIdentificationNumber(value string) {
	o.TaxIdentificationNumber = (*Max35Text)(&value)
}

func (o *OrganisationIdentification2) AddProprietaryIdentification() *GenericIdentification3 {
	o.ProprietaryIdentification = new(GenericIdentification3)
	return o.ProprietaryIdentification
}
