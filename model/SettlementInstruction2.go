package model

// Provides further details on the settlement of the instruction.
type SettlementInstruction2 struct {

	// Method used to settle the (batch of) payment instructions.
	SettlementMethod *SettlementMethod2Code `xml:"SttlmMtd"`

	// A specific purpose account used to post debit and credit entries as a result of the transaction.
	SettlementAccount *CashAccount24 `xml:"SttlmAcct,omitempty"`

	// Specification of a pre-agreed offering between clearing agents or the channel through which the payment instruction is processed.
	ClearingSystem *ClearingSystemIdentification3Choice `xml:"ClrSys,omitempty"`
}

func (s *SettlementInstruction2) SetSettlementMethod(value string) {
	s.SettlementMethod = (*SettlementMethod2Code)(&value)
}

func (s *SettlementInstruction2) AddSettlementAccount() *CashAccount24 {
	s.SettlementAccount = new(CashAccount24)
	return s.SettlementAccount
}

func (s *SettlementInstruction2) AddClearingSystem() *ClearingSystemIdentification3Choice {
	s.ClearingSystem = new(ClearingSystemIdentification3Choice)
	return s.ClearingSystem
}
