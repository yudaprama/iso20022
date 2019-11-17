package model

// Details of the standing settlement instruction to be applied.
type StandingSettlementInstruction11 struct {

	// Specifies what settlement standing instruction database is to be used to derive the settlement parties involved in the transaction.
	SettlementStandingInstructionDatabase *SettlementStandingInstructionDatabase4Choice `xml:"SttlmStgInstrDB"`

	// Identification of the buyer or seller in a standing settlement instruction enabling to derive the Standing Settlement Instruction.
	Counterparty *Counterparty9Choice `xml:"CtrPty"`

	// Vendor of the Settlement Standing Instruction database requested to be consulted.
	Vendor *PartyIdentification100 `xml:"Vndr,omitempty"`

	// Delivering parties, other than the seller, needed for deriving the standing settlement instruction (for example, depository) or provided for information purposes (for example, instructing party settlement chain).
	OtherDeliveringSettlementParties *SettlementParties36 `xml:"OthrDlvrgSttlmPties,omitempty"`

	// Receiving parties, other than the buyer, needed for deriving the standing settlement instruction (for example, depository) or provided for information purposes (for example, instructing party settlement chain).
	OtherReceivingSettlementParties *SettlementParties36 `xml:"OthrRcvgSttlmPties,omitempty"`
}

func (s *StandingSettlementInstruction11) AddSettlementStandingInstructionDatabase() *SettlementStandingInstructionDatabase4Choice {
	s.SettlementStandingInstructionDatabase = new(SettlementStandingInstructionDatabase4Choice)
	return s.SettlementStandingInstructionDatabase
}

func (s *StandingSettlementInstruction11) AddCounterparty() *Counterparty9Choice {
	s.Counterparty = new(Counterparty9Choice)
	return s.Counterparty
}

func (s *StandingSettlementInstruction11) AddVendor() *PartyIdentification100 {
	s.Vendor = new(PartyIdentification100)
	return s.Vendor
}

func (s *StandingSettlementInstruction11) AddOtherDeliveringSettlementParties() *SettlementParties36 {
	s.OtherDeliveringSettlementParties = new(SettlementParties36)
	return s.OtherDeliveringSettlementParties
}

func (s *StandingSettlementInstruction11) AddOtherReceivingSettlementParties() *SettlementParties36 {
	s.OtherReceivingSettlementParties = new(SettlementParties36)
	return s.OtherReceivingSettlementParties
}
