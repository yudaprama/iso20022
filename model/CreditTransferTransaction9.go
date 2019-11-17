package model

// Provides further details specific to the individual transaction(s) included in the message.
type CreditTransferTransaction9 struct {

	// Unique identification, as assigned by a sending party, to unambiguously identify the credit instruction within the message.
	CreditIdentification *Max35Text `xml:"CdtId"`

	// Identifies whether a single entry per individual direct debit transaction or a batch entry for the sum of the amounts of all transactions within the group of a message is requested.
	// Usage: Batch booking is used to request and not order a possible batch booking.
	BatchBooking *BatchBookingIndicator `xml:"BtchBookg,omitempty"`

	// Further specifies the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation21 `xml:"PmtTpInf,omitempty"`

	// Amount of money moved between the instructing agent and the instructed agent.
	TotalInterbankSettlementAmount *ActiveCurrencyAndAmount `xml:"TtlIntrBkSttlmAmt,omitempty"`

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	InterbankSettlementDate *ISODate `xml:"IntrBkSttlmDt,omitempty"`

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	InstructingAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstgAgt,omitempty"`

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	InstructedAgent *BranchAndFinancialInstitutionIdentification5 `xml:"InstdAgt,omitempty"`

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
	CreditorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"CdtrAgt,omitempty"`

	// Unambiguous identification of the account of the creditor agent at its servicing agent to which a credit entry will be made as a result of the payment transaction.
	CreditorAgentAccount *CashAccount24 `xml:"CdtrAgtAcct,omitempty"`

	// Financial institution that receives an amount of money from the financial institutional debtor.
	Creditor *BranchAndFinancialInstitutionIdentification5 `xml:"Cdtr"`

	// Unambiguous identification of the account of the creditor to which a credit entry will be posted as a result of the payment transaction.
	CreditorAccount *CashAccount24 `xml:"CdtrAcct,omitempty"`

	// Ultimate financial institution to which an amount of money is due.
	UltimateCreditor *BranchAndFinancialInstitutionIdentification5 `xml:"UltmtCdtr,omitempty"`

	// Further information related to the processing of the payment instruction, provided by the initiating party, and intended for the creditor agent.
	InstructionForCreditorAgent []*InstructionForCreditorAgent2 `xml:"InstrForCdtrAgt,omitempty"`

	// Provides information on the individual debit transaction(s) included in the message.
	DirectDebitTransactionInformation []*DirectDebitTransactionInformation15 `xml:"DrctDbtTxInf"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *CreditTransferTransaction9) SetCreditIdentification(value string) {
	c.CreditIdentification = (*Max35Text)(&value)
}

func (c *CreditTransferTransaction9) SetBatchBooking(value string) {
	c.BatchBooking = (*BatchBookingIndicator)(&value)
}

func (c *CreditTransferTransaction9) AddPaymentTypeInformation() *PaymentTypeInformation21 {
	c.PaymentTypeInformation = new(PaymentTypeInformation21)
	return c.PaymentTypeInformation
}

func (c *CreditTransferTransaction9) SetTotalInterbankSettlementAmount(value, currency string) {
	c.TotalInterbankSettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CreditTransferTransaction9) SetInterbankSettlementDate(value string) {
	c.InterbankSettlementDate = (*ISODate)(&value)
}

func (c *CreditTransferTransaction9) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification5 {
	c.InstructingAgent = new(BranchAndFinancialInstitutionIdentification5)
	return c.InstructingAgent
}

func (c *CreditTransferTransaction9) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification5 {
	c.InstructedAgent = new(BranchAndFinancialInstitutionIdentification5)
	return c.InstructedAgent
}

func (c *CreditTransferTransaction9) AddIntermediaryAgent1() *BranchAndFinancialInstitutionIdentification5 {
	c.IntermediaryAgent1 = new(BranchAndFinancialInstitutionIdentification5)
	return c.IntermediaryAgent1
}

func (c *CreditTransferTransaction9) AddIntermediaryAgent1Account() *CashAccount24 {
	c.IntermediaryAgent1Account = new(CashAccount24)
	return c.IntermediaryAgent1Account
}

func (c *CreditTransferTransaction9) AddIntermediaryAgent2() *BranchAndFinancialInstitutionIdentification5 {
	c.IntermediaryAgent2 = new(BranchAndFinancialInstitutionIdentification5)
	return c.IntermediaryAgent2
}

func (c *CreditTransferTransaction9) AddIntermediaryAgent2Account() *CashAccount24 {
	c.IntermediaryAgent2Account = new(CashAccount24)
	return c.IntermediaryAgent2Account
}

func (c *CreditTransferTransaction9) AddIntermediaryAgent3() *BranchAndFinancialInstitutionIdentification5 {
	c.IntermediaryAgent3 = new(BranchAndFinancialInstitutionIdentification5)
	return c.IntermediaryAgent3
}

func (c *CreditTransferTransaction9) AddIntermediaryAgent3Account() *CashAccount24 {
	c.IntermediaryAgent3Account = new(CashAccount24)
	return c.IntermediaryAgent3Account
}

func (c *CreditTransferTransaction9) AddCreditorAgent() *BranchAndFinancialInstitutionIdentification5 {
	c.CreditorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return c.CreditorAgent
}

func (c *CreditTransferTransaction9) AddCreditorAgentAccount() *CashAccount24 {
	c.CreditorAgentAccount = new(CashAccount24)
	return c.CreditorAgentAccount
}

func (c *CreditTransferTransaction9) AddCreditor() *BranchAndFinancialInstitutionIdentification5 {
	c.Creditor = new(BranchAndFinancialInstitutionIdentification5)
	return c.Creditor
}

func (c *CreditTransferTransaction9) AddCreditorAccount() *CashAccount24 {
	c.CreditorAccount = new(CashAccount24)
	return c.CreditorAccount
}

func (c *CreditTransferTransaction9) AddUltimateCreditor() *BranchAndFinancialInstitutionIdentification5 {
	c.UltimateCreditor = new(BranchAndFinancialInstitutionIdentification5)
	return c.UltimateCreditor
}

func (c *CreditTransferTransaction9) AddInstructionForCreditorAgent() *InstructionForCreditorAgent2 {
	newValue := new(InstructionForCreditorAgent2)
	c.InstructionForCreditorAgent = append(c.InstructionForCreditorAgent, newValue)
	return newValue
}

func (c *CreditTransferTransaction9) AddDirectDebitTransactionInformation() *DirectDebitTransactionInformation15 {
	newValue := new(DirectDebitTransactionInformation15)
	c.DirectDebitTransactionInformation = append(c.DirectDebitTransactionInformation, newValue)
	return newValue
}

func (c *CreditTransferTransaction9) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
