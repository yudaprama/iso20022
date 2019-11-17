package model

// Set of elements providing information specific to the individual transaction(s) included in the message.
type CreditTransferTransactionInformation2 struct {

	// Set of elements to reference a payment instruction.
	PaymentIdentification *PaymentIdentification2 `xml:"PmtId"`

	// Set of elements used to further specify the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation3 `xml:"PmtTpInf,omitempty"`

	// Amount of money moved between the instructing agent and the instructed agent.
	InterbankSettlementAmount *CurrencyAndAmount `xml:"IntrBkSttlmAmt"`

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	InterbankSettlementDate *ISODate `xml:"IntrBkSttlmDt,omitempty"`

	// Provides information on the occurred settlement time(s) of the payment transaction.
	SettlementTimeIndication *SettlementDateTimeIndication1 `xml:"SttlmTmIndctn,omitempty"`

	// Provides information on the requested settlement time of the payment instruction.
	SettlementTimeRequest *SettlementTimeRequest1 `xml:"SttlmTmReq,omitempty"`

	// Point in time when the payment order from the initiating party meets the processing conditions of the account servicing agent (debtor's agent in case of a credit transfer, creditor's agent in case of a direct debit). This means - amongst others - that the account servicing agent has received the payment order and has applied checks as eg, authorisation, availability of funds.
	AcceptanceDateTime *ISODateTime `xml:"AccptncDtTm,omitempty"`

	// Date used for the correction of the value date of a cash pool movement that has been posted with a different value date.
	PoolingAdjustmentDate *ISODate `xml:"PoolgAdjstmntDt,omitempty"`

	// Amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as ordered by the initiating party.
	InstructedAmount *CurrencyAndAmount `xml:"InstdAmt,omitempty"`

	// The factor used for conversion of an amount from one currency to another. This reflects the price at which one currency was bought with another currency.
	ExchangeRate *BaseOneRate `xml:"XchgRate,omitempty"`

	// Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr"`

	// Provides information on the charges related to the payment transaction.
	ChargesInformation []*ChargesInformation1 `xml:"ChrgsInf,omitempty"`

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
	// Usage: If more than two intermediary agents are present, then IntermediaryAgent2 identifies the agent between the intermediary agent 1 and the intermediary agent 3.
	IntermediaryAgent2 *BranchAndFinancialInstitutionIdentification3 `xml:"IntrmyAgt2,omitempty"`

	// Unambiguous identification of the account of the intermediary agent 2 at its servicing agent in the payment chain.
	IntermediaryAgent2Account *CashAccount7 `xml:"IntrmyAgt2Acct,omitempty"`

	// Agent between the debtor agent and creditor agent.
	//
	// Usage: If IntermediaryAgent3 is present, then it identifies the agent between the intermediary agent 2 and the creditor agent.
	IntermediaryAgent3 *BranchAndFinancialInstitutionIdentification3 `xml:"IntrmyAgt3,omitempty"`

	// Unambiguous identification of the account of the intermediary agent 3 at its servicing agent in the payment chain.
	IntermediaryAgent3Account *CashAccount7 `xml:"IntrmyAgt3Acct,omitempty"`

	// Ultimate party that owes an amount of money to the (ultimate) creditor.
	UltimateDebtor *PartyIdentification8 `xml:"UltmtDbtr,omitempty"`

	// Party that initiates the payment. In the payment context, this can either be the debtor (in a credit transfer), the creditor (in a direct debit), or a party that initiates the payment on behalf of the debtor or creditor.
	InitiatingParty *PartyIdentification8 `xml:"InitgPty,omitempty"`

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentification8 `xml:"Dbtr"`

	// Unambiguous identification of the account of the debtor to which a debit entry will be made as a result of the transaction.
	DebtorAccount *CashAccount7 `xml:"DbtrAcct,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification3 `xml:"DbtrAgt"`

	// Unambiguous identification of the account of the debtor agent at its servicing agent in the payment chain.
	DebtorAgentAccount *CashAccount7 `xml:"DbtrAgtAcct,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *BranchAndFinancialInstitutionIdentification3 `xml:"CdtrAgt"`

	// Unambiguous identification of the account of the creditor agent at its servicing agent to which a credit entry will be made as a result of the payment transaction.
	CreditorAgentAccount *CashAccount7 `xml:"CdtrAgtAcct,omitempty"`

	// Party to which an amount of money is due.
	Creditor *PartyIdentification8 `xml:"Cdtr"`

	// Unambiguous identification of the account of the creditor to which a credit entry will be posted as a result of the payment transaction.
	CreditorAccount *CashAccount7 `xml:"CdtrAcct,omitempty"`

	// Ultimate party to which an amount of money is due.
	UltimateCreditor *PartyIdentification8 `xml:"UltmtCdtr,omitempty"`

	// Further information related to the processing of the payment instruction that may need to be acted upon by the creditor agent.
	//
	// Usage: The instruction can relate to a level of service, can be an instruction to be executed by the creditor's agent, or can be information required by the creditor's agent to process the instruction.
	InstructionForCreditorAgent []*InstructionForCreditorAgent1 `xml:"InstrForCdtrAgt,omitempty"`

	// Further information related to the processing of the payment instruction that may need to be acted upon by the next agent.
	//
	// Usage: The next agent may not be the creditor agent.
	// The instruction can relate to a level of service, can be an instruction that has to be executed by the agent, or can be information required by the next agent.
	InstructionForNextAgent []*InstructionForNextAgent1 `xml:"InstrForNxtAgt,omitempty"`

	// Underlying reason for the payment transaction.
	//
	// Usage: Purpose is used by the end-customers, i.e. initiating party, (ultimate) debtor, (ultimate) creditor to provide information concerning the nature of the payment. Purpose is a content element, which is not used for processing by any of the agents involved in the payment chain.
	Purpose *Purpose1Choice `xml:"Purp,omitempty"`

	// Information needed due to regulatory and statutory requirements.
	RegulatoryReporting []*RegulatoryReporting2 `xml:"RgltryRptg,omitempty"`

	// Information related to the handling of the remittance information by any of the agents in the transaction processing chain.
	RelatedRemittanceInformation []*RemittanceLocation1 `xml:"RltdRmtInf,omitempty"`

	// Information supplied to enable the matching of an entry with the items that the transfer is intended to settle, such as commercial invoices in an accounts' receivable system.
	RemittanceInformation *RemittanceInformation1 `xml:"RmtInf,omitempty"`
}

func (c *CreditTransferTransactionInformation2) AddPaymentIdentification() *PaymentIdentification2 {
	c.PaymentIdentification = new(PaymentIdentification2)
	return c.PaymentIdentification
}

func (c *CreditTransferTransactionInformation2) AddPaymentTypeInformation() *PaymentTypeInformation3 {
	c.PaymentTypeInformation = new(PaymentTypeInformation3)
	return c.PaymentTypeInformation
}

func (c *CreditTransferTransactionInformation2) SetInterbankSettlementAmount(value, currency string) {
	c.InterbankSettlementAmount = NewCurrencyAndAmount(value, currency)
}

func (c *CreditTransferTransactionInformation2) SetInterbankSettlementDate(value string) {
	c.InterbankSettlementDate = (*ISODate)(&value)
}

func (c *CreditTransferTransactionInformation2) AddSettlementTimeIndication() *SettlementDateTimeIndication1 {
	c.SettlementTimeIndication = new(SettlementDateTimeIndication1)
	return c.SettlementTimeIndication
}

func (c *CreditTransferTransactionInformation2) AddSettlementTimeRequest() *SettlementTimeRequest1 {
	c.SettlementTimeRequest = new(SettlementTimeRequest1)
	return c.SettlementTimeRequest
}

func (c *CreditTransferTransactionInformation2) SetAcceptanceDateTime(value string) {
	c.AcceptanceDateTime = (*ISODateTime)(&value)
}

func (c *CreditTransferTransactionInformation2) SetPoolingAdjustmentDate(value string) {
	c.PoolingAdjustmentDate = (*ISODate)(&value)
}

func (c *CreditTransferTransactionInformation2) SetInstructedAmount(value, currency string) {
	c.InstructedAmount = NewCurrencyAndAmount(value, currency)
}

func (c *CreditTransferTransactionInformation2) SetExchangeRate(value string) {
	c.ExchangeRate = (*BaseOneRate)(&value)
}

func (c *CreditTransferTransactionInformation2) SetChargeBearer(value string) {
	c.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (c *CreditTransferTransactionInformation2) AddChargesInformation() *ChargesInformation1 {
	newValue := new(ChargesInformation1)
	c.ChargesInformation = append(c.ChargesInformation, newValue)
	return newValue
}

func (c *CreditTransferTransactionInformation2) AddPreviousInstructingAgent() *BranchAndFinancialInstitutionIdentification3 {
	c.PreviousInstructingAgent = new(BranchAndFinancialInstitutionIdentification3)
	return c.PreviousInstructingAgent
}

func (c *CreditTransferTransactionInformation2) AddPreviousInstructingAgentAccount() *CashAccount7 {
	c.PreviousInstructingAgentAccount = new(CashAccount7)
	return c.PreviousInstructingAgentAccount
}

func (c *CreditTransferTransactionInformation2) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification3 {
	c.InstructingAgent = new(BranchAndFinancialInstitutionIdentification3)
	return c.InstructingAgent
}

func (c *CreditTransferTransactionInformation2) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification3 {
	c.InstructedAgent = new(BranchAndFinancialInstitutionIdentification3)
	return c.InstructedAgent
}

func (c *CreditTransferTransactionInformation2) AddIntermediaryAgent1() *BranchAndFinancialInstitutionIdentification3 {
	c.IntermediaryAgent1 = new(BranchAndFinancialInstitutionIdentification3)
	return c.IntermediaryAgent1
}

func (c *CreditTransferTransactionInformation2) AddIntermediaryAgent1Account() *CashAccount7 {
	c.IntermediaryAgent1Account = new(CashAccount7)
	return c.IntermediaryAgent1Account
}

func (c *CreditTransferTransactionInformation2) AddIntermediaryAgent2() *BranchAndFinancialInstitutionIdentification3 {
	c.IntermediaryAgent2 = new(BranchAndFinancialInstitutionIdentification3)
	return c.IntermediaryAgent2
}

func (c *CreditTransferTransactionInformation2) AddIntermediaryAgent2Account() *CashAccount7 {
	c.IntermediaryAgent2Account = new(CashAccount7)
	return c.IntermediaryAgent2Account
}

func (c *CreditTransferTransactionInformation2) AddIntermediaryAgent3() *BranchAndFinancialInstitutionIdentification3 {
	c.IntermediaryAgent3 = new(BranchAndFinancialInstitutionIdentification3)
	return c.IntermediaryAgent3
}

func (c *CreditTransferTransactionInformation2) AddIntermediaryAgent3Account() *CashAccount7 {
	c.IntermediaryAgent3Account = new(CashAccount7)
	return c.IntermediaryAgent3Account
}

func (c *CreditTransferTransactionInformation2) AddUltimateDebtor() *PartyIdentification8 {
	c.UltimateDebtor = new(PartyIdentification8)
	return c.UltimateDebtor
}

func (c *CreditTransferTransactionInformation2) AddInitiatingParty() *PartyIdentification8 {
	c.InitiatingParty = new(PartyIdentification8)
	return c.InitiatingParty
}

func (c *CreditTransferTransactionInformation2) AddDebtor() *PartyIdentification8 {
	c.Debtor = new(PartyIdentification8)
	return c.Debtor
}

func (c *CreditTransferTransactionInformation2) AddDebtorAccount() *CashAccount7 {
	c.DebtorAccount = new(CashAccount7)
	return c.DebtorAccount
}

func (c *CreditTransferTransactionInformation2) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification3 {
	c.DebtorAgent = new(BranchAndFinancialInstitutionIdentification3)
	return c.DebtorAgent
}

func (c *CreditTransferTransactionInformation2) AddDebtorAgentAccount() *CashAccount7 {
	c.DebtorAgentAccount = new(CashAccount7)
	return c.DebtorAgentAccount
}

func (c *CreditTransferTransactionInformation2) AddCreditorAgent() *BranchAndFinancialInstitutionIdentification3 {
	c.CreditorAgent = new(BranchAndFinancialInstitutionIdentification3)
	return c.CreditorAgent
}

func (c *CreditTransferTransactionInformation2) AddCreditorAgentAccount() *CashAccount7 {
	c.CreditorAgentAccount = new(CashAccount7)
	return c.CreditorAgentAccount
}

func (c *CreditTransferTransactionInformation2) AddCreditor() *PartyIdentification8 {
	c.Creditor = new(PartyIdentification8)
	return c.Creditor
}

func (c *CreditTransferTransactionInformation2) AddCreditorAccount() *CashAccount7 {
	c.CreditorAccount = new(CashAccount7)
	return c.CreditorAccount
}

func (c *CreditTransferTransactionInformation2) AddUltimateCreditor() *PartyIdentification8 {
	c.UltimateCreditor = new(PartyIdentification8)
	return c.UltimateCreditor
}

func (c *CreditTransferTransactionInformation2) AddInstructionForCreditorAgent() *InstructionForCreditorAgent1 {
	newValue := new(InstructionForCreditorAgent1)
	c.InstructionForCreditorAgent = append(c.InstructionForCreditorAgent, newValue)
	return newValue
}

func (c *CreditTransferTransactionInformation2) AddInstructionForNextAgent() *InstructionForNextAgent1 {
	newValue := new(InstructionForNextAgent1)
	c.InstructionForNextAgent = append(c.InstructionForNextAgent, newValue)
	return newValue
}

func (c *CreditTransferTransactionInformation2) AddPurpose() *Purpose1Choice {
	c.Purpose = new(Purpose1Choice)
	return c.Purpose
}

func (c *CreditTransferTransactionInformation2) AddRegulatoryReporting() *RegulatoryReporting2 {
	newValue := new(RegulatoryReporting2)
	c.RegulatoryReporting = append(c.RegulatoryReporting, newValue)
	return newValue
}

func (c *CreditTransferTransactionInformation2) AddRelatedRemittanceInformation() *RemittanceLocation1 {
	newValue := new(RemittanceLocation1)
	c.RelatedRemittanceInformation = append(c.RelatedRemittanceInformation, newValue)
	return newValue
}

func (c *CreditTransferTransactionInformation2) AddRemittanceInformation() *RemittanceInformation1 {
	c.RemittanceInformation = new(RemittanceInformation1)
	return c.RemittanceInformation
}
