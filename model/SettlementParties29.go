package model

// Identification of a settlement party by a choice between a BIC or a name and address or a party identification.
type SettlementParties29 struct {

	// Financial institution from which cash will be transferred.
	DeliveryAgent *PartyIdentification73Choice `xml:"DlvryAgt,omitempty"`

	// Party, within the settlement chain, between the delivery and receiving agents.
	Intermediary *PartyIdentification73Choice `xml:"Intrmy,omitempty"`

	// Financial institution where the payee will receive the funds.
	ReceivingAgent *PartyIdentification73Choice `xml:"RcvgAgt"`

	// Ultimate institution that will receive the funds when different from the trading or counterparty side.
	BeneficiaryInstitution *PartyIdentification73Choice `xml:"BnfcryInstn,omitempty"`
}

func (s *SettlementParties29) AddDeliveryAgent() *PartyIdentification73Choice {
	s.DeliveryAgent = new(PartyIdentification73Choice)
	return s.DeliveryAgent
}

func (s *SettlementParties29) AddIntermediary() *PartyIdentification73Choice {
	s.Intermediary = new(PartyIdentification73Choice)
	return s.Intermediary
}

func (s *SettlementParties29) AddReceivingAgent() *PartyIdentification73Choice {
	s.ReceivingAgent = new(PartyIdentification73Choice)
	return s.ReceivingAgent
}

func (s *SettlementParties29) AddBeneficiaryInstitution() *PartyIdentification73Choice {
	s.BeneficiaryInstitution = new(PartyIdentification73Choice)
	return s.BeneficiaryInstitution
}
