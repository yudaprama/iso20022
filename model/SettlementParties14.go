package model

// Chain of parties involved in the settlement of a transaction, including receipts and deliveries, book transfers, treasury deals, or other activities, resulting in the movement of a security or amount of money from one account to another.
type SettlementParties14 struct {

	// First party in the settlement chain. In a plain vanilla settlement, it is the Central Securities Depository where the counterparty requests to receive the financial instrument or from where the counterparty delivers the financial instruments.
	Depository *PartyIdentification48 `xml:"Dpstry,omitempty"`

	// Party that, in a settlement chain interacts with the depository.
	Party1 *PartyIdentificationAndAccount45 `xml:"Pty1,omitempty"`

	// Party that, in a settlement chain interacts with the party 1.
	Party2 *PartyIdentificationAndAccount45 `xml:"Pty2,omitempty"`

	// Party that, in a settlement chain interacts with the party 2.
	Party3 *PartyIdentificationAndAccount45 `xml:"Pty3,omitempty"`

	// Party that, in a settlement chain interacts with the party 3.
	Party4 *PartyIdentificationAndAccount45 `xml:"Pty4,omitempty"`

	// Party that, in a settlement chain interacts with the party 4.
	Party5 *PartyIdentificationAndAccount45 `xml:"Pty5,omitempty"`
}

func (s *SettlementParties14) AddDepository() *PartyIdentification48 {
	s.Depository = new(PartyIdentification48)
	return s.Depository
}

func (s *SettlementParties14) AddParty1() *PartyIdentificationAndAccount45 {
	s.Party1 = new(PartyIdentificationAndAccount45)
	return s.Party1
}

func (s *SettlementParties14) AddParty2() *PartyIdentificationAndAccount45 {
	s.Party2 = new(PartyIdentificationAndAccount45)
	return s.Party2
}

func (s *SettlementParties14) AddParty3() *PartyIdentificationAndAccount45 {
	s.Party3 = new(PartyIdentificationAndAccount45)
	return s.Party3
}

func (s *SettlementParties14) AddParty4() *PartyIdentificationAndAccount45 {
	s.Party4 = new(PartyIdentificationAndAccount45)
	return s.Party4
}

func (s *SettlementParties14) AddParty5() *PartyIdentificationAndAccount45 {
	s.Party5 = new(PartyIdentificationAndAccount45)
	return s.Party5
}
