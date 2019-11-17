package model

// Chain of parties involved in the settlement of a transaction, including receipts and deliveries, book transfers, treasury deals, or other activities, resulting in the movement of a security or amount of money from one account to another.
type SettlementParties2 struct {

	// First party in the settlement chain. In a plain vanilla settlement, it is the Central Securities Depository where the counterparty requests to receive the financial instrument or from where the counterparty delivers the financial instruments.
	Depository *PartyIdentification36 `xml:"Dpstry,omitempty"`

	// Party that, in a settlement chain interacts with the depository.
	Party1 *PartyIdentificationAndAccount16 `xml:"Pty1,omitempty"`

	// Party that, in a settlement chain interacts with the party 1.
	Party2 *PartyIdentificationAndAccount16 `xml:"Pty2,omitempty"`

	// Party that, in a settlement chain interacts with the party 2.
	Party3 *PartyIdentificationAndAccount16 `xml:"Pty3,omitempty"`

	// Party that, in a settlement chain interacts with the party 3.
	Party4 *PartyIdentificationAndAccount16 `xml:"Pty4,omitempty"`

	// Party that, in a settlement chain interacts with the party 4.
	Party5 *PartyIdentificationAndAccount16 `xml:"Pty5,omitempty"`
}

func (s *SettlementParties2) AddDepository() *PartyIdentification36 {
	s.Depository = new(PartyIdentification36)
	return s.Depository
}

func (s *SettlementParties2) AddParty1() *PartyIdentificationAndAccount16 {
	s.Party1 = new(PartyIdentificationAndAccount16)
	return s.Party1
}

func (s *SettlementParties2) AddParty2() *PartyIdentificationAndAccount16 {
	s.Party2 = new(PartyIdentificationAndAccount16)
	return s.Party2
}

func (s *SettlementParties2) AddParty3() *PartyIdentificationAndAccount16 {
	s.Party3 = new(PartyIdentificationAndAccount16)
	return s.Party3
}

func (s *SettlementParties2) AddParty4() *PartyIdentificationAndAccount16 {
	s.Party4 = new(PartyIdentificationAndAccount16)
	return s.Party4
}

func (s *SettlementParties2) AddParty5() *PartyIdentificationAndAccount16 {
	s.Party5 = new(PartyIdentificationAndAccount16)
	return s.Party5
}
