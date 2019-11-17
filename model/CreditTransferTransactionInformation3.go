package model

// Set of elements providing information specific to the individual transaction(s) included in the message.
type CreditTransferTransactionInformation3 struct {

	// Set of elements to reference a payment instruction.
	PaymentIdentification *PaymentIdentification2 `xml:"PmtId"`

	// Set of elements used to further specify the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation5 `xml:"PmtTpInf,omitempty"`

	// Amount of money moved between the instructing agent and the instructed agent.
	InterbankSettlementAmount *CurrencyAndAmount `xml:"IntrBkSttlmAmt"`

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	InterbankSettlementDate *ISODate `xml:"IntrBkSttlmDt,omitempty"`

	// Provides information on the occurred settlement time(s) of the payment transaction.
	SettlementTimeIndication *SettlementDateTimeIndication1 `xml:"SttlmTmIndctn,omitempty"`

	// Provides information on the requested settlement time of the payment instruction.
	SettlementTimeRequest *SettlementTimeRequest1 `xml:"SttlmTmReq,omitempty"`

	// Agent immediately prior to the instructing agent.
	PreviousInstructingAgent *BranchAndFinancialInstitutionIdentification3 `xml:"PrvsInstgAgt,omitempty"`

	// Unambiguous identification of the account of the previous instructing agent at its servicing agent in the payment chain.
	PreviousInstructingAgentAccount *CashAccount7 `xml:"PrvsInstgAgtAcct,omitempty"`

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	InstructingAgent *BranchAndFinancialInstitutionIdentification3 `xml:"InstgAgt,omitempty"`

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	InstructedAgent *BranchAndFinancialInstitutionIdentification3 `xml:"InstdAgt,omitempty"`

	// Agent between the debtor agent and creditor agent.
	//
	// Usage: If more than one intermediary agent is present, then IntermediaryAgent1 identifies the agent between the debtor agent and the intermediary agent 2.
	IntermediaryAgent1 *BranchAndFinancialInstitutionIdentification3 `xml:"IntrmyAgt1,omitempty"`

	// Unambiguous identification of the account of the intermediary agent 1 at its servicing agent in the payment chain.
	IntermediaryAgent1Account *CashAccount7 `xml:"IntrmyAgt1Acct,omitempty"`

	// Agent between the debtor agent and creditor agent.
	//
	// Usage: If more than one intermediary agent is present, then IntermediaryAgent2 identifies the agent between the intermediary agent 1 and the intermediary agent 3.
	IntermediaryAgent2 *BranchAndFinancialInstitutionIdentification3 `xml:"IntrmyAgt2,omitempty"`

	// Unambiguous identification of the account of the intermediary agent 2 at its servicing agent in the payment chain.
	IntermediaryAgent2Account *CashAccount7 `xml:"IntrmyAgt2Acct,omitempty"`

	// Agent between the debtor agent and creditor agent.
	//
	// Usage: If more than one intermediary agent is present, then IntermediaryAgent3 identifies the agent between the  intermediary agent 2 and the creditor agent.
	IntermediaryAgent3 *BranchAndFinancialInstitutionIdentification3 `xml:"IntrmyAgt3,omitempty"`

	// Unambiguous identification of the account of the intermediary agent 3 at its servicing agent in the payment chain.
	IntermediaryAgent3Account *CashAccount7 `xml:"IntrmyAgt3Acct,omitempty"`

	// Ultimate financial institution that owes an amount of money to the (ultimate) institutional creditor.
	UltimateDebtor *BranchAndFinancialInstitutionIdentification3 `xml:"UltmtDbtr,omitempty"`

	// Financial institution that owes an amount of money to the (ultimate) financial institutional creditor.
	Debtor *BranchAndFinancialInstitutionIdentification3 `xml:"Dbtr"`

	// Unambiguous identification of the account of the debtor to which a debit entry will be made as a result of the transaction.
	DebtorAccount *CashAccount7 `xml:"DbtrAcct,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification3 `xml:"DbtrAgt,omitempty"`

	// Unambiguous identification of the account of the debtor agent at its servicing agent in the payment chain.
	DebtorAgentAccount *CashAccount7 `xml:"DbtrAgtAcct,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *BranchAndFinancialInstitutionIdentification3 `xml:"CdtrAgt,omitempty"`

	// Unambiguous identification of the account of the creditor agent at its servicing agent to which a credit entry will be made as a result of the payment transaction.
	CreditorAgentAccount *CashAccount7 `xml:"CdtrAgtAcct,omitempty"`

	// Financial institution that receives an amount of money from the financial institutional debtor.
	Creditor *BranchAndFinancialInstitutionIdentification3 `xml:"Cdtr"`

	// Unambiguous identification of the account of the creditor to which a credit entry will be posted as a result of the payment transaction.
	CreditorAccount *CashAccount7 `xml:"CdtrAcct,omitempty"`

	// Ultimate financial institution to which an amount of money is due.
	UltimateCreditor *BranchAndFinancialInstitutionIdentification3 `xml:"UltmtCdtr,omitempty"`

	// Further information related to the processing of the payment instruction that may need to be acted upon by the creditor agent.
	//
	// Usage: The instruction can relate to a level of service, can be an instruction to be executed by the creditor's agent, or can be information required by the creditor's agent to process the instruction.
	InstructionForCreditorAgent []*InstructionForCreditorAgent2 `xml:"InstrForCdtrAgt,omitempty"`

	// Further information related to the processing of the payment instruction that may need to be acted upon by the next agent.
	//
	// Usage: The next agent may not be the creditor agent.
	// The instruction can relate to a level of service, can be an instruction that has to be executed by the agent, or can be information required by the next agent.
	InstructionForNextAgent []*InstructionForNextAgent1 `xml:"InstrForNxtAgt,omitempty"`

	// Information supplied to enable the matching of an entry with the items that the transfer is intended to settle.
	RemittanceInformation *RemittanceInformation2 `xml:"RmtInf,omitempty"`
}

func (c *CreditTransferTransactionInformation3) AddPaymentIdentification() *PaymentIdentification2 {
	c.PaymentIdentification = new(PaymentIdentification2)
	return c.PaymentIdentification
}

func (c *CreditTransferTransactionInformation3) AddPaymentTypeInformation() *PaymentTypeInformation5 {
	c.PaymentTypeInformation = new(PaymentTypeInformation5)
	return c.PaymentTypeInformation
}

func (c *CreditTransferTransactionInformation3) SetInterbankSettlementAmount(value, currency string) {
	c.InterbankSettlementAmount = NewCurrencyAndAmount(value, currency)
}

func (c *CreditTransferTransactionInformation3) SetInterbankSettlementDate(value string) {
	c.InterbankSettlementDate = (*ISODate)(&value)
}

func (c *CreditTransferTransactionInformation3) AddSettlementTimeIndication() *SettlementDateTimeIndication1 {
	c.SettlementTimeIndication = new(SettlementDateTimeIndication1)
	return c.SettlementTimeIndication
}

func (c *CreditTransferTransactionInformation3) AddSettlementTimeRequest() *SettlementTimeRequest1 {
	c.SettlementTimeRequest = new(SettlementTimeRequest1)
	return c.SettlementTimeRequest
}

func (c *CreditTransferTransactionInformation3) AddPreviousInstructingAgent() *BranchAndFinancialInstitutionIdentification3 {
	c.PreviousInstructingAgent = new(BranchAndFinancialInstitutionIdentification3)
	return c.PreviousInstructingAgent
}

func (c *CreditTransferTransactionInformation3) AddPreviousInstructingAgentAccount() *CashAccount7 {
	c.PreviousInstructingAgentAccount = new(CashAccount7)
	return c.PreviousInstructingAgentAccount
}

func (c *CreditTransferTransactionInformation3) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification3 {
	c.InstructingAgent = new(BranchAndFinancialInstitutionIdentification3)
	return c.InstructingAgent
}

func (c *CreditTransferTransactionInformation3) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification3 {
	c.InstructedAgent = new(BranchAndFinancialInstitutionIdentification3)
	return c.InstructedAgent
}

func (c *CreditTransferTransactionInformation3) AddIntermediaryAgent1() *BranchAndFinancialInstitutionIdentification3 {
	c.IntermediaryAgent1 = new(BranchAndFinancialInstitutionIdentification3)
	return c.IntermediaryAgent1
}

func (c *CreditTransferTransactionInformation3) AddIntermediaryAgent1Account() *CashAccount7 {
	c.IntermediaryAgent1Account = new(CashAccount7)
	return c.IntermediaryAgent1Account
}

func (c *CreditTransferTransactionInformation3) AddIntermediaryAgent2() *BranchAndFinancialInstitutionIdentification3 {
	c.IntermediaryAgent2 = new(BranchAndFinancialInstitutionIdentification3)
	return c.IntermediaryAgent2
}

func (c *CreditTransferTransactionInformation3) AddIntermediaryAgent2Account() *CashAccount7 {
	c.IntermediaryAgent2Account = new(CashAccount7)
	return c.IntermediaryAgent2Account
}

func (c *CreditTransferTransactionInformation3) AddIntermediaryAgent3() *BranchAndFinancialInstitutionIdentification3 {
	c.IntermediaryAgent3 = new(BranchAndFinancialInstitutionIdentification3)
	return c.IntermediaryAgent3
}

func (c *CreditTransferTransactionInformation3) AddIntermediaryAgent3Account() *CashAccount7 {
	c.IntermediaryAgent3Account = new(CashAccount7)
	return c.IntermediaryAgent3Account
}

func (c *CreditTransferTransactionInformation3) AddUltimateDebtor() *BranchAndFinancialInstitutionIdentification3 {
	c.UltimateDebtor = new(BranchAndFinancialInstitutionIdentification3)
	return c.UltimateDebtor
}

func (c *CreditTransferTransactionInformation3) AddDebtor() *BranchAndFinancialInstitutionIdentification3 {
	c.Debtor = new(BranchAndFinancialInstitutionIdentification3)
	return c.Debtor
}

func (c *CreditTransferTransactionInformation3) AddDebtorAccount() *CashAccount7 {
	c.DebtorAccount = new(CashAccount7)
	return c.DebtorAccount
}

func (c *CreditTransferTransactionInformation3) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification3 {
	c.DebtorAgent = new(BranchAndFinancialInstitutionIdentification3)
	return c.DebtorAgent
}

func (c *CreditTransferTransactionInformation3) AddDebtorAgentAccount() *CashAccount7 {
	c.DebtorAgentAccount = new(CashAccount7)
	return c.DebtorAgentAccount
}

func (c *CreditTransferTransactionInformation3) AddCreditorAgent() *BranchAndFinancialInstitutionIdentification3 {
	c.CreditorAgent = new(BranchAndFinancialInstitutionIdentification3)
	return c.CreditorAgent
}

func (c *CreditTransferTransactionInformation3) AddCreditorAgentAccount() *CashAccount7 {
	c.CreditorAgentAccount = new(CashAccount7)
	return c.CreditorAgentAccount
}

func (c *CreditTransferTransactionInformation3) AddCreditor() *BranchAndFinancialInstitutionIdentification3 {
	c.Creditor = new(BranchAndFinancialInstitutionIdentification3)
	return c.Creditor
}

func (c *CreditTransferTransactionInformation3) AddCreditorAccount() *CashAccount7 {
	c.CreditorAccount = new(CashAccount7)
	return c.CreditorAccount
}

func (c *CreditTransferTransactionInformation3) AddUltimateCreditor() *BranchAndFinancialInstitutionIdentification3 {
	c.UltimateCreditor = new(BranchAndFinancialInstitutionIdentification3)
	return c.UltimateCreditor
}

func (c *CreditTransferTransactionInformation3) AddInstructionForCreditorAgent() *InstructionForCreditorAgent2 {
	newValue := new(InstructionForCreditorAgent2)
	c.InstructionForCreditorAgent = append(c.InstructionForCreditorAgent, newValue)
	return newValue
}

func (c *CreditTransferTransactionInformation3) AddInstructionForNextAgent() *InstructionForNextAgent1 {
	newValue := new(InstructionForNextAgent1)
	c.InstructionForNextAgent = append(c.InstructionForNextAgent, newValue)
	return newValue
}

func (c *CreditTransferTransactionInformation3) AddRemittanceInformation() *RemittanceInformation2 {
	c.RemittanceInformation = new(RemittanceInformation2)
	return c.RemittanceInformation
}
