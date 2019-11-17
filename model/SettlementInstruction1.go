package model

// Provides further details on the settlement of the instruction.
type SettlementInstruction1 struct {

	// Method used to settle the (batch of) payment instructions.
	SettlementMethod *SettlementMethod1Code `xml:"SttlmMtd"`

	// A specific purpose account used to post debit and credit entries as a result of the transaction.
	SettlementAccount *CashAccount24 `xml:"SttlmAcct,omitempty"`

	// Specification of a pre-agreed offering between clearing agents or the channel through which the payment instruction is processed.
	ClearingSystem *ClearingSystemIdentification3Choice `xml:"ClrSys,omitempty"`

	// Agent through which the instructing agent will reimburse the instructed agent.
	//
	// Usage: If InstructingAgent and InstructedAgent have the same reimbursement agent, then only InstructingReimbursementAgent must be used.
	InstructingReimbursementAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstgRmbrsmntAgt,omitempty"`

	// Unambiguous identification of the account of the instructing reimbursement agent account at its servicing agent in the payment chain.
	InstructingReimbursementAgentAccount *CashAccount24 `xml:"InstgRmbrsmntAgtAcct,omitempty"`

	// Agent at which the instructed agent will be reimbursed.
	// Usage: If InstructedReimbursementAgent contains a branch of the InstructedAgent, then the party in InstructedAgent will claim reimbursement from that branch/will be paid by that branch.
	// Usage: If InstructingAgent and InstructedAgent have the same reimbursement agent, then only InstructingReimbursementAgent must be used.
	InstructedReimbursementAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstdRmbrsmntAgt,omitempty"`

	// Unambiguous identification of the account of the instructed reimbursement agent account at its servicing agent in the payment chain.
	InstructedReimbursementAgentAccount *CashAccount24 `xml:"InstdRmbrsmntAgtAcct,omitempty"`

	// Agent at which the instructed agent will be reimbursed.
	// Usage: If ThirdReimbursementAgent contains a branch of the InstructedAgent, then the party in InstructedAgent will claim reimbursement from that branch/will be paid by that branch.
	ThirdReimbursementAgent *BranchAndFinancialInstitutionIdentification5 `xml:"ThrdRmbrsmntAgt,omitempty"`

	// Unambiguous identification of the account of the third reimbursement agent account at its servicing agent in the payment chain.
	ThirdReimbursementAgentAccount *CashAccount24 `xml:"ThrdRmbrsmntAgtAcct,omitempty"`
}

func (s *SettlementInstruction1) SetSettlementMethod(value string) {
	s.SettlementMethod = (*SettlementMethod1Code)(&value)
}

func (s *SettlementInstruction1) AddSettlementAccount() *CashAccount24 {
	s.SettlementAccount = new(CashAccount24)
	return s.SettlementAccount
}

func (s *SettlementInstruction1) AddClearingSystem() *ClearingSystemIdentification3Choice {
	s.ClearingSystem = new(ClearingSystemIdentification3Choice)
	return s.ClearingSystem
}

func (s *SettlementInstruction1) AddInstructingReimbursementAgent() *BranchAndFinancialInstitutionIdentification5 {
	s.InstructingReimbursementAgent = new(BranchAndFinancialInstitutionIdentification5)
	return s.InstructingReimbursementAgent
}

func (s *SettlementInstruction1) AddInstructingReimbursementAgentAccount() *CashAccount24 {
	s.InstructingReimbursementAgentAccount = new(CashAccount24)
	return s.InstructingReimbursementAgentAccount
}

func (s *SettlementInstruction1) AddInstructedReimbursementAgent() *BranchAndFinancialInstitutionIdentification5 {
	s.InstructedReimbursementAgent = new(BranchAndFinancialInstitutionIdentification5)
	return s.InstructedReimbursementAgent
}

func (s *SettlementInstruction1) AddInstructedReimbursementAgentAccount() *CashAccount24 {
	s.InstructedReimbursementAgentAccount = new(CashAccount24)
	return s.InstructedReimbursementAgentAccount
}

func (s *SettlementInstruction1) AddThirdReimbursementAgent() *BranchAndFinancialInstitutionIdentification5 {
	s.ThirdReimbursementAgent = new(BranchAndFinancialInstitutionIdentification5)
	return s.ThirdReimbursementAgent
}

func (s *SettlementInstruction1) AddThirdReimbursementAgentAccount() *CashAccount24 {
	s.ThirdReimbursementAgentAccount = new(CashAccount24)
	return s.ThirdReimbursementAgentAccount
}
