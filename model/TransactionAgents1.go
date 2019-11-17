package model

// Set of elements providing information specific to the individual transaction(s) included in the message.
type TransactionAgents1 struct {

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification3 `xml:"DbtrAgt,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *BranchAndFinancialInstitutionIdentification3 `xml:"CdtrAgt,omitempty"`

	// Agent between the debtor agent and creditor agent.
	//
	// Usage: If more than one intermediary agent is present, then IntermediaryAgent1 identifies the agent between the debtor agent and the intermediary agent 2.
	IntermediaryAgent1 *BranchAndFinancialInstitutionIdentification3 `xml:"IntrmyAgt1,omitempty"`

	// Agent between the debtor agent and creditor agent.
	//
	// Usage: If more than two intermediary agents are present, then IntermediaryAgent2 identifies the agent between the intermediary agent 1 and the intermediary agent 3.
	IntermediaryAgent2 *BranchAndFinancialInstitutionIdentification3 `xml:"IntrmyAgt2,omitempty"`

	// Agent between the debtor agent and creditor agent.
	//
	// Usage: If IntermediaryAgent3 is present, then it identifies the agent between the intermediary agent 2 and the creditor agent.
	IntermediaryAgent3 *BranchAndFinancialInstitutionIdentification3 `xml:"IntrmyAgt3,omitempty"`

	// Party that receives securities from the delivering agent at the place of settlement, eg, central securities depository.
	// Can also be used in the context of treasury operations.
	ReceivingAgent *BranchAndFinancialInstitutionIdentification3 `xml:"RcvgAgt,omitempty"`

	// Party that delivers securities to the receiving agent at the place of settlement, eg, central securities depository.
	// Can also be used in the context of treasury operations.
	DeliveringAgent *BranchAndFinancialInstitutionIdentification3 `xml:"DlvrgAgt,omitempty"`

	// Legal entity that has the right to issue securities.
	IssuingAgent *BranchAndFinancialInstitutionIdentification3 `xml:"IssgAgt,omitempty"`

	// Place where settlement of the securities takes place.
	//
	// Note : this is typed by a financial institution identification - as this is the standard way of identifying eg securities settlement agent/central system.
	SettlementPlace *BranchAndFinancialInstitutionIdentification3 `xml:"SttlmPlc,omitempty"`

	// Proprietary agent related to the underlying transaction.
	Proprietary []*ProprietaryAgent1 `xml:"Prtry,omitempty"`
}

func (t *TransactionAgents1) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification3 {
	t.DebtorAgent = new(BranchAndFinancialInstitutionIdentification3)
	return t.DebtorAgent
}

func (t *TransactionAgents1) AddCreditorAgent() *BranchAndFinancialInstitutionIdentification3 {
	t.CreditorAgent = new(BranchAndFinancialInstitutionIdentification3)
	return t.CreditorAgent
}

func (t *TransactionAgents1) AddIntermediaryAgent1() *BranchAndFinancialInstitutionIdentification3 {
	t.IntermediaryAgent1 = new(BranchAndFinancialInstitutionIdentification3)
	return t.IntermediaryAgent1
}

func (t *TransactionAgents1) AddIntermediaryAgent2() *BranchAndFinancialInstitutionIdentification3 {
	t.IntermediaryAgent2 = new(BranchAndFinancialInstitutionIdentification3)
	return t.IntermediaryAgent2
}

func (t *TransactionAgents1) AddIntermediaryAgent3() *BranchAndFinancialInstitutionIdentification3 {
	t.IntermediaryAgent3 = new(BranchAndFinancialInstitutionIdentification3)
	return t.IntermediaryAgent3
}

func (t *TransactionAgents1) AddReceivingAgent() *BranchAndFinancialInstitutionIdentification3 {
	t.ReceivingAgent = new(BranchAndFinancialInstitutionIdentification3)
	return t.ReceivingAgent
}

func (t *TransactionAgents1) AddDeliveringAgent() *BranchAndFinancialInstitutionIdentification3 {
	t.DeliveringAgent = new(BranchAndFinancialInstitutionIdentification3)
	return t.DeliveringAgent
}

func (t *TransactionAgents1) AddIssuingAgent() *BranchAndFinancialInstitutionIdentification3 {
	t.IssuingAgent = new(BranchAndFinancialInstitutionIdentification3)
	return t.IssuingAgent
}

func (t *TransactionAgents1) AddSettlementPlace() *BranchAndFinancialInstitutionIdentification3 {
	t.SettlementPlace = new(BranchAndFinancialInstitutionIdentification3)
	return t.SettlementPlace
}

func (t *TransactionAgents1) AddProprietary() *ProprietaryAgent1 {
	newValue := new(ProprietaryAgent1)
	t.Proprietary = append(t.Proprietary, newValue)
	return newValue
}
