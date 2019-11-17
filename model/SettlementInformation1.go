package model

// Further information required for the settlement the transaction.
type SettlementInformation1 struct {

	// Method used to settle the (batch of) payment instructions.
	SettlementMethod *SettlementMethod1Code `xml:"SttlmMtd"`

	// A specific purpose account used to post debit and credit entries as a result of the transaction.
	SettlementAccount *CashAccount7 `xml:"SttlmAcct,omitempty"`

	// Specification of a pre-agreed offering between clearing agents or the channel through which the payment instruction is processed.
	ClearingSystem *ClearingSystemIdentification1Choice `xml:"ClrSys,omitempty"`

	// Agent through which the instructing agent will reimburse the instructed agent.
	//
	// Usage: If InstructingAgent and InstructedAgent have the same reimbursement agent, then only InstructingReimbursementAgent must be used.
	InstructingReimbursementAgent *BranchAndFinancialInstitutionIdentification3 `xml:"InstgRmbrsmntAgt,omitempty"`

	// Unambiguous identification of the account of the instructing reimbursement agent account at its servicing agent in the payment chain.
	InstructingReimbursementAgentAccount *CashAccount7 `xml:"InstgRmbrsmntAgtAcct,omitempty"`

	// Agent at which the instructed agent will be reimbursed.
	//
	// Usage: If InstructedReimbursementAgent contains a branch of the InstructedAgent, then the instructed agent will claim reimbursement from that branch/will be paid by that branch.
	//
	// Usage: If InstructingAgent and InstructedAgent have the same reimbursement agent, then only InstructingReimbursementAgent must be used.
	InstructedReimbursementAgent *BranchAndFinancialInstitutionIdentification3 `xml:"InstdRmbrsmntAgt,omitempty"`

	// Unambiguous identification of the account of the instructed reimbursement agent account at its servicing agent in the payment chain.
	InstructedReimbursementAgentAccount *CashAccount7 `xml:"InstdRmbrsmntAgtAcct,omitempty"`

	// Instructed agent's branch where the amount of money will be made available when different from the instructed reimbursement agent.
	ThirdReimbursementAgent *BranchAndFinancialInstitutionIdentification3 `xml:"ThrdRmbrsmntAgt,omitempty"`

	// Unambiguous identification of the account of the third reimbursement agent account at its servicing agent in the payment chain.
	ThirdReimbursementAgentAccount *CashAccount7 `xml:"ThrdRmbrsmntAgtAcct,omitempty"`
}

func (s *SettlementInformation1) SetSettlementMethod(value string) {
	s.SettlementMethod = (*SettlementMethod1Code)(&value)
}

func (s *SettlementInformation1) AddSettlementAccount() *CashAccount7 {
	s.SettlementAccount = new(CashAccount7)
	return s.SettlementAccount
}

func (s *SettlementInformation1) AddClearingSystem() *ClearingSystemIdentification1Choice {
	s.ClearingSystem = new(ClearingSystemIdentification1Choice)
	return s.ClearingSystem
}

func (s *SettlementInformation1) AddInstructingReimbursementAgent() *BranchAndFinancialInstitutionIdentification3 {
	s.InstructingReimbursementAgent = new(BranchAndFinancialInstitutionIdentification3)
	return s.InstructingReimbursementAgent
}

func (s *SettlementInformation1) AddInstructingReimbursementAgentAccount() *CashAccount7 {
	s.InstructingReimbursementAgentAccount = new(CashAccount7)
	return s.InstructingReimbursementAgentAccount
}

func (s *SettlementInformation1) AddInstructedReimbursementAgent() *BranchAndFinancialInstitutionIdentification3 {
	s.InstructedReimbursementAgent = new(BranchAndFinancialInstitutionIdentification3)
	return s.InstructedReimbursementAgent
}

func (s *SettlementInformation1) AddInstructedReimbursementAgentAccount() *CashAccount7 {
	s.InstructedReimbursementAgentAccount = new(CashAccount7)
	return s.InstructedReimbursementAgentAccount
}

func (s *SettlementInformation1) AddThirdReimbursementAgent() *BranchAndFinancialInstitutionIdentification3 {
	s.ThirdReimbursementAgent = new(BranchAndFinancialInstitutionIdentification3)
	return s.ThirdReimbursementAgent
}

func (s *SettlementInformation1) AddThirdReimbursementAgentAccount() *CashAccount7 {
	s.ThirdReimbursementAgentAccount = new(CashAccount7)
	return s.ThirdReimbursementAgentAccount
}
