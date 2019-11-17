package model

// Details of the standing settlement instruction to be applied.
type StandingSettlementInstruction5 struct {

	// Specifies what settlement standing instruction database is to be used to derive the settlement parties involved in the transaction.
	SettlementStandingInstructionDatabase *SettlementStandingInstructionDatabase1Choice `xml:"SttlmStgInstrDB"`

	// Identification of the buyer or seller in a standing settlement instruction enabling to derive the Standing Settlement Instruction.
	Counterparty *Counterparty5Choice `xml:"CtrPty"`

	// Vendor of the Settlement Standing Instruction database requested to be consulted.
	Vendor *PartyIdentification45Choice `xml:"Vndr,omitempty"`

	// Delivering parties, other than the seller, needed for deriving the standing settlement instruction (for example, depository) or provided for information purposes (for example, instructing party settlement chain).
	OtherDeliveringSettlementParties *SettlementParties14 `xml:"OthrDlvrgSttlmPties,omitempty"`

	// Receiving parties, other than the buyer, needed for deriving the standing settlement instruction (for example, depository) or provided for information purposes (for example, instructing party settlement chain).
	OtherReceivingSettlementParties *SettlementParties14 `xml:"OthrRcvgSttlmPties,omitempty"`
}

func (s *StandingSettlementInstruction5) AddSettlementStandingInstructionDatabase() *SettlementStandingInstructionDatabase1Choice {
	s.SettlementStandingInstructionDatabase = new(SettlementStandingInstructionDatabase1Choice)
	return s.SettlementStandingInstructionDatabase
}

func (s *StandingSettlementInstruction5) AddCounterparty() *Counterparty5Choice {
	s.Counterparty = new(Counterparty5Choice)
	return s.Counterparty
}

func (s *StandingSettlementInstruction5) AddVendor() *PartyIdentification45Choice {
	s.Vendor = new(PartyIdentification45Choice)
	return s.Vendor
}

func (s *StandingSettlementInstruction5) AddOtherDeliveringSettlementParties() *SettlementParties14 {
	s.OtherDeliveringSettlementParties = new(SettlementParties14)
	return s.OtherDeliveringSettlementParties
}

func (s *StandingSettlementInstruction5) AddOtherReceivingSettlementParties() *SettlementParties14 {
	s.OtherReceivingSettlementParties = new(SettlementParties14)
	return s.OtherReceivingSettlementParties
}
