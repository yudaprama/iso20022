package model

// Set of elements used to provide information on the settlement of the instruction.
type SettlementInformation15 struct {

	// Agent through which the instructing agent will reimburse the instructed agent.
	// Usage: If InstructingAgent and InstructedAgent have the same reimbursement agent, then only InstructingReimbursementAgent must be used.
	InstructingReimbursementAgent *BranchAndFinancialInstitutionIdentification4 `xml:"InstgRmbrsmntAgt,omitempty"`

	// Unambiguous identification of the account of the instructing reimbursement agent account at its servicing agent in the payment chain.
	InstructingReimbursementAgentAccount *CashAccount16 `xml:"InstgRmbrsmntAgtAcct,omitempty"`

	// Agent at which the instructed agent will be reimbursed.
	// Usage: If InstructedReimbursementAgent contains a branch of the InstructedAgent, then the party in InstructedAgent will claim reimbursement from that branch/will be paid by that branch.
	// Usage: If InstructingAgent and InstructedAgent have the same reimbursement agent, then only InstructingReimbursementAgent must be used.
	InstructedReimbursementAgent *BranchAndFinancialInstitutionIdentification4 `xml:"InstdRmbrsmntAgt,omitempty"`

	// Unambiguous identification of the account of the instructed reimbursement agent account at its servicing agent in the payment chain.
	InstructedReimbursementAgentAccount *CashAccount16 `xml:"InstdRmbrsmntAgtAcct,omitempty"`
}

func (s *SettlementInformation15) AddInstructingReimbursementAgent() *BranchAndFinancialInstitutionIdentification4 {
	s.InstructingReimbursementAgent = new(BranchAndFinancialInstitutionIdentification4)
	return s.InstructingReimbursementAgent
}

func (s *SettlementInformation15) AddInstructingReimbursementAgentAccount() *CashAccount16 {
	s.InstructingReimbursementAgentAccount = new(CashAccount16)
	return s.InstructingReimbursementAgentAccount
}

func (s *SettlementInformation15) AddInstructedReimbursementAgent() *BranchAndFinancialInstitutionIdentification4 {
	s.InstructedReimbursementAgent = new(BranchAndFinancialInstitutionIdentification4)
	return s.InstructedReimbursementAgent
}

func (s *SettlementInformation15) AddInstructedReimbursementAgentAccount() *CashAccount16 {
	s.InstructedReimbursementAgentAccount = new(CashAccount16)
	return s.InstructedReimbursementAgentAccount
}
