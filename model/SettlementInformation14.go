package model

// Set of elements used to provide information on the settlement of the instruction.
type SettlementInformation14 struct {

	// Method used to settle the (batch of) payment instructions.
	SettlementMethod *SettlementMethod2Code `xml:"SttlmMtd"`

	// A specific purpose account used to post debit and credit entries as a result of the transaction.
	SettlementAccount *CashAccount16 `xml:"SttlmAcct,omitempty"`

	// Specification of a pre-agreed offering between clearing agents or the channel through which the payment instruction is processed.
	ClearingSystem *ClearingSystemIdentification3Choice `xml:"ClrSys,omitempty"`
}

func (s *SettlementInformation14) SetSettlementMethod(value string) {
	s.SettlementMethod = (*SettlementMethod2Code)(&value)
}

func (s *SettlementInformation14) AddSettlementAccount() *CashAccount16 {
	s.SettlementAccount = new(CashAccount16)
	return s.SettlementAccount
}

func (s *SettlementInformation14) AddClearingSystem() *ClearingSystemIdentification3Choice {
	s.ClearingSystem = new(ClearingSystemIdentification3Choice)
	return s.ClearingSystem
}
