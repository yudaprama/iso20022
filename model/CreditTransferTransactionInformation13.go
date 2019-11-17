package model

// Set of elements used to provide information specific to the individual transaction(s) included in the message.
type CreditTransferTransactionInformation13 struct {

	// Set of elements used to reference a payment instruction.
	PaymentIdentification *PaymentIdentification3 `xml:"PmtId"`

	// Set of elements used to further specify the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation23 `xml:"PmtTpInf,omitempty"`

	// Amount of money moved between the instructing agent and the instructed agent.
	InterbankSettlementAmount *ActiveCurrencyAndAmount `xml:"IntrBkSttlmAmt"`

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	InterbankSettlementDate *ISODate `xml:"IntrBkSttlmDt,omitempty"`

	// Indicator of the urgency or order of importance that the instructing party would like the instructed party to apply to the processing of the settlement instruction.
	SettlementPriority *Priority3Code `xml:"SttlmPrty,omitempty"`

	// Provides information on the occurred settlement time(s) of the payment transaction.
	SettlementTimeIndication *SettlementDateTimeIndication1 `xml:"SttlmTmIndctn,omitempty"`

	// Provides information on the requested settlement time(s) of the payment instruction.
	SettlementTimeRequest *SettlementTimeRequest2 `xml:"SttlmTmReq,omitempty"`

	// Agent immediately prior to the instructing agent.
	PreviousInstructingAgent *BranchAndFinancialInstitutionIdentification4 `xml:"PrvsInstgAgt,omitempty"`

	// Unambiguous identification of the account of the previous instructing agent at its servicing agent in the payment chain.
	PreviousInstructingAgentAccount *CashAccount16 `xml:"PrvsInstgAgtAcct,omitempty"`

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	InstructingAgent *BranchAndFinancialInstitutionIdentification4 `xml:"InstgAgt,omitempty"`

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	InstructedAgent *BranchAndFinancialInstitutionIdentification4 `xml:"InstdAgt,omitempty"`

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

	// Ultimate financial institution that owes an amount of money to the (ultimate) institutional creditor.
	UltimateDebtor *BranchAndFinancialInstitutionIdentification4 `xml:"UltmtDbtr,omitempty"`

	// Financial institution that owes an amount of money to the (ultimate) financial institutional creditor.
	Debtor *BranchAndFinancialInstitutionIdentification4 `xml:"Dbtr"`

	// Unambiguous identification of the account of the debtor to which a debit entry will be made as a result of the transaction.
	DebtorAccount *CashAccount16 `xml:"DbtrAcct,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification4 `xml:"DbtrAgt,omitempty"`

	// Unambiguous identification of the account of the debtor agent at its servicing agent in the payment chain.
	DebtorAgentAccount *CashAccount16 `xml:"DbtrAgtAcct,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *BranchAndFinancialInstitutionIdentification4 `xml:"CdtrAgt,omitempty"`

	// Unambiguous identification of the account of the creditor agent at its servicing agent to which a credit entry will be made as a result of the payment transaction.
	CreditorAgentAccount *CashAccount16 `xml:"CdtrAgtAcct,omitempty"`

	// Financial institution that receives an amount of money from the financial institutional debtor.
	Creditor *BranchAndFinancialInstitutionIdentification4 `xml:"Cdtr"`

	// Unambiguous identification of the account of the creditor to which a credit entry will be posted as a result of the payment transaction.
	CreditorAccount *CashAccount16 `xml:"CdtrAcct,omitempty"`

	// Ultimate financial institution to which an amount of money is due.
	UltimateCreditor *BranchAndFinancialInstitutionIdentification4 `xml:"UltmtCdtr,omitempty"`

	// Further information related to the processing of the payment instruction, provided by the initiating party, and intended for the creditor agent.
	InstructionForCreditorAgent []*InstructionForCreditorAgent2 `xml:"InstrForCdtrAgt,omitempty"`

	// Further information related to the processing of the payment instruction that may need to be acted upon by the next agent.
	//
	// Usage: The next agent may not be the creditor agent.
	// The instruction can relate to a level of service, can be an instruction that has to be executed by the agent, or can be information required by the next agent.
	InstructionForNextAgent []*InstructionForNextAgent1 `xml:"InstrForNxtAgt,omitempty"`

	// Information supplied to enable the matching of an entry with the items that the transfer is intended to settle, such as commercial invoices in an accounts' receivable system.
	RemittanceInformation *RemittanceInformation2 `xml:"RmtInf,omitempty"`

	// Set of elements used to provide information on the underlying customer credit transfer for which cover is provided.
	UnderlyingCustomerCreditTransfer *CreditTransferTransactionInformation12 `xml:"UndrlygCstmrCdtTrf,omitempty"`
}

func (c *CreditTransferTransactionInformation13) AddPaymentIdentification() *PaymentIdentification3 {
	c.PaymentIdentification = new(PaymentIdentification3)
	return c.PaymentIdentification
}

func (c *CreditTransferTransactionInformation13) AddPaymentTypeInformation() *PaymentTypeInformation23 {
	c.PaymentTypeInformation = new(PaymentTypeInformation23)
	return c.PaymentTypeInformation
}

func (c *CreditTransferTransactionInformation13) SetInterbankSettlementAmount(value, currency string) {
	c.InterbankSettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CreditTransferTransactionInformation13) SetInterbankSettlementDate(value string) {
	c.InterbankSettlementDate = (*ISODate)(&value)
}

func (c *CreditTransferTransactionInformation13) SetSettlementPriority(value string) {
	c.SettlementPriority = (*Priority3Code)(&value)
}

func (c *CreditTransferTransactionInformation13) AddSettlementTimeIndication() *SettlementDateTimeIndication1 {
	c.SettlementTimeIndication = new(SettlementDateTimeIndication1)
	return c.SettlementTimeIndication
}

func (c *CreditTransferTransactionInformation13) AddSettlementTimeRequest() *SettlementTimeRequest2 {
	c.SettlementTimeRequest = new(SettlementTimeRequest2)
	return c.SettlementTimeRequest
}

func (c *CreditTransferTransactionInformation13) AddPreviousInstructingAgent() *BranchAndFinancialInstitutionIdentification4 {
	c.PreviousInstructingAgent = new(BranchAndFinancialInstitutionIdentification4)
	return c.PreviousInstructingAgent
}

func (c *CreditTransferTransactionInformation13) AddPreviousInstructingAgentAccount() *CashAccount16 {
	c.PreviousInstructingAgentAccount = new(CashAccount16)
	return c.PreviousInstructingAgentAccount
}

func (c *CreditTransferTransactionInformation13) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification4 {
	c.InstructingAgent = new(BranchAndFinancialInstitutionIdentification4)
	return c.InstructingAgent
}

func (c *CreditTransferTransactionInformation13) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification4 {
	c.InstructedAgent = new(BranchAndFinancialInstitutionIdentification4)
	return c.InstructedAgent
}

func (c *CreditTransferTransactionInformation13) AddIntermediaryAgent1() *BranchAndFinancialInstitutionIdentification4 {
	c.IntermediaryAgent1 = new(BranchAndFinancialInstitutionIdentification4)
	return c.IntermediaryAgent1
}

func (c *CreditTransferTransactionInformation13) AddIntermediaryAgent1Account() *CashAccount16 {
	c.IntermediaryAgent1Account = new(CashAccount16)
	return c.IntermediaryAgent1Account
}

func (c *CreditTransferTransactionInformation13) AddIntermediaryAgent2() *BranchAndFinancialInstitutionIdentification4 {
	c.IntermediaryAgent2 = new(BranchAndFinancialInstitutionIdentification4)
	return c.IntermediaryAgent2
}

func (c *CreditTransferTransactionInformation13) AddIntermediaryAgent2Account() *CashAccount16 {
	c.IntermediaryAgent2Account = new(CashAccount16)
	return c.IntermediaryAgent2Account
}

func (c *CreditTransferTransactionInformation13) AddIntermediaryAgent3() *BranchAndFinancialInstitutionIdentification4 {
	c.IntermediaryAgent3 = new(BranchAndFinancialInstitutionIdentification4)
	return c.IntermediaryAgent3
}

func (c *CreditTransferTransactionInformation13) AddIntermediaryAgent3Account() *CashAccount16 {
	c.IntermediaryAgent3Account = new(CashAccount16)
	return c.IntermediaryAgent3Account
}

func (c *CreditTransferTransactionInformation13) AddUltimateDebtor() *BranchAndFinancialInstitutionIdentification4 {
	c.UltimateDebtor = new(BranchAndFinancialInstitutionIdentification4)
	return c.UltimateDebtor
}

func (c *CreditTransferTransactionInformation13) AddDebtor() *BranchAndFinancialInstitutionIdentification4 {
	c.Debtor = new(BranchAndFinancialInstitutionIdentification4)
	return c.Debtor
}

func (c *CreditTransferTransactionInformation13) AddDebtorAccount() *CashAccount16 {
	c.DebtorAccount = new(CashAccount16)
	return c.DebtorAccount
}

func (c *CreditTransferTransactionInformation13) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification4 {
	c.DebtorAgent = new(BranchAndFinancialInstitutionIdentification4)
	return c.DebtorAgent
}

func (c *CreditTransferTransactionInformation13) AddDebtorAgentAccount() *CashAccount16 {
	c.DebtorAgentAccount = new(CashAccount16)
	return c.DebtorAgentAccount
}

func (c *CreditTransferTransactionInformation13) AddCreditorAgent() *BranchAndFinancialInstitutionIdentification4 {
	c.CreditorAgent = new(BranchAndFinancialInstitutionIdentification4)
	return c.CreditorAgent
}

func (c *CreditTransferTransactionInformation13) AddCreditorAgentAccount() *CashAccount16 {
	c.CreditorAgentAccount = new(CashAccount16)
	return c.CreditorAgentAccount
}

func (c *CreditTransferTransactionInformation13) AddCreditor() *BranchAndFinancialInstitutionIdentification4 {
	c.Creditor = new(BranchAndFinancialInstitutionIdentification4)
	return c.Creditor
}

func (c *CreditTransferTransactionInformation13) AddCreditorAccount() *CashAccount16 {
	c.CreditorAccount = new(CashAccount16)
	return c.CreditorAccount
}

func (c *CreditTransferTransactionInformation13) AddUltimateCreditor() *BranchAndFinancialInstitutionIdentification4 {
	c.UltimateCreditor = new(BranchAndFinancialInstitutionIdentification4)
	return c.UltimateCreditor
}

func (c *CreditTransferTransactionInformation13) AddInstructionForCreditorAgent() *InstructionForCreditorAgent2 {
	newValue := new(InstructionForCreditorAgent2)
	c.InstructionForCreditorAgent = append(c.InstructionForCreditorAgent, newValue)
	return newValue
}

func (c *CreditTransferTransactionInformation13) AddInstructionForNextAgent() *InstructionForNextAgent1 {
	newValue := new(InstructionForNextAgent1)
	c.InstructionForNextAgent = append(c.InstructionForNextAgent, newValue)
	return newValue
}

func (c *CreditTransferTransactionInformation13) AddRemittanceInformation() *RemittanceInformation2 {
	c.RemittanceInformation = new(RemittanceInformation2)
	return c.RemittanceInformation
}

func (c *CreditTransferTransactionInformation13) AddUnderlyingCustomerCreditTransfer() *CreditTransferTransactionInformation12 {
	c.UnderlyingCustomerCreditTransfer = new(CreditTransferTransactionInformation12)
	return c.UnderlyingCustomerCreditTransfer
}
