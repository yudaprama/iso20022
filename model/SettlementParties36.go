package model

// Chain of parties involved in the settlement of a transaction, including receipts and deliveries, book transfers, treasury deals, or other activities, resulting in the movement of a security or amount of money from one account to another.
type SettlementParties36 struct {

	// First party in the settlement chain. In a plain vanilla settlement, it is the Central Securities Depository where the counterparty requests to receive the financial instrument or from where the counterparty delivers the financial instruments.
	Depository *PartyIdentification75 `xml:"Dpstry,omitempty"`

	// Party that, in a settlement chain interacts with the depository.
	Party1 *PartyIdentificationAndAccount106 `xml:"Pty1,omitempty"`

	// Party that, in a settlement chain interacts with the party 1.
	Party2 *PartyIdentificationAndAccount106 `xml:"Pty2,omitempty"`

	// Party that, in a settlement chain interacts with the party 2.
	Party3 *PartyIdentificationAndAccount106 `xml:"Pty3,omitempty"`

	// Party that, in a settlement chain interacts with the party 3.
	Party4 *PartyIdentificationAndAccount106 `xml:"Pty4,omitempty"`

	// Party that, in a settlement chain interacts with the party 4.
	Party5 *PartyIdentificationAndAccount106 `xml:"Pty5,omitempty"`
}

func (s *SettlementParties36) AddDepository() *PartyIdentification75 {
	s.Depository = new(PartyIdentification75)
	return s.Depository
}

func (s *SettlementParties36) AddParty1() *PartyIdentificationAndAccount106 {
	s.Party1 = new(PartyIdentificationAndAccount106)
	return s.Party1
}

func (s *SettlementParties36) AddParty2() *PartyIdentificationAndAccount106 {
	s.Party2 = new(PartyIdentificationAndAccount106)
	return s.Party2
}

func (s *SettlementParties36) AddParty3() *PartyIdentificationAndAccount106 {
	s.Party3 = new(PartyIdentificationAndAccount106)
	return s.Party3
}

func (s *SettlementParties36) AddParty4() *PartyIdentificationAndAccount106 {
	s.Party4 = new(PartyIdentificationAndAccount106)
	return s.Party4
}

func (s *SettlementParties36) AddParty5() *PartyIdentificationAndAccount106 {
	s.Party5 = new(PartyIdentificationAndAccount106)
	return s.Party5
}
