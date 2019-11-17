package model

// Chain of parties involved in the settlement of a transaction, including receipts and deliveries, book transfers, treasury deals, or other activities, resulting in the movement of a security or amount of money from one account to another.
type SettlementParties26 struct {

	// Party that, in a settlement chain interacts with the party 1.
	Party2 *PartyIdentificationAndAccount42 `xml:"Pty2,omitempty"`

	// Party that, in a settlement chain interacts with the party 2.
	Party3 *PartyIdentificationAndAccount42 `xml:"Pty3,omitempty"`

	// Party that, in a settlement chain interacts with the party 3.
	Party4 *PartyIdentificationAndAccount42 `xml:"Pty4,omitempty"`

	// Party that, in a settlement chain interacts with the party 4.
	Party5 *PartyIdentificationAndAccount42 `xml:"Pty5,omitempty"`
}

func (s *SettlementParties26) AddParty2() *PartyIdentificationAndAccount42 {
	s.Party2 = new(PartyIdentificationAndAccount42)
	return s.Party2
}

func (s *SettlementParties26) AddParty3() *PartyIdentificationAndAccount42 {
	s.Party3 = new(PartyIdentificationAndAccount42)
	return s.Party3
}

func (s *SettlementParties26) AddParty4() *PartyIdentificationAndAccount42 {
	s.Party4 = new(PartyIdentificationAndAccount42)
	return s.Party4
}

func (s *SettlementParties26) AddParty5() *PartyIdentificationAndAccount42 {
	s.Party5 = new(PartyIdentificationAndAccount42)
	return s.Party5
}
