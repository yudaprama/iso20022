package model

// Set of elements used to provide information specific to the individual direct debit transaction(s) included in the message.
type DirectDebitTransactionInformation10 struct {

	// Set of elements used to reference a payment instruction.
	PaymentIdentification *PaymentIdentification3 `xml:"PmtId"`

	// Set of elements used to further specify the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation22 `xml:"PmtTpInf,omitempty"`

	// Amount of money moved between the instructing agent and the instructed agent.
	InterbankSettlementAmount *ActiveCurrencyAndAmount `xml:"IntrBkSttlmAmt"`

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	InterbankSettlementDate *ISODate `xml:"IntrBkSttlmDt,omitempty"`

	// Amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as ordered by the initiating party.
	InstructedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"InstdAmt,omitempty"`

	// Factor used to convert an amount from one currency into another. This reflects the price at which one currency was bought with another currency.
	ExchangeRate *BaseOneRate `xml:"XchgRate,omitempty"`

	// Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr"`

	// Provides information on the charges to be paid by the charge bearer(s) related to the payment transaction.
	ChargesInformation []*ChargesInformation5 `xml:"ChrgsInf,omitempty"`

	// Date and time at which the creditor requests that the amount of money is to be collected from the debtor.
	RequestedCollectionDate *ISODate `xml:"ReqdColltnDt,omitempty"`

	// Set of elements providing information specific to the direct debit mandate.
	DirectDebitTransaction *DirectDebitTransaction6 `xml:"DrctDbtTx,omitempty"`

	// Party to which an amount of money is due.
	Creditor *PartyIdentification32 `xml:"Cdtr"`

	// Unambiguous identification of the account of the creditor to which a credit entry will be posted as a result of the payment transaction.
	CreditorAccount *CashAccount16 `xml:"CdtrAcct,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *BranchAndFinancialInstitutionIdentification4 `xml:"CdtrAgt"`

	// Unambiguous identification of the account of the creditor agent at its servicing agent to which a credit entry will be made as a result of the payment transaction.
	CreditorAgentAccount *CashAccount16 `xml:"CdtrAgtAcct,omitempty"`

	// Ultimate party to which an amount of money is due.
	UltimateCreditor *PartyIdentification32 `xml:"UltmtCdtr,omitempty"`

	// Party that initiates the payment.
	// Usage: This can be either the creditor or a party that initiates the direct debit on behalf of the creditor.
	InitiatingParty *PartyIdentification32 `xml:"InitgPty,omitempty"`

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

	// Agent between the debtor agent and creditor agent.
	//
	// Usage: If IntermediaryAgent3 is present, then it identifies the agent between the intermediary agent 2 and the debtor agent.
	IntermediaryAgent3 *BranchAndFinancialInstitutionIdentification4 `xml:"IntrmyAgt3,omitempty"`

	// Unambiguous identification of the account of the intermediary agent 3 at its servicing agent in the payment chain.
	IntermediaryAgent3Account *CashAccount16 `xml:"IntrmyAgt3Acct,omitempty"`

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentification32 `xml:"Dbtr"`

	// Unambiguous identification of the account of the debtor to which a debit entry will be made as a result of the transaction.
	DebtorAccount *CashAccount16 `xml:"DbtrAcct"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification4 `xml:"DbtrAgt"`

	// Unambiguous identification of the account of the debtor agent at its servicing agent in the payment chain.
	DebtorAgentAccount *CashAccount16 `xml:"DbtrAgtAcct,omitempty"`

	// Ultimate party that owes an amount of money to the (ultimate) creditor.
	UltimateDebtor *PartyIdentification32 `xml:"UltmtDbtr,omitempty"`

	// Underlying reason for the payment transaction.
	// Usage: Purpose is used by the end-customers, that is initiating party, (ultimate) debtor, (ultimate) creditor to provide information concerning the nature of the payment. Purpose is a content element, which is not used for processing by any of the agents involved in the payment chain.
	Purpose *Purpose2Choice `xml:"Purp,omitempty"`

	// Information needed due to regulatory and statutory requirements.
	RegulatoryReporting []*RegulatoryReporting3 `xml:"RgltryRptg,omitempty"`

	// Set of elements used to provide information related to the handling of the remittance information by any of the agents in the transaction processing chain.
	RelatedRemittanceInformation []*RemittanceLocation2 `xml:"RltdRmtInf,omitempty"`

	// Information supplied to enable the matching of an entry with the items that the transfer is intended to settle, such as commercial invoices in an accounts' receivable system.
	RemittanceInformation *RemittanceInformation5 `xml:"RmtInf,omitempty"`
}

func (d *DirectDebitTransactionInformation10) AddPaymentIdentification() *PaymentIdentification3 {
	d.PaymentIdentification = new(PaymentIdentification3)
	return d.PaymentIdentification
}

func (d *DirectDebitTransactionInformation10) AddPaymentTypeInformation() *PaymentTypeInformation22 {
	d.PaymentTypeInformation = new(PaymentTypeInformation22)
	return d.PaymentTypeInformation
}

func (d *DirectDebitTransactionInformation10) SetInterbankSettlementAmount(value, currency string) {
	d.InterbankSettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (d *DirectDebitTransactionInformation10) SetInterbankSettlementDate(value string) {
	d.InterbankSettlementDate = (*ISODate)(&value)
}

func (d *DirectDebitTransactionInformation10) SetInstructedAmount(value, currency string) {
	d.InstructedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (d *DirectDebitTransactionInformation10) SetExchangeRate(value string) {
	d.ExchangeRate = (*BaseOneRate)(&value)
}

func (d *DirectDebitTransactionInformation10) SetChargeBearer(value string) {
	d.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (d *DirectDebitTransactionInformation10) AddChargesInformation() *ChargesInformation5 {
	newValue := new(ChargesInformation5)
	d.ChargesInformation = append(d.ChargesInformation, newValue)
	return newValue
}

func (d *DirectDebitTransactionInformation10) SetRequestedCollectionDate(value string) {
	d.RequestedCollectionDate = (*ISODate)(&value)
}

func (d *DirectDebitTransactionInformation10) AddDirectDebitTransaction() *DirectDebitTransaction6 {
	d.DirectDebitTransaction = new(DirectDebitTransaction6)
	return d.DirectDebitTransaction
}

func (d *DirectDebitTransactionInformation10) AddCreditor() *PartyIdentification32 {
	d.Creditor = new(PartyIdentification32)
	return d.Creditor
}

func (d *DirectDebitTransactionInformation10) AddCreditorAccount() *CashAccount16 {
	d.CreditorAccount = new(CashAccount16)
	return d.CreditorAccount
}

func (d *DirectDebitTransactionInformation10) AddCreditorAgent() *BranchAndFinancialInstitutionIdentification4 {
	d.CreditorAgent = new(BranchAndFinancialInstitutionIdentification4)
	return d.CreditorAgent
}

func (d *DirectDebitTransactionInformation10) AddCreditorAgentAccount() *CashAccount16 {
	d.CreditorAgentAccount = new(CashAccount16)
	return d.CreditorAgentAccount
}

func (d *DirectDebitTransactionInformation10) AddUltimateCreditor() *PartyIdentification32 {
	d.UltimateCreditor = new(PartyIdentification32)
	return d.UltimateCreditor
}

func (d *DirectDebitTransactionInformation10) AddInitiatingParty() *PartyIdentification32 {
	d.InitiatingParty = new(PartyIdentification32)
	return d.InitiatingParty
}

func (d *DirectDebitTransactionInformation10) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification4 {
	d.InstructingAgent = new(BranchAndFinancialInstitutionIdentification4)
	return d.InstructingAgent
}

func (d *DirectDebitTransactionInformation10) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification4 {
	d.InstructedAgent = new(BranchAndFinancialInstitutionIdentification4)
	return d.InstructedAgent
}

func (d *DirectDebitTransactionInformation10) AddIntermediaryAgent1() *BranchAndFinancialInstitutionIdentification4 {
	d.IntermediaryAgent1 = new(BranchAndFinancialInstitutionIdentification4)
	return d.IntermediaryAgent1
}

func (d *DirectDebitTransactionInformation10) AddIntermediaryAgent1Account() *CashAccount16 {
	d.IntermediaryAgent1Account = new(CashAccount16)
	return d.IntermediaryAgent1Account
}

func (d *DirectDebitTransactionInformation10) AddIntermediaryAgent2() *BranchAndFinancialInstitutionIdentification4 {
	d.IntermediaryAgent2 = new(BranchAndFinancialInstitutionIdentification4)
	return d.IntermediaryAgent2
}

func (d *DirectDebitTransactionInformation10) AddIntermediaryAgent2Account() *CashAccount16 {
	d.IntermediaryAgent2Account = new(CashAccount16)
	return d.IntermediaryAgent2Account
}

func (d *DirectDebitTransactionInformation10) AddIntermediaryAgent3() *BranchAndFinancialInstitutionIdentification4 {
	d.IntermediaryAgent3 = new(BranchAndFinancialInstitutionIdentification4)
	return d.IntermediaryAgent3
}

func (d *DirectDebitTransactionInformation10) AddIntermediaryAgent3Account() *CashAccount16 {
	d.IntermediaryAgent3Account = new(CashAccount16)
	return d.IntermediaryAgent3Account
}

func (d *DirectDebitTransactionInformation10) AddDebtor() *PartyIdentification32 {
	d.Debtor = new(PartyIdentification32)
	return d.Debtor
}

func (d *DirectDebitTransactionInformation10) AddDebtorAccount() *CashAccount16 {
	d.DebtorAccount = new(CashAccount16)
	return d.DebtorAccount
}

func (d *DirectDebitTransactionInformation10) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification4 {
	d.DebtorAgent = new(BranchAndFinancialInstitutionIdentification4)
	return d.DebtorAgent
}

func (d *DirectDebitTransactionInformation10) AddDebtorAgentAccount() *CashAccount16 {
	d.DebtorAgentAccount = new(CashAccount16)
	return d.DebtorAgentAccount
}

func (d *DirectDebitTransactionInformation10) AddUltimateDebtor() *PartyIdentification32 {
	d.UltimateDebtor = new(PartyIdentification32)
	return d.UltimateDebtor
}

func (d *DirectDebitTransactionInformation10) AddPurpose() *Purpose2Choice {
	d.Purpose = new(Purpose2Choice)
	return d.Purpose
}

func (d *DirectDebitTransactionInformation10) AddRegulatoryReporting() *RegulatoryReporting3 {
	newValue := new(RegulatoryReporting3)
	d.RegulatoryReporting = append(d.RegulatoryReporting, newValue)
	return newValue
}

func (d *DirectDebitTransactionInformation10) AddRelatedRemittanceInformation() *RemittanceLocation2 {
	newValue := new(RemittanceLocation2)
	d.RelatedRemittanceInformation = append(d.RelatedRemittanceInformation, newValue)
	return newValue
}

func (d *DirectDebitTransactionInformation10) AddRemittanceInformation() *RemittanceInformation5 {
	d.RemittanceInformation = new(RemittanceInformation5)
	return d.RemittanceInformation
}
