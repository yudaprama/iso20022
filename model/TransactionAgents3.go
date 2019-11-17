package model

// Provides further details on the agents specific to the individual transaction.
type TransactionAgents3 struct {

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"DbtrAgt,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"CdtrAgt,omitempty"`

	// Agent between the debtor's agent and the creditor's agent.
	//
	// Usage: If more than one intermediary agent is present, then IntermediaryAgent1 identifies the agent between the DebtorAgent and the IntermediaryAgent2.
	IntermediaryAgent1 *BranchAndFinancialInstitutionIdentification5 `xml:"IntrmyAgt1,omitempty"`

	// Agent between the debtor's agent and the creditor's agent.
	//
	// Usage: If more than two intermediary agents are present, then IntermediaryAgent2 identifies the agent between the IntermediaryAgent1 and the IntermediaryAgent3.
	IntermediaryAgent2 *BranchAndFinancialInstitutionIdentification5 `xml:"IntrmyAgt2,omitempty"`

	// Agent between the debtor's agent and the creditor's agent.
	//
	// Usage: If IntermediaryAgent3 is present, then it identifies the agent between the IntermediaryAgent 2 and the CreditorAgent.
	IntermediaryAgent3 *BranchAndFinancialInstitutionIdentification5 `xml:"IntrmyAgt3,omitempty"`

	// Party that receives securities from the delivering agent at the place of settlement, such as central securities depository.
	// Can also be used in the context of treasury operations.
	ReceivingAgent *BranchAndFinancialInstitutionIdentification5 `xml:"RcvgAgt,omitempty"`

	// Party that delivers securities to the receiving agent at the place of settlement, such as a central securities depository.
	// Can also be used in the context of treasury operations.
	DeliveringAgent *BranchAndFinancialInstitutionIdentification5 `xml:"DlvrgAgt,omitempty"`

	// Legal entity that has the right to issue securities.
	IssuingAgent *BranchAndFinancialInstitutionIdentification5 `xml:"IssgAgt,omitempty"`

	// Place where settlement of the securities takes place.
	// Usage: This is typed by a financial institution identification as this is the standard way to identify a securities settlement agent/central system.
	SettlementPlace *BranchAndFinancialInstitutionIdentification5 `xml:"SttlmPlc,omitempty"`

	// Proprietary agent related to the underlying transaction.
	Proprietary []*ProprietaryAgent3 `xml:"Prtry,omitempty"`
}

func (t *TransactionAgents3) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification5 {
	t.DebtorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return t.DebtorAgent
}

func (t *TransactionAgents3) AddCreditorAgent() *BranchAndFinancialInstitutionIdentification5 {
	t.CreditorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return t.CreditorAgent
}

func (t *TransactionAgents3) AddIntermediaryAgent1() *BranchAndFinancialInstitutionIdentification5 {
	t.IntermediaryAgent1 = new(BranchAndFinancialInstitutionIdentification5)
	return t.IntermediaryAgent1
}

func (t *TransactionAgents3) AddIntermediaryAgent2() *BranchAndFinancialInstitutionIdentification5 {
	t.IntermediaryAgent2 = new(BranchAndFinancialInstitutionIdentification5)
	return t.IntermediaryAgent2
}

func (t *TransactionAgents3) AddIntermediaryAgent3() *BranchAndFinancialInstitutionIdentification5 {
	t.IntermediaryAgent3 = new(BranchAndFinancialInstitutionIdentification5)
	return t.IntermediaryAgent3
}

func (t *TransactionAgents3) AddReceivingAgent() *BranchAndFinancialInstitutionIdentification5 {
	t.ReceivingAgent = new(BranchAndFinancialInstitutionIdentification5)
	return t.ReceivingAgent
}

func (t *TransactionAgents3) AddDeliveringAgent() *BranchAndFinancialInstitutionIdentification5 {
	t.DeliveringAgent = new(BranchAndFinancialInstitutionIdentification5)
	return t.DeliveringAgent
}

func (t *TransactionAgents3) AddIssuingAgent() *BranchAndFinancialInstitutionIdentification5 {
	t.IssuingAgent = new(BranchAndFinancialInstitutionIdentification5)
	return t.IssuingAgent
}

func (t *TransactionAgents3) AddSettlementPlace() *BranchAndFinancialInstitutionIdentification5 {
	t.SettlementPlace = new(BranchAndFinancialInstitutionIdentification5)
	return t.SettlementPlace
}

func (t *TransactionAgents3) AddProprietary() *ProprietaryAgent3 {
	newValue := new(ProprietaryAgent3)
	t.Proprietary = append(t.Proprietary, newValue)
	return newValue
}
