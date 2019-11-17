package model

// Chain of parties involved in the settlement of a transaction, including receipts and deliveries, book transfers, treasury deals, or other activities, resulting in the movement of a security or amount of money from one account to another.
type SettlementParties12 struct {

	// First party in the settlement chain. In a plain vanilla settlement, it is the Central Securities Depository where the counterparty requests to receive the financial instrument or from where the counterparty delivers the financial instruments.
	Depository *PartyIdentification46 `xml:"Dpstry,omitempty"`

	// Party that, in a settlement chain interacts with the depository.
	Party1 *PartyIdentificationAndAccount43 `xml:"Pty1,omitempty"`

	// Party that, in a settlement chain interacts with the party 1.
	Party2 *PartyIdentificationAndAccount43 `xml:"Pty2,omitempty"`

	// Party that, in a settlement chain interacts with the party 2.
	Party3 *PartyIdentificationAndAccount43 `xml:"Pty3,omitempty"`

	// Party that, in a settlement chain interacts with the party 3.
	Party4 *PartyIdentificationAndAccount43 `xml:"Pty4,omitempty"`

	// Party that, in a settlement chain interacts with the party 4.
	Party5 *PartyIdentificationAndAccount43 `xml:"Pty5,omitempty"`
}

func (s *SettlementParties12) AddDepository() *PartyIdentification46 {
	s.Depository = new(PartyIdentification46)
	return s.Depository
}

func (s *SettlementParties12) AddParty1() *PartyIdentificationAndAccount43 {
	s.Party1 = new(PartyIdentificationAndAccount43)
	return s.Party1
}

func (s *SettlementParties12) AddParty2() *PartyIdentificationAndAccount43 {
	s.Party2 = new(PartyIdentificationAndAccount43)
	return s.Party2
}

func (s *SettlementParties12) AddParty3() *PartyIdentificationAndAccount43 {
	s.Party3 = new(PartyIdentificationAndAccount43)
	return s.Party3
}

func (s *SettlementParties12) AddParty4() *PartyIdentificationAndAccount43 {
	s.Party4 = new(PartyIdentificationAndAccount43)
	return s.Party4
}

func (s *SettlementParties12) AddParty5() *PartyIdentificationAndAccount43 {
	s.Party5 = new(PartyIdentificationAndAccount43)
	return s.Party5
}
