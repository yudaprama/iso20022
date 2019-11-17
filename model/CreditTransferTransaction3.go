package model

// Provides further details specific to the individual transaction(s) included in the message.
type CreditTransferTransaction3 struct {

	// Ultimate party that owes an amount of money to the (ultimate) creditor.
	UltimateDebtor *PartyIdentification43 `xml:"UltmtDbtr,omitempty"`

	// Party that initiates the payment.
	// Usage: This can be either the debtor or a party that initiates the credit transfer on behalf of the debtor.
	InitiatingParty *PartyIdentification43 `xml:"InitgPty,omitempty"`

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentification43 `xml:"Dbtr"`

	// Unambiguous identification of the account of the debtor to which a debit entry will be made as a result of the transaction.
	DebtorAccount *CashAccount24 `xml:"DbtrAcct,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"DbtrAgt"`

	// Unambiguous identification of the account of the debtor agent at its servicing agent in the payment chain.
	DebtorAgentAccount *CashAccount24 `xml:"DbtrAgtAcct,omitempty"`

	// Agent immediately prior to the instructing agent.
	PreviousInstructingAgent *BranchAndFinancialInstitutionIdentification5 `xml:"PrvsInstgAgt,omitempty"`

	// Unambiguous identification of the account of the previous instructing agent at its servicing agent in the payment chain.
	PreviousInstructingAgentAccount *CashAccount24 `xml:"PrvsInstgAgtAcct,omitempty"`

	// Agent between the debtor's agent and the creditor's agent.
	//
	// Usage: If more than one intermediary agent is present, then IntermediaryAgent1 identifies the agent between the DebtorAgent and the IntermediaryAgent2.
	IntermediaryAgent1 *BranchAndFinancialInstitutionIdentification5 `xml:"IntrmyAgt1,omitempty"`

	// Unambiguous identification of the account of the intermediary agent 1 at its servicing agent in the payment chain.
	IntermediaryAgent1Account *CashAccount24 `xml:"IntrmyAgt1Acct,omitempty"`

	// Agent between the debtor's agent and the creditor's agent.
	//
	// Usage: If more than two intermediary agents are present, then IntermediaryAgent2 identifies the agent between the IntermediaryAgent1 and the IntermediaryAgent3.
	IntermediaryAgent2 *BranchAndFinancialInstitutionIdentification5 `xml:"IntrmyAgt2,omitempty"`

	// Unambiguous identification of the account of the intermediary agent 2 at its servicing agent in the payment chain.
	IntermediaryAgent2Account *CashAccount24 `xml:"IntrmyAgt2Acct,omitempty"`

	// Agent between the debtor's agent and the creditor's agent.
	//
	// Usage: If IntermediaryAgent3 is present, then it identifies the agent between the IntermediaryAgent 2 and the CreditorAgent.
	IntermediaryAgent3 *BranchAndFinancialInstitutionIdentification5 `xml:"IntrmyAgt3,omitempty"`

	// Unambiguous identification of the account of the intermediary agent 3 at its servicing agent in the payment chain.
	IntermediaryAgent3Account *CashAccount24 `xml:"IntrmyAgt3Acct,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"CdtrAgt"`

	// Unambiguous identification of the account of the creditor agent at its servicing agent to which a credit entry will be made as a result of the payment transaction.
	CreditorAgentAccount *CashAccount24 `xml:"CdtrAgtAcct,omitempty"`

	// Party to which an amount of money is due.
	Creditor *PartyIdentification43 `xml:"Cdtr"`

	// Unambiguous identification of the account of the creditor to which a credit entry will be posted as a result of the payment transaction.
	CreditorAccount *CashAccount24 `xml:"CdtrAcct,omitempty"`

	// Ultimate party to which an amount of money is due.
	UltimateCreditor *PartyIdentification43 `xml:"UltmtCdtr,omitempty"`

	// Information supplied to enable the matching of an entry with the items that the transfer is intended to settle, such as commercial invoices in an accounts' receivable system.
	RemittanceInformation *RemittanceInformation7 `xml:"RmtInf,omitempty"`

	// Amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as ordered by the initiating party.
	InstructedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"InstdAmt,omitempty"`
}

func (c *CreditTransferTransaction3) AddUltimateDebtor() *PartyIdentification43 {
	c.UltimateDebtor = new(PartyIdentification43)
	return c.UltimateDebtor
}

func (c *CreditTransferTransaction3) AddInitiatingParty() *PartyIdentification43 {
	c.InitiatingParty = new(PartyIdentification43)
	return c.InitiatingParty
}

func (c *CreditTransferTransaction3) AddDebtor() *PartyIdentification43 {
	c.Debtor = new(PartyIdentification43)
	return c.Debtor
}

func (c *CreditTransferTransaction3) AddDebtorAccount() *CashAccount24 {
	c.DebtorAccount = new(CashAccount24)
	return c.DebtorAccount
}

func (c *CreditTransferTransaction3) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification5 {
	c.DebtorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return c.DebtorAgent
}

func (c *CreditTransferTransaction3) AddDebtorAgentAccount() *CashAccount24 {
	c.DebtorAgentAccount = new(CashAccount24)
	return c.DebtorAgentAccount
}

func (c *CreditTransferTransaction3) AddPreviousInstructingAgent() *BranchAndFinancialInstitutionIdentification5 {
	c.PreviousInstructingAgent = new(BranchAndFinancialInstitutionIdentification5)
	return c.PreviousInstructingAgent
}

func (c *CreditTransferTransaction3) AddPreviousInstructingAgentAccount() *CashAccount24 {
	c.PreviousInstructingAgentAccount = new(CashAccount24)
	return c.PreviousInstructingAgentAccount
}

func (c *CreditTransferTransaction3) AddIntermediaryAgent1() *BranchAndFinancialInstitutionIdentification5 {
	c.IntermediaryAgent1 = new(BranchAndFinancialInstitutionIdentification5)
	return c.IntermediaryAgent1
}

func (c *CreditTransferTransaction3) AddIntermediaryAgent1Account() *CashAccount24 {
	c.IntermediaryAgent1Account = new(CashAccount24)
	return c.IntermediaryAgent1Account
}

func (c *CreditTransferTransaction3) AddIntermediaryAgent2() *BranchAndFinancialInstitutionIdentification5 {
	c.IntermediaryAgent2 = new(BranchAndFinancialInstitutionIdentification5)
	return c.IntermediaryAgent2
}

func (c *CreditTransferTransaction3) AddIntermediaryAgent2Account() *CashAccount24 {
	c.IntermediaryAgent2Account = new(CashAccount24)
	return c.IntermediaryAgent2Account
}

func (c *CreditTransferTransaction3) AddIntermediaryAgent3() *BranchAndFinancialInstitutionIdentification5 {
	c.IntermediaryAgent3 = new(BranchAndFinancialInstitutionIdentification5)
	return c.IntermediaryAgent3
}

func (c *CreditTransferTransaction3) AddIntermediaryAgent3Account() *CashAccount24 {
	c.IntermediaryAgent3Account = new(CashAccount24)
	return c.IntermediaryAgent3Account
}

func (c *CreditTransferTransaction3) AddCreditorAgent() *BranchAndFinancialInstitutionIdentification5 {
	c.CreditorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return c.CreditorAgent
}

func (c *CreditTransferTransaction3) AddCreditorAgentAccount() *CashAccount24 {
	c.CreditorAgentAccount = new(CashAccount24)
	return c.CreditorAgentAccount
}

func (c *CreditTransferTransaction3) AddCreditor() *PartyIdentification43 {
	c.Creditor = new(PartyIdentification43)
	return c.Creditor
}

func (c *CreditTransferTransaction3) AddCreditorAccount() *CashAccount24 {
	c.CreditorAccount = new(CashAccount24)
	return c.CreditorAccount
}

func (c *CreditTransferTransaction3) AddUltimateCreditor() *PartyIdentification43 {
	c.UltimateCreditor = new(PartyIdentification43)
	return c.UltimateCreditor
}

func (c *CreditTransferTransaction3) AddRemittanceInformation() *RemittanceInformation7 {
	c.RemittanceInformation = new(RemittanceInformation7)
	return c.RemittanceInformation
}

func (c *CreditTransferTransaction3) SetInstructedAmount(value, currency string) {
	c.InstructedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}
