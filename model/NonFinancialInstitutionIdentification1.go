package model

// Unique and unambiguous identification of a non-financial institution.
type NonFinancialInstitutionIdentification1 struct {

	// Code allocated to a non-financial institution by the ISO 9362 Registration Authority as described in ISO 9362 "Banking - Banking telecommunication messages - Business identifier code (BIC)".
	BEI *BEIIdentifier `xml:"BEI,omitempty"`

	// Global Location Number.  A non-significant reference number used to identify legal entities, functional entities or physical entities according to the European Association for Numbering (EAN) numbering scheme rules. The number is used to retrieve detailed information linked to it.
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

func (n *NonFinancialInstitutionIdentification1) SetBEI(value string) {
	n.BEI = (*BEIIdentifier)(&value)
}

func (n *NonFinancialInstitutionIdentification1) SetEANGLN(value string) {
	n.EANGLN = (*EANGLNIdentifier)(&value)
}

func (n *NonFinancialInstitutionIdentification1) SetCHIPSUniversalIdentification(value string) {
	n.CHIPSUniversalIdentification = (*CHIPSUniversalIdentifier)(&value)
}

func (n *NonFinancialInstitutionIdentification1) SetDUNS(value string) {
	n.DUNS = (*DunsIdentifier)(&value)
}

func (n *NonFinancialInstitutionIdentification1) SetBankPartyIdentification(value string) {
	n.BankPartyIdentification = (*Max35Text)(&value)
}

func (n *NonFinancialInstitutionIdentification1) SetTaxIdentificationNumber(value string) {
	n.TaxIdentificationNumber = (*Max35Text)(&value)
}

func (n *NonFinancialInstitutionIdentification1) AddProprietaryIdentification() *GenericIdentification3 {
	n.ProprietaryIdentification = new(GenericIdentification3)
	return n.ProprietaryIdentification
}
