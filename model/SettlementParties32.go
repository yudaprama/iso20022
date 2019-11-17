package model

// Chain of parties involved in the settlement of a transaction, including receipts and deliveries, book transfers, treasury deals, or other activities, resulting in the movement of a security or amount of money from one account to another.
type SettlementParties32 struct {

	// First party in the settlement chain. In a plain vanilla settlement, it is the central securities depository where the counterparty requests to receive the financial instrument or from where the counterparty delivers the financial instruments.
	Depository *PartyIdentification63 `xml:"Dpstry"`

	// Party that, in a settlement chain, interacts with the depository. This may also be known as the “local agent”, “sub-custodian”, “receiving agent” or “delivering agent”.
	Party1 *PartyIdentificationAndAccount95 `xml:"Pty1,omitempty"`

	// Party that, in a settlement chain, interacts with party 1. This may also be known as the “investment manager” or “custodian”.
	Party2 *PartyIdentificationAndAccount95 `xml:"Pty2,omitempty"`

	// Party that, in a settlement chain, interacts with party 2.
	Party3 *PartyIdentificationAndAccount95 `xml:"Pty3,omitempty"`

	// Party that, in a settlement chain, interacts with party 3.
	Party4 *PartyIdentificationAndAccount95 `xml:"Pty4,omitempty"`

	// Party that, in a settlement chain, interacts with party 4.
	Party5 *PartyIdentificationAndAccount95 `xml:"Pty5,omitempty"`
}

func (s *SettlementParties32) AddDepository() *PartyIdentification63 {
	s.Depository = new(PartyIdentification63)
	return s.Depository
}

func (s *SettlementParties32) AddParty1() *PartyIdentificationAndAccount95 {
	s.Party1 = new(PartyIdentificationAndAccount95)
	return s.Party1
}

func (s *SettlementParties32) AddParty2() *PartyIdentificationAndAccount95 {
	s.Party2 = new(PartyIdentificationAndAccount95)
	return s.Party2
}

func (s *SettlementParties32) AddParty3() *PartyIdentificationAndAccount95 {
	s.Party3 = new(PartyIdentificationAndAccount95)
	return s.Party3
}

func (s *SettlementParties32) AddParty4() *PartyIdentificationAndAccount95 {
	s.Party4 = new(PartyIdentificationAndAccount95)
	return s.Party4
}

func (s *SettlementParties32) AddParty5() *PartyIdentificationAndAccount95 {
	s.Party5 = new(PartyIdentificationAndAccount95)
	return s.Party5
}
