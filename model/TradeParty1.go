package model

// Specifies an entity involved in a trade activity.
type TradeParty1 struct {

	// Unique identification, as assigned by an organisation, to unambiguously identify a party.
	PartyIdentification *PartyIdentification45 `xml:"PtyId"`

	// Legally constituted organization specified for this trade party.
	LegalOrganisation *LegalOrganisation1 `xml:"LglOrg,omitempty"`

	// Entity involved in an activity.
	TaxParty []*TaxParty3 `xml:"TaxPty,omitempty"`
}

func (t *TradeParty1) AddPartyIdentification() *PartyIdentification45 {
	t.PartyIdentification = new(PartyIdentification45)
	return t.PartyIdentification
}

func (t *TradeParty1) AddLegalOrganisation() *LegalOrganisation1 {
	t.LegalOrganisation = new(LegalOrganisation1)
	return t.LegalOrganisation
}

func (t *TradeParty1) AddTaxParty() *TaxParty3 {
	newValue := new(TaxParty3)
	t.TaxParty = append(t.TaxParty, newValue)
	return newValue
}
