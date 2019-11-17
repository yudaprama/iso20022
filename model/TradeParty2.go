package model

// Specifies an entity involved in a trade activity.
type TradeParty2 struct {

	// Unique identification, as assigned by an organisation, to unambiguously identify a party.
	PartyIdentification *PartyIdentification77 `xml:"PtyId"`

	// Legally constituted organization specified for this trade party.
	LegalOrganisation *LegalOrganisation2 `xml:"LglOrg,omitempty"`

	// Entity involved in an activity.
	TaxParty []*TaxParty3 `xml:"TaxPty,omitempty"`
}

func (t *TradeParty2) AddPartyIdentification() *PartyIdentification77 {
	t.PartyIdentification = new(PartyIdentification77)
	return t.PartyIdentification
}

func (t *TradeParty2) AddLegalOrganisation() *LegalOrganisation2 {
	t.LegalOrganisation = new(LegalOrganisation2)
	return t.LegalOrganisation
}

func (t *TradeParty2) AddTaxParty() *TaxParty3 {
	newValue := new(TaxParty3)
	t.TaxParty = append(t.TaxParty, newValue)
	return newValue
}
