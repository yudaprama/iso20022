package model

// Specifies settlement parties (delivering/receiving).
type SettlementParties4 struct {

	// First receiving party in the settlement chain. In a plain vanilla settlement, it is the Central Securities Depository where the receiving side of the transaction requests to receive the financial instrument.
	Depository *PartyIdentification35 `xml:"Dpstry"`

	// Party that interacts with the Depository.
	Party1 *PartyIdentificationAndAccount14 `xml:"Pty1"`

	// Party that interacts with the Party1.
	Party2 *PartyIdentificationAndAccount14 `xml:"Pty2,omitempty"`

	// Party that interacts with the Party2.
	Party3 *PartyIdentificationAndAccount14 `xml:"Pty3,omitempty"`
}

func (s *SettlementParties4) AddDepository() *PartyIdentification35 {
	s.Depository = new(PartyIdentification35)
	return s.Depository
}

func (s *SettlementParties4) AddParty1() *PartyIdentificationAndAccount14 {
	s.Party1 = new(PartyIdentificationAndAccount14)
	return s.Party1
}

func (s *SettlementParties4) AddParty2() *PartyIdentificationAndAccount14 {
	s.Party2 = new(PartyIdentificationAndAccount14)
	return s.Party2
}

func (s *SettlementParties4) AddParty3() *PartyIdentificationAndAccount14 {
	s.Party3 = new(PartyIdentificationAndAccount14)
	return s.Party3
}
