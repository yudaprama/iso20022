package model

// Further information required for the settlement the transaction.
type SettlementInformation2 struct {

	// Method used to settle the (batch of) payment instructions.
	SettlementMethod *SettlementMethod2Code `xml:"SttlmMtd"`

	// A specific purpose account used to post debit and credit entries as a result of the transaction.
	SettlementAccount *CashAccount7 `xml:"SttlmAcct,omitempty"`

	// Specification of a pre-agreed offering between clearing agents or the channel through which the payment instruction is processed.
	ClearingSystem *ClearingSystemIdentification1Choice `xml:"ClrSys,omitempty"`
}

func (s *SettlementInformation2) SetSettlementMethod(value string) {
	s.SettlementMethod = (*SettlementMethod2Code)(&value)
}

func (s *SettlementInformation2) AddSettlementAccount() *CashAccount7 {
	s.SettlementAccount = new(CashAccount7)
	return s.SettlementAccount
}

func (s *SettlementInformation2) AddClearingSystem() *ClearingSystemIdentification1Choice {
	s.ClearingSystem = new(ClearingSystemIdentification1Choice)
	return s.ClearingSystem
}
