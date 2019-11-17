package model

// Set of elements used to identify an organization or a person.
type PartyIdentification72 struct {

	// Unique identification, as assigned by an organisation, to unambiguously identify a party.
	PartyIdentification *PartyIdentification43 `xml:"PtyId"`

	// Identifier and name of an organisation that is allocated by an institution.
	LegalOrganisation *LegalOrganisation1 `xml:"LglOrg,omitempty"`

	// Tax registration details.
	TaxParty *TaxParty1 `xml:"TaxPty,omitempty"`
}

func (p *PartyIdentification72) AddPartyIdentification() *PartyIdentification43 {
	p.PartyIdentification = new(PartyIdentification43)
	return p.PartyIdentification
}

func (p *PartyIdentification72) AddLegalOrganisation() *LegalOrganisation1 {
	p.LegalOrganisation = new(LegalOrganisation1)
	return p.LegalOrganisation
}

func (p *PartyIdentification72) AddTaxParty() *TaxParty1 {
	p.TaxParty = new(TaxParty1)
	return p.TaxParty
}
