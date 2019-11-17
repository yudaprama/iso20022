package model

// Chain of parties involved in the settlement of a transaction, including receipts and deliveries, book transfers, treasury deals, or other activities, resulting in the movement of a security or amount of money from one account to another.
type SettlementParties35 struct {

	// Parties through which settlement is to take place.
	StandingSettlementParties *SettlementParties32 `xml:"StgSttlmPties"`

	// Identifier needed for settlement purposes. This identifier could be, for example, an identifier that identifies an institution or agent at a CDS or ICSD (Depository Trust Clearing Corporation (DTC) Institution ID or DTC Agent ID). It could also be a local tax identification number or an ‘investor identification’, as mandated by local market practice.
	LocalMarketIdentification []*GenericIdentification49 `xml:"LclMktId,omitempty"`

	// Registration information required for settlement. For some markets, for example, Spain (Iberclear) registration details are mandatory and should be part of the SSI. In some cases, the name of the institution is different than what's provided in the BIC Directory. If this is the case, the name should be provided.
	RegistrationDetails *PartyIdentification99Choice `xml:"RegnDtls,omitempty"`
}

func (s *SettlementParties35) AddStandingSettlementParties() *SettlementParties32 {
	s.StandingSettlementParties = new(SettlementParties32)
	return s.StandingSettlementParties
}

func (s *SettlementParties35) AddLocalMarketIdentification() *GenericIdentification49 {
	newValue := new(GenericIdentification49)
	s.LocalMarketIdentification = append(s.LocalMarketIdentification, newValue)
	return newValue
}

func (s *SettlementParties35) AddRegistrationDetails() *PartyIdentification99Choice {
	s.RegistrationDetails = new(PartyIdentification99Choice)
	return s.RegistrationDetails
}
