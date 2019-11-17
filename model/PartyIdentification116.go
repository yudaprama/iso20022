package model

// Unique identification, as assigned by an organisation, to unambiguously identify a party.
type PartyIdentification116 struct {

	// Unique identification, as assigned by an organisation, to unambiguously identify a party.
	PartyIdentification *OrganisationIdentification28 `xml:"PtyId"`

	// Identifier and name of an organisation that is allocated by an institution.
	LegalOrganisation *LegalOrganisation1 `xml:"LglOrg,omitempty"`

	// TaxParty
	TaxParty *TaxParty1 `xml:"TaxPty,omitempty"`
}

func (p *PartyIdentification116) AddPartyIdentification() *OrganisationIdentification28 {
	p.PartyIdentification = new(OrganisationIdentification28)
	return p.PartyIdentification
}

func (p *PartyIdentification116) AddLegalOrganisation() *LegalOrganisation1 {
	p.LegalOrganisation = new(LegalOrganisation1)
	return p.LegalOrganisation
}

func (p *PartyIdentification116) AddTaxParty() *TaxParty1 {
	p.TaxParty = new(TaxParty1)
	return p.TaxParty
}
