package model

// Details of the standing settlement instruction to be applied.
type StandingSettlementInstruction12 struct {

	// Specifies what settlement standing instruction database is to be used to derive the settlement parties involved in the transaction.
	SettlementStandingInstructionDatabase *SettlementStandingInstructionDatabase5Choice `xml:"SttlmStgInstrDB"`

	// Identification of the buyer or seller in a standing settlement instruction enabling to derive the Standing Settlement Instruction.
	Counterparty *Counterparty10Choice `xml:"CtrPty"`

	// Vendor of the Settlement Standing Instruction database requested to be consulted.
	Vendor *PartyIdentification111 `xml:"Vndr,omitempty"`

	// Delivering parties, other than the seller, needed for deriving the standing settlement instruction (for example, depository) or provided for information purposes (for example, instructing party settlement chain).
	OtherDeliveringSettlementParties *SettlementParties44 `xml:"OthrDlvrgSttlmPties,omitempty"`

	// Receiving parties, other than the buyer, needed for deriving the standing settlement instruction (for example, depository) or provided for information purposes (for example, instructing party settlement chain).
	OtherReceivingSettlementParties *SettlementParties44 `xml:"OthrRcvgSttlmPties,omitempty"`
}

func (s *StandingSettlementInstruction12) AddSettlementStandingInstructionDatabase() *SettlementStandingInstructionDatabase5Choice {
	s.SettlementStandingInstructionDatabase = new(SettlementStandingInstructionDatabase5Choice)
	return s.SettlementStandingInstructionDatabase
}

func (s *StandingSettlementInstruction12) AddCounterparty() *Counterparty10Choice {
	s.Counterparty = new(Counterparty10Choice)
	return s.Counterparty
}

func (s *StandingSettlementInstruction12) AddVendor() *PartyIdentification111 {
	s.Vendor = new(PartyIdentification111)
	return s.Vendor
}

func (s *StandingSettlementInstruction12) AddOtherDeliveringSettlementParties() *SettlementParties44 {
	s.OtherDeliveringSettlementParties = new(SettlementParties44)
	return s.OtherDeliveringSettlementParties
}

func (s *StandingSettlementInstruction12) AddOtherReceivingSettlementParties() *SettlementParties44 {
	s.OtherReceivingSettlementParties = new(SettlementParties44)
	return s.OtherReceivingSettlementParties
}
