package model

// Set of characteristics that apply to the the direct debit transaction(s).
type DirectDebitTransactionInformation2 struct {

	// Set of elements to reference a payment instruction.
	PaymentIdentification *PaymentIdentification2 `xml:"PmtId"`

	// Set of elements used to further specify the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation4 `xml:"PmtTpInf,omitempty"`

	// Amount of money moved between the instructing agent and the instructed agent.
	InterbankSettlementAmount *CurrencyAndAmount `xml:"IntrBkSttlmAmt"`

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	InterbankSettlementDate *ISODate `xml:"IntrBkSttlmDt,omitempty"`

	// Amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as ordered by the initiating party.
	InstructedAmount *CurrencyAndAmount `xml:"InstdAmt,omitempty"`

	// The factor used for conversion of an amount from one currency into another. This reflects the price at which one currency was bought with another currency.
	ExchangeRate *BaseOneRate `xml:"XchgRate,omitempty"`

	// Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr"`

	// Provides information on the charges related to the payment transaction.
	ChargesInformation []*ChargesInformation1 `xml:"ChrgsInf,omitempty"`

	// Date at which the creditor requests the amount of money to be collected from the debtor.
	RequestedCollectionDate *ISODate `xml:"ReqdColltnDt,omitempty"`

	// Set of elements providing information specific to the direct debit mandate.
	DirectDebitTransaction *DirectDebitTransaction1 `xml:"DrctDbtTx,omitempty"`

	// Party to which an amount of money is due.
	Creditor *PartyIdentification8 `xml:"Cdtr"`

	// Unambiguous identification of the account of the creditor to which a credit entry will be posted as a result of the payment transaction.
	CreditorAccount *CashAccount7 `xml:"CdtrAcct,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *BranchAndFinancialInstitutionIdentification3 `xml:"CdtrAgt"`

	// Unambiguous identification of the account of the creditor agent at its servicing agent to which a credit entry will be made as a result of the payment transaction.
	CreditorAgentAccount *CashAccount7 `xml:"CdtrAgtAcct,omitempty"`

	// Ultimate party to which an amount of money is due.
	UltimateCreditor *PartyIdentification8 `xml:"UltmtCdtr,omitempty"`

	// Party that initiates the payment. In the payment context, this can either be the debtor (in a credit transfer), the creditor (in a direct debit), or a party that initiates the payment on behalf of the debtor or creditor.
	InitiatingParty *PartyIdentification8 `xml:"InitgPty,omitempty"`

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	InstructingAgent *BranchAndFinancialInstitutionIdentification3 `xml:"InstgAgt,omitempty"`

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	InstructedAgent *BranchAndFinancialInstitutionIdentification3 `xml:"InstdAgt,omitempty"`

	// Agent between the debtor agent and creditor agent.
	//
	// Usage: If more than one intermediary agent is present, then IntermediaryAgent1 identifies the agent between the creditor agent and the intermediary agent 2.
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
	// Usage: If IntermediaryAgent3 is present, then it identifies the agent between the intermediary agent 2 and the debtor agent.
	IntermediaryAgent3 *BranchAndFinancialInstitutionIdentification3 `xml:"IntrmyAgt3,omitempty"`

	// Unambiguous identification of the account of the intermediary agent 3 at its servicing agent in the payment chain.
	IntermediaryAgent3Account *CashAccount7 `xml:"IntrmyAgt3Acct,omitempty"`

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentification8 `xml:"Dbtr"`

	// Unambiguous identification of the account of the debtor to which a debit entry will be made as a result of the transaction.
	DebtorAccount *CashAccount7 `xml:"DbtrAcct"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification3 `xml:"DbtrAgt"`

	// Unambiguous identification of the account of the debtor agent at its servicing agent in the payment chain.
	DebtorAgentAccount *CashAccount7 `xml:"DbtrAgtAcct,omitempty"`

	// Ultimate party that owes an amount of money to the (ultimate) creditor.
	UltimateDebtor *PartyIdentification8 `xml:"UltmtDbtr,omitempty"`

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

func (d *DirectDebitTransactionInformation2) AddPaymentIdentification() *PaymentIdentification2 {
	d.PaymentIdentification = new(PaymentIdentification2)
	return d.PaymentIdentification
}

func (d *DirectDebitTransactionInformation2) AddPaymentTypeInformation() *PaymentTypeInformation4 {
	d.PaymentTypeInformation = new(PaymentTypeInformation4)
	return d.PaymentTypeInformation
}

func (d *DirectDebitTransactionInformation2) SetInterbankSettlementAmount(value, currency string) {
	d.InterbankSettlementAmount = NewCurrencyAndAmount(value, currency)
}

func (d *DirectDebitTransactionInformation2) SetInterbankSettlementDate(value string) {
	d.InterbankSettlementDate = (*ISODate)(&value)
}

func (d *DirectDebitTransactionInformation2) SetInstructedAmount(value, currency string) {
	d.InstructedAmount = NewCurrencyAndAmount(value, currency)
}

func (d *DirectDebitTransactionInformation2) SetExchangeRate(value string) {
	d.ExchangeRate = (*BaseOneRate)(&value)
}

func (d *DirectDebitTransactionInformation2) SetChargeBearer(value string) {
	d.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (d *DirectDebitTransactionInformation2) AddChargesInformation() *ChargesInformation1 {
	newValue := new(ChargesInformation1)
	d.ChargesInformation = append(d.ChargesInformation, newValue)
	return newValue
}

func (d *DirectDebitTransactionInformation2) SetRequestedCollectionDate(value string) {
	d.RequestedCollectionDate = (*ISODate)(&value)
}

func (d *DirectDebitTransactionInformation2) AddDirectDebitTransaction() *DirectDebitTransaction1 {
	d.DirectDebitTransaction = new(DirectDebitTransaction1)
	return d.DirectDebitTransaction
}

func (d *DirectDebitTransactionInformation2) AddCreditor() *PartyIdentification8 {
	d.Creditor = new(PartyIdentification8)
	return d.Creditor
}

func (d *DirectDebitTransactionInformation2) AddCreditorAccount() *CashAccount7 {
	d.CreditorAccount = new(CashAccount7)
	return d.CreditorAccount
}

func (d *DirectDebitTransactionInformation2) AddCreditorAgent() *BranchAndFinancialInstitutionIdentification3 {
	d.CreditorAgent = new(BranchAndFinancialInstitutionIdentification3)
	return d.CreditorAgent
}

func (d *DirectDebitTransactionInformation2) AddCreditorAgentAccount() *CashAccount7 {
	d.CreditorAgentAccount = new(CashAccount7)
	return d.CreditorAgentAccount
}

func (d *DirectDebitTransactionInformation2) AddUltimateCreditor() *PartyIdentification8 {
	d.UltimateCreditor = new(PartyIdentification8)
	return d.UltimateCreditor
}

func (d *DirectDebitTransactionInformation2) AddInitiatingParty() *PartyIdentification8 {
	d.InitiatingParty = new(PartyIdentification8)
	return d.InitiatingParty
}

func (d *DirectDebitTransactionInformation2) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification3 {
	d.InstructingAgent = new(BranchAndFinancialInstitutionIdentification3)
	return d.InstructingAgent
}

func (d *DirectDebitTransactionInformation2) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification3 {
	d.InstructedAgent = new(BranchAndFinancialInstitutionIdentification3)
	return d.InstructedAgent
}

func (d *DirectDebitTransactionInformation2) AddIntermediaryAgent1() *BranchAndFinancialInstitutionIdentification3 {
	d.IntermediaryAgent1 = new(BranchAndFinancialInstitutionIdentification3)
	return d.IntermediaryAgent1
}

func (d *DirectDebitTransactionInformation2) AddIntermediaryAgent1Account() *CashAccount7 {
	d.IntermediaryAgent1Account = new(CashAccount7)
	return d.IntermediaryAgent1Account
}

func (d *DirectDebitTransactionInformation2) AddIntermediaryAgent2() *BranchAndFinancialInstitutionIdentification3 {
	d.IntermediaryAgent2 = new(BranchAndFinancialInstitutionIdentification3)
	return d.IntermediaryAgent2
}

func (d *DirectDebitTransactionInformation2) AddIntermediaryAgent2Account() *CashAccount7 {
	d.IntermediaryAgent2Account = new(CashAccount7)
	return d.IntermediaryAgent2Account
}

func (d *DirectDebitTransactionInformation2) AddIntermediaryAgent3() *BranchAndFinancialInstitutionIdentification3 {
	d.IntermediaryAgent3 = new(BranchAndFinancialInstitutionIdentification3)
	return d.IntermediaryAgent3
}

func (d *DirectDebitTransactionInformation2) AddIntermediaryAgent3Account() *CashAccount7 {
	d.IntermediaryAgent3Account = new(CashAccount7)
	return d.IntermediaryAgent3Account
}

func (d *DirectDebitTransactionInformation2) AddDebtor() *PartyIdentification8 {
	d.Debtor = new(PartyIdentification8)
	return d.Debtor
}

func (d *DirectDebitTransactionInformation2) AddDebtorAccount() *CashAccount7 {
	d.DebtorAccount = new(CashAccount7)
	return d.DebtorAccount
}

func (d *DirectDebitTransactionInformation2) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification3 {
	d.DebtorAgent = new(BranchAndFinancialInstitutionIdentification3)
	return d.DebtorAgent
}

func (d *DirectDebitTransactionInformation2) AddDebtorAgentAccount() *CashAccount7 {
	d.DebtorAgentAccount = new(CashAccount7)
	return d.DebtorAgentAccount
}

func (d *DirectDebitTransactionInformation2) AddUltimateDebtor() *PartyIdentification8 {
	d.UltimateDebtor = new(PartyIdentification8)
	return d.UltimateDebtor
}

func (d *DirectDebitTransactionInformation2) AddPurpose() *Purpose1Choice {
	d.Purpose = new(Purpose1Choice)
	return d.Purpose
}

func (d *DirectDebitTransactionInformation2) AddRegulatoryReporting() *RegulatoryReporting2 {
	newValue := new(RegulatoryReporting2)
	d.RegulatoryReporting = append(d.RegulatoryReporting, newValue)
	return newValue
}

func (d *DirectDebitTransactionInformation2) AddRelatedRemittanceInformation() *RemittanceLocation1 {
	newValue := new(RemittanceLocation1)
	d.RelatedRemittanceInformation = append(d.RelatedRemittanceInformation, newValue)
	return newValue
}

func (d *DirectDebitTransactionInformation2) AddRemittanceInformation() *RemittanceInformation1 {
	d.RemittanceInformation = new(RemittanceInformation1)
	return d.RemittanceInformation
}
