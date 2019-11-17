package model

// Details of the standing settlement instruction to be applied.
type StandingSettlementInstruction4 struct {

	// Specifies what settlement standing instruction database is to be used to derive the settlement parties involved in the transaction.
	SettlementStandingInstructionDatabase *SettlementStandingInstructionDatabase1Choice `xml:"SttlmStgInstrDB"`

	// Identification of the buyer or seller in a standing settlement instruction enabling to derive the Standing Settlement Instruction.
	Counterparty *Counterparty4Choice `xml:"CtrPty"`

	// Vendor of the Settlement Standing Instruction database requested to be consulted.
	Vendor *PartyIdentification43Choice `xml:"Vndr,omitempty"`

	// Delivering parties, other than the seller, needed for deriving the standing settlement instruction (for example, depository) or provided for information purposes (for example, instructing party settlement chain).
	OtherDeliveringSettlementParties *SettlementParties11 `xml:"OthrDlvrgSttlmPties,omitempty"`

	// Receiving parties, other than the buyer, needed for deriving the standing settlement instruction (for example, depository) or provided for information purposes (for example, instructing party settlement chain).
	OtherReceivingSettlementParties *SettlementParties11 `xml:"OthrRcvgSttlmPties,omitempty"`
}

func (s *StandingSettlementInstruction4) AddSettlementStandingInstructionDatabase() *SettlementStandingInstructionDatabase1Choice {
	s.SettlementStandingInstructionDatabase = new(SettlementStandingInstructionDatabase1Choice)
	return s.SettlementStandingInstructionDatabase
}

func (s *StandingSettlementInstruction4) AddCounterparty() *Counterparty4Choice {
	s.Counterparty = new(Counterparty4Choice)
	return s.Counterparty
}

func (s *StandingSettlementInstruction4) AddVendor() *PartyIdentification43Choice {
	s.Vendor = new(PartyIdentification43Choice)
	return s.Vendor
}

func (s *StandingSettlementInstruction4) AddOtherDeliveringSettlementParties() *SettlementParties11 {
	s.OtherDeliveringSettlementParties = new(SettlementParties11)
	return s.OtherDeliveringSettlementParties
}

func (s *StandingSettlementInstruction4) AddOtherReceivingSettlementParties() *SettlementParties11 {
	s.OtherReceivingSettlementParties = new(SettlementParties11)
	return s.OtherReceivingSettlementParties
}
