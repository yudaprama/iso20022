package model

// Details of the standing settlement instruction to be applied.
type StandingSettlementInstruction9 struct {

	// Specifies what settlement standing instruction database is to be used to derive the settlement parties involved in the transaction.
	SettlementStandingInstructionDatabase *SettlementStandingInstructionDatabase3Choice `xml:"SttlmStgInstrDB"`

	// Vendor of the Settlement Standing Instruction database requested to be consulted.
	Vendor *PartyIdentification32Choice `xml:"Vndr,omitempty"`

	// Delivering parties, other than the seller, needed for deriving the standing settlement instruction (for example, depository) or provided for information purposes (for example, instructing party settlement chain).
	OtherDeliveringSettlementParties *SettlementParties23 `xml:"OthrDlvrgSttlmPties,omitempty"`

	// Receiving parties, other than the buyer, needed for deriving the standing settlement instruction (for example, depository) or provided for information purposes (for example, instructing party settlement chain).
	OtherReceivingSettlementParties *SettlementParties23 `xml:"OthrRcvgSttlmPties,omitempty"`
}

func (s *StandingSettlementInstruction9) AddSettlementStandingInstructionDatabase() *SettlementStandingInstructionDatabase3Choice {
	s.SettlementStandingInstructionDatabase = new(SettlementStandingInstructionDatabase3Choice)
	return s.SettlementStandingInstructionDatabase
}

func (s *StandingSettlementInstruction9) AddVendor() *PartyIdentification32Choice {
	s.Vendor = new(PartyIdentification32Choice)
	return s.Vendor
}

func (s *StandingSettlementInstruction9) AddOtherDeliveringSettlementParties() *SettlementParties23 {
	s.OtherDeliveringSettlementParties = new(SettlementParties23)
	return s.OtherDeliveringSettlementParties
}

func (s *StandingSettlementInstruction9) AddOtherReceivingSettlementParties() *SettlementParties23 {
	s.OtherReceivingSettlementParties = new(SettlementParties23)
	return s.OtherReceivingSettlementParties
}
