package model

// Set of elements used to provide information on the agents specific to the individual transaction.
type TransactionAgents2 struct {

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification4 `xml:"DbtrAgt,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *BranchAndFinancialInstitutionIdentification4 `xml:"CdtrAgt,omitempty"`

	// Agent between the debtor's agent and the creditor's agent.
	//
	// Usage: If more than one intermediary agent is present, then IntermediaryAgent1 identifies the agent between the DebtorAgent and the IntermediaryAgent2.
	IntermediaryAgent1 *BranchAndFinancialInstitutionIdentification4 `xml:"IntrmyAgt1,omitempty"`

	// Agent between the debtor's agent and the creditor's agent.
	//
	// Usage: If more than two intermediary agents are present, then IntermediaryAgent2 identifies the agent between the IntermediaryAgent1 and the IntermediaryAgent3.
	IntermediaryAgent2 *BranchAndFinancialInstitutionIdentification4 `xml:"IntrmyAgt2,omitempty"`

	// Agent between the debtor's agent and the creditor's agent.
	//
	// Usage: If IntermediaryAgent3 is present, then it identifies the agent between the IntermediaryAgent 2 and the CreditorAgent.
	IntermediaryAgent3 *BranchAndFinancialInstitutionIdentification4 `xml:"IntrmyAgt3,omitempty"`

	// Party that receives securities from the delivering agent at the place of settlement, such as central securities depository.
	// Can also be used in the context of treasury operations.
	ReceivingAgent *BranchAndFinancialInstitutionIdentification4 `xml:"RcvgAgt,omitempty"`

	// Party that delivers securities to the receiving agent at the place of settlement, such as a central securities depository.
	// Can also be used in the context of treasury operations.
	DeliveringAgent *BranchAndFinancialInstitutionIdentification4 `xml:"DlvrgAgt,omitempty"`

	// Legal entity that has the right to issue securities.
	IssuingAgent *BranchAndFinancialInstitutionIdentification4 `xml:"IssgAgt,omitempty"`

	// Place where settlement of the securities takes place.
	// Usage: This is typed by a financial institution identification as this is the standard way to identify a securities settlement agent/central system.
	SettlementPlace *BranchAndFinancialInstitutionIdentification4 `xml:"SttlmPlc,omitempty"`

	// Proprietary agent related to the underlying transaction.
	Proprietary []*ProprietaryAgent2 `xml:"Prtry,omitempty"`
}

func (t *TransactionAgents2) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification4 {
	t.DebtorAgent = new(BranchAndFinancialInstitutionIdentification4)
	return t.DebtorAgent
}

func (t *TransactionAgents2) AddCreditorAgent() *BranchAndFinancialInstitutionIdentification4 {
	t.CreditorAgent = new(BranchAndFinancialInstitutionIdentification4)
	return t.CreditorAgent
}

func (t *TransactionAgents2) AddIntermediaryAgent1() *BranchAndFinancialInstitutionIdentification4 {
	t.IntermediaryAgent1 = new(BranchAndFinancialInstitutionIdentification4)
	return t.IntermediaryAgent1
}

func (t *TransactionAgents2) AddIntermediaryAgent2() *BranchAndFinancialInstitutionIdentification4 {
	t.IntermediaryAgent2 = new(BranchAndFinancialInstitutionIdentification4)
	return t.IntermediaryAgent2
}

func (t *TransactionAgents2) AddIntermediaryAgent3() *BranchAndFinancialInstitutionIdentification4 {
	t.IntermediaryAgent3 = new(BranchAndFinancialInstitutionIdentification4)
	return t.IntermediaryAgent3
}

func (t *TransactionAgents2) AddReceivingAgent() *BranchAndFinancialInstitutionIdentification4 {
	t.ReceivingAgent = new(BranchAndFinancialInstitutionIdentification4)
	return t.ReceivingAgent
}

func (t *TransactionAgents2) AddDeliveringAgent() *BranchAndFinancialInstitutionIdentification4 {
	t.DeliveringAgent = new(BranchAndFinancialInstitutionIdentification4)
	return t.DeliveringAgent
}

func (t *TransactionAgents2) AddIssuingAgent() *BranchAndFinancialInstitutionIdentification4 {
	t.IssuingAgent = new(BranchAndFinancialInstitutionIdentification4)
	return t.IssuingAgent
}

func (t *TransactionAgents2) AddSettlementPlace() *BranchAndFinancialInstitutionIdentification4 {
	t.SettlementPlace = new(BranchAndFinancialInstitutionIdentification4)
	return t.SettlementPlace
}

func (t *TransactionAgents2) AddProprietary() *ProprietaryAgent2 {
	newValue := new(ProprietaryAgent2)
	t.Proprietary = append(t.Proprietary, newValue)
	return newValue
}
