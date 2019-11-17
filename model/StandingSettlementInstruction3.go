package model

// Details of the standing settlement instruction to be applied.
type StandingSettlementInstruction3 struct {

	// Specifies what settlement standing instruction database is to be used to derive the settlement parties involved in the transaction.
	SettlementStandingInstructionDatabase *SettlementStandingInstructionDatabase1Choice `xml:"SttlmStgInstrDB"`

	// Identification of the buyer or seller in a standing settlement instruction enabling to derive the Standing Settlement Instruction.
	Counterparty *Counterparty3Choice `xml:"CtrPty"`

	// Vendor of the Settlement Standing Instruction database requested to be consulted.
	Vendor *PartyIdentification49Choice `xml:"Vndr,omitempty"`

	// Delivering parties, other than the seller, needed for deriving the standing settlement instruction (for example, depository) or provided for information purposes (for example, instructing party settlement chain).
	OtherDeliveringSettlementParties *SettlementParties10 `xml:"OthrDlvrgSttlmPties,omitempty"`

	// Receiving parties, other than the buyer, needed for deriving the standing settlement instruction (for example, depository) or provided for information purposes (for example, instructing party settlement chain).
	OtherReceivingSettlementParties *SettlementParties10 `xml:"OthrRcvgSttlmPties,omitempty"`
}

func (s *StandingSettlementInstruction3) AddSettlementStandingInstructionDatabase() *SettlementStandingInstructionDatabase1Choice {
	s.SettlementStandingInstructionDatabase = new(SettlementStandingInstructionDatabase1Choice)
	return s.SettlementStandingInstructionDatabase
}

func (s *StandingSettlementInstruction3) AddCounterparty() *Counterparty3Choice {
	s.Counterparty = new(Counterparty3Choice)
	return s.Counterparty
}

func (s *StandingSettlementInstruction3) AddVendor() *PartyIdentification49Choice {
	s.Vendor = new(PartyIdentification49Choice)
	return s.Vendor
}

func (s *StandingSettlementInstruction3) AddOtherDeliveringSettlementParties() *SettlementParties10 {
	s.OtherDeliveringSettlementParties = new(SettlementParties10)
	return s.OtherDeliveringSettlementParties
}

func (s *StandingSettlementInstruction3) AddOtherReceivingSettlementParties() *SettlementParties10 {
	s.OtherReceivingSettlementParties = new(SettlementParties10)
	return s.OtherReceivingSettlementParties
}
