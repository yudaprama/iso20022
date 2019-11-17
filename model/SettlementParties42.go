package model

// Specifies settlement parties (delivering/receiving).
type SettlementParties42 struct {

	// First receiving party in the settlement chain. In a plain vanilla settlement, it is the central securities depository where the receiving side of the transaction requests to receive the financial instrument.
	Depository *PartyIdentification92 `xml:"Dpstry,omitempty"`

	// Party that interacts with the depository.
	Party1 *PartyIdentificationAndAccount122 `xml:"Pty1,omitempty"`

	// Party that interacts with the party1.
	Party2 *PartyIdentificationAndAccount122 `xml:"Pty2,omitempty"`

	// Party that interacts with the party2.
	Party3 *PartyIdentificationAndAccount122 `xml:"Pty3,omitempty"`
}

func (s *SettlementParties42) AddDepository() *PartyIdentification92 {
	s.Depository = new(PartyIdentification92)
	return s.Depository
}

func (s *SettlementParties42) AddParty1() *PartyIdentificationAndAccount122 {
	s.Party1 = new(PartyIdentificationAndAccount122)
	return s.Party1
}

func (s *SettlementParties42) AddParty2() *PartyIdentificationAndAccount122 {
	s.Party2 = new(PartyIdentificationAndAccount122)
	return s.Party2
}

func (s *SettlementParties42) AddParty3() *PartyIdentificationAndAccount122 {
	s.Party3 = new(PartyIdentificationAndAccount122)
	return s.Party3
}
