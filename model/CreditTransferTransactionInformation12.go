package model

// Set of elements used to provide information specific to the individual transaction(s) included in the message.
type CreditTransferTransactionInformation12 struct {

	// Ultimate party that owes an amount of money to the (ultimate) creditor.
	UltimateDebtor *PartyIdentification32 `xml:"UltmtDbtr,omitempty"`

	// Party that initiates the payment.
	// Usage: This can be either the debtor or a party that initiates the credit transfer on behalf of the debtor.
	InitiatingParty *PartyIdentification32 `xml:"InitgPty,omitempty"`

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentification32 `xml:"Dbtr"`

	// Unambiguous identification of the account of the debtor to which a debit entry will be made as a result of the transaction.
	DebtorAccount *CashAccount16 `xml:"DbtrAcct,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification4 `xml:"DbtrAgt"`

	// Unambiguous identification of the account of the debtor agent at its servicing agent in the payment chain.
	DebtorAgentAccount *CashAccount16 `xml:"DbtrAgtAcct,omitempty"`

	// Agent immediately prior to the instructing agent.
	PreviousInstructingAgent *BranchAndFinancialInstitutionIdentification4 `xml:"PrvsInstgAgt,omitempty"`

	// Unambiguous identification of the account of the previous instructing agent at its servicing agent in the payment chain.
	PreviousInstructingAgentAccount *CashAccount16 `xml:"PrvsInstgAgtAcct,omitempty"`

	// Agent between the debtor's agent and the creditor's agent.
	//
	// Usage: If more than one intermediary agent is present, then IntermediaryAgent1 identifies the agent between the DebtorAgent and the IntermediaryAgent2.
	IntermediaryAgent1 *BranchAndFinancialInstitutionIdentification4 `xml:"IntrmyAgt1,omitempty"`

	// Unambiguous identification of the account of the intermediary agent 1 at its servicing agent in the payment chain.
	IntermediaryAgent1Account *CashAccount16 `xml:"IntrmyAgt1Acct,omitempty"`

	// Agent between the debtor's agent and the creditor's agent.
	//
	// Usage: If more than two intermediary agents are present, then IntermediaryAgent2 identifies the agent between the IntermediaryAgent1 and the IntermediaryAgent3.
	IntermediaryAgent2 *BranchAndFinancialInstitutionIdentification4 `xml:"IntrmyAgt2,omitempty"`

	// Unambiguous identification of the account of the intermediary agent 2 at its servicing agent in the payment chain.
	IntermediaryAgent2Account *CashAccount16 `xml:"IntrmyAgt2Acct,omitempty"`

	// Agent between the debtor's agent and the creditor's agent.
	//
	// Usage: If IntermediaryAgent3 is present, then it identifies the agent between the IntermediaryAgent 2 and the CreditorAgent.
	IntermediaryAgent3 *BranchAndFinancialInstitutionIdentification4 `xml:"IntrmyAgt3,omitempty"`

	// Unambiguous identification of the account of the intermediary agent 3 at its servicing agent in the payment chain.
	IntermediaryAgent3Account *CashAccount16 `xml:"IntrmyAgt3Acct,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *BranchAndFinancialInstitutionIdentification4 `xml:"CdtrAgt"`

	// Unambiguous identification of the account of the creditor agent at its servicing agent to which a credit entry will be made as a result of the payment transaction.
	CreditorAgentAccount *CashAccount16 `xml:"CdtrAgtAcct,omitempty"`

	// Party to which an amount of money is due.
	Creditor *PartyIdentification32 `xml:"Cdtr"`

	// Unambiguous identification of the account of the creditor to which a credit entry will be posted as a result of the payment transaction.
	CreditorAccount *CashAccount16 `xml:"CdtrAcct,omitempty"`

	// Ultimate party to which an amount of money is due.
	UltimateCreditor *PartyIdentification32 `xml:"UltmtCdtr,omitempty"`

	// Information supplied to enable the matching of an entry with the items that the transfer is intended to settle, such as commercial invoices in an accounts' receivable system.
	RemittanceInformation *RemittanceInformation5 `xml:"RmtInf,omitempty"`

	// Amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as ordered by the initiating party.
	InstructedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"InstdAmt,omitempty"`
}

func (c *CreditTransferTransactionInformation12) AddUltimateDebtor() *PartyIdentification32 {
	c.UltimateDebtor = new(PartyIdentification32)
	return c.UltimateDebtor
}

func (c *CreditTransferTransactionInformation12) AddInitiatingParty() *PartyIdentification32 {
	c.InitiatingParty = new(PartyIdentification32)
	return c.InitiatingParty
}

func (c *CreditTransferTransactionInformation12) AddDebtor() *PartyIdentification32 {
	c.Debtor = new(PartyIdentification32)
	return c.Debtor
}

func (c *CreditTransferTransactionInformation12) AddDebtorAccount() *CashAccount16 {
	c.DebtorAccount = new(CashAccount16)
	return c.DebtorAccount
}

func (c *CreditTransferTransactionInformation12) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification4 {
	c.DebtorAgent = new(BranchAndFinancialInstitutionIdentification4)
	return c.DebtorAgent
}

func (c *CreditTransferTransactionInformation12) AddDebtorAgentAccount() *CashAccount16 {
	c.DebtorAgentAccount = new(CashAccount16)
	return c.DebtorAgentAccount
}

func (c *CreditTransferTransactionInformation12) AddPreviousInstructingAgent() *BranchAndFinancialInstitutionIdentification4 {
	c.PreviousInstructingAgent = new(BranchAndFinancialInstitutionIdentification4)
	return c.PreviousInstructingAgent
}

func (c *CreditTransferTransactionInformation12) AddPreviousInstructingAgentAccount() *CashAccount16 {
	c.PreviousInstructingAgentAccount = new(CashAccount16)
	return c.PreviousInstructingAgentAccount
}

func (c *CreditTransferTransactionInformation12) AddIntermediaryAgent1() *BranchAndFinancialInstitutionIdentification4 {
	c.IntermediaryAgent1 = new(BranchAndFinancialInstitutionIdentification4)
	return c.IntermediaryAgent1
}

func (c *CreditTransferTransactionInformation12) AddIntermediaryAgent1Account() *CashAccount16 {
	c.IntermediaryAgent1Account = new(CashAccount16)
	return c.IntermediaryAgent1Account
}

func (c *CreditTransferTransactionInformation12) AddIntermediaryAgent2() *BranchAndFinancialInstitutionIdentification4 {
	c.IntermediaryAgent2 = new(BranchAndFinancialInstitutionIdentification4)
	return c.IntermediaryAgent2
}

func (c *CreditTransferTransactionInformation12) AddIntermediaryAgent2Account() *CashAccount16 {
	c.IntermediaryAgent2Account = new(CashAccount16)
	return c.IntermediaryAgent2Account
}

func (c *CreditTransferTransactionInformation12) AddIntermediaryAgent3() *BranchAndFinancialInstitutionIdentification4 {
	c.IntermediaryAgent3 = new(BranchAndFinancialInstitutionIdentification4)
	return c.IntermediaryAgent3
}

func (c *CreditTransferTransactionInformation12) AddIntermediaryAgent3Account() *CashAccount16 {
	c.IntermediaryAgent3Account = new(CashAccount16)
	return c.IntermediaryAgent3Account
}

func (c *CreditTransferTransactionInformation12) AddCreditorAgent() *BranchAndFinancialInstitutionIdentification4 {
	c.CreditorAgent = new(BranchAndFinancialInstitutionIdentification4)
	return c.CreditorAgent
}

func (c *CreditTransferTransactionInformation12) AddCreditorAgentAccount() *CashAccount16 {
	c.CreditorAgentAccount = new(CashAccount16)
	return c.CreditorAgentAccount
}

func (c *CreditTransferTransactionInformation12) AddCreditor() *PartyIdentification32 {
	c.Creditor = new(PartyIdentification32)
	return c.Creditor
}

func (c *CreditTransferTransactionInformation12) AddCreditorAccount() *CashAccount16 {
	c.CreditorAccount = new(CashAccount16)
	return c.CreditorAccount
}

func (c *CreditTransferTransactionInformation12) AddUltimateCreditor() *PartyIdentification32 {
	c.UltimateCreditor = new(PartyIdentification32)
	return c.UltimateCreditor
}

func (c *CreditTransferTransactionInformation12) AddRemittanceInformation() *RemittanceInformation5 {
	c.RemittanceInformation = new(RemittanceInformation5)
	return c.RemittanceInformation
}

func (c *CreditTransferTransactionInformation12) SetInstructedAmount(value, currency string) {
	c.InstructedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}
