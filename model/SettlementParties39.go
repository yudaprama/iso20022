package model

// Chain of parties involved in the settlement of a transaction, including receipts and deliveries, book transfers, treasury deals, or other activities, resulting in the movement of a security or amount of money from one account to another.
type SettlementParties39 struct {

	// Party that, in a settlement chain interacts with the party 1.
	Party2 *PartyIdentificationAndAccount106 `xml:"Pty2,omitempty"`

	// Party that, in a settlement chain interacts with the party 2.
	Party3 *PartyIdentificationAndAccount106 `xml:"Pty3,omitempty"`

	// Party that, in a settlement chain interacts with the party 3.
	Party4 *PartyIdentificationAndAccount106 `xml:"Pty4,omitempty"`

	// Party that, in a settlement chain interacts with the party 4.
	Party5 *PartyIdentificationAndAccount106 `xml:"Pty5,omitempty"`
}

func (s *SettlementParties39) AddParty2() *PartyIdentificationAndAccount106 {
	s.Party2 = new(PartyIdentificationAndAccount106)
	return s.Party2
}

func (s *SettlementParties39) AddParty3() *PartyIdentificationAndAccount106 {
	s.Party3 = new(PartyIdentificationAndAccount106)
	return s.Party3
}

func (s *SettlementParties39) AddParty4() *PartyIdentificationAndAccount106 {
	s.Party4 = new(PartyIdentificationAndAccount106)
	return s.Party4
}

func (s *SettlementParties39) AddParty5() *PartyIdentificationAndAccount106 {
	s.Party5 = new(PartyIdentificationAndAccount106)
	return s.Party5
}
