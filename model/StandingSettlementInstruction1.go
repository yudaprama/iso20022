package model

// Details of the standing settlement instruction to be applied.
type StandingSettlementInstruction1 struct {

	// Specifies what settlement standing instruction database is to be used to derive the settlement parties involved in the transaction.
	SettlementStandingInstructionDatabase *SettlementStandingInstructionDatabase1Choice `xml:"SttlmStgInstrDB"`

	// Identification of the buyer or seller in a standing settlement instruction enabling to derive the Standing Settlement Instruction.
	Counterparty *Counterparty1Choice `xml:"CtrPty"`

	// Vendor of the Settlement Standing Instruction database requested to be consulted.
	Vendor *PartyIdentification10Choice `xml:"Vndr,omitempty"`

	// Delivering parties, other than the seller, needed for deriving the standing settlement instruction (for example, depository) or provided for information purposes (for example, instructing party settlement chain).
	OtherDeliveringSettlementParties *SettlementParties5 `xml:"OthrDlvrgSttlmPties,omitempty"`

	// Receiving parties, other than the buyer, needed for deriving the standing settlement instruction (for example, depository) or provided for information purposes (for example, instructing party settlement chain).
	OtherReceivingSettlementParties *SettlementParties5 `xml:"OthrRcvgSttlmPties,omitempty"`
}

func (s *StandingSettlementInstruction1) AddSettlementStandingInstructionDatabase() *SettlementStandingInstructionDatabase1Choice {
	s.SettlementStandingInstructionDatabase = new(SettlementStandingInstructionDatabase1Choice)
	return s.SettlementStandingInstructionDatabase
}

func (s *StandingSettlementInstruction1) AddCounterparty() *Counterparty1Choice {
	s.Counterparty = new(Counterparty1Choice)
	return s.Counterparty
}

func (s *StandingSettlementInstruction1) AddVendor() *PartyIdentification10Choice {
	s.Vendor = new(PartyIdentification10Choice)
	return s.Vendor
}

func (s *StandingSettlementInstruction1) AddOtherDeliveringSettlementParties() *SettlementParties5 {
	s.OtherDeliveringSettlementParties = new(SettlementParties5)
	return s.OtherDeliveringSettlementParties
}

func (s *StandingSettlementInstruction1) AddOtherReceivingSettlementParties() *SettlementParties5 {
	s.OtherReceivingSettlementParties = new(SettlementParties5)
	return s.OtherReceivingSettlementParties
}
