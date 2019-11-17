package model

// Provides further details on the settlement of the instruction.
type SettlementInstruction3 struct {

	// Agent through which the instructing agent will reimburse the instructed agent.
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
}

func (s *SettlementInstruction3) AddInstructingReimbursementAgent() *BranchAndFinancialInstitutionIdentification5 {
	s.InstructingReimbursementAgent = new(BranchAndFinancialInstitutionIdentification5)
	return s.InstructingReimbursementAgent
}

func (s *SettlementInstruction3) AddInstructingReimbursementAgentAccount() *CashAccount24 {
	s.InstructingReimbursementAgentAccount = new(CashAccount24)
	return s.InstructingReimbursementAgentAccount
}

func (s *SettlementInstruction3) AddInstructedReimbursementAgent() *BranchAndFinancialInstitutionIdentification5 {
	s.InstructedReimbursementAgent = new(BranchAndFinancialInstitutionIdentification5)
	return s.InstructedReimbursementAgent
}

func (s *SettlementInstruction3) AddInstructedReimbursementAgentAccount() *CashAccount24 {
	s.InstructedReimbursementAgentAccount = new(CashAccount24)
	return s.InstructedReimbursementAgentAccount
}
