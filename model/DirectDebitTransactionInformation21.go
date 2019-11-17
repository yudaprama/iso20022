package model

// Provides further details specific to the individual direct debit transaction(s) included in the message.
type DirectDebitTransactionInformation21 struct {

	// Set of elements used to reference a payment instruction.
	PaymentIdentification *PaymentIdentification3 `xml:"PmtId"`

	// Set of elements used to further specify the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation25 `xml:"PmtTpInf,omitempty"`

	// Amount of money moved between the instructing agent and the instructed agent.
	InterbankSettlementAmount *ActiveCurrencyAndAmount `xml:"IntrBkSttlmAmt"`

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	InterbankSettlementDate *ISODate `xml:"IntrBkSttlmDt,omitempty"`

	// Indicator of the urgency or order of importance that the instructing party would like the instructed party to apply to the processing of the settlement instruction.
	SettlementPriority *Priority3Code `xml:"SttlmPrty,omitempty"`

	// Amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as ordered by the initiating party.
	// Usage: This amount has to be transported unchanged through the transaction chain.
	InstructedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"InstdAmt,omitempty"`

	// Factor used to convert an amount from one currency into another. This reflects the price at which one currency was bought with another currency.
	ExchangeRate *BaseOneRate `xml:"XchgRate,omitempty"`

	// Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr"`

	// Provides information on the charges to be paid by the charge bearer(s) related to the payment transaction.
	ChargesInformation []*Charges2 `xml:"ChrgsInf,omitempty"`

	// Date and time at which the creditor requests that the amount of money is to be collected from the debtor.
	RequestedCollectionDate *ISODate `xml:"ReqdColltnDt,omitempty"`

	// Provides information specific to the direct debit mandate.
	DirectDebitTransaction *DirectDebitTransaction9 `xml:"DrctDbtTx,omitempty"`

	// Party to which an amount of money is due.
	Creditor *PartyIdentification43 `xml:"Cdtr"`

	// Unambiguous identification of the account of the creditor to which a credit entry will be posted as a result of the payment transaction.
	CreditorAccount *CashAccount24 `xml:"CdtrAcct,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"CdtrAgt"`

	// Unambiguous identification of the account of the creditor agent at its servicing agent to which a credit entry will be made as a result of the payment transaction.
	CreditorAgentAccount *CashAccount24 `xml:"CdtrAgtAcct,omitempty"`

	// Ultimate party to which an amount of money is due.
	UltimateCreditor *PartyIdentification43 `xml:"UltmtCdtr,omitempty"`

	// Party that initiates the payment.
	// Usage: This can be either the creditor or a party that initiates the direct debit on behalf of the creditor.
	InitiatingParty *PartyIdentification43 `xml:"InitgPty,omitempty"`

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

	// Agent between the debtor agent and creditor agent.
	//
	// Usage: If IntermediaryAgent3 is present, then it identifies the agent between the intermediary agent 2 and the debtor agent.
	IntermediaryAgent3 *BranchAndFinancialInstitutionIdentification5 `xml:"IntrmyAgt3,omitempty"`

	// Unambiguous identification of the account of the intermediary agent 3 at its servicing agent in the payment chain.
	IntermediaryAgent3Account *CashAccount24 `xml:"IntrmyAgt3Acct,omitempty"`

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentification43 `xml:"Dbtr"`

	// Unambiguous identification of the account of the debtor to which a debit entry will be made as a result of the transaction.
	DebtorAccount *CashAccount24 `xml:"DbtrAcct"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"DbtrAgt"`

	// Unambiguous identification of the account of the debtor agent at its servicing agent in the payment chain.
	DebtorAgentAccount *CashAccount24 `xml:"DbtrAgtAcct,omitempty"`

	// Ultimate party that owes an amount of money to the (ultimate) creditor.
	UltimateDebtor *PartyIdentification43 `xml:"UltmtDbtr,omitempty"`

	// Underlying reason for the payment transaction.
	// Usage: Purpose is used by the end-customers, that is initiating party, (ultimate) debtor, (ultimate) creditor to provide information concerning the nature of the payment. Purpose is a content element, which is not used for processing by any of the agents involved in the payment chain.
	Purpose *Purpose2Choice `xml:"Purp,omitempty"`

	// Information needed due to regulatory and statutory requirements.
	RegulatoryReporting []*RegulatoryReporting3 `xml:"RgltryRptg,omitempty"`

	// Provides information related to the handling of the remittance information by any of the agents in the transaction processing chain.
	RelatedRemittanceInformation []*RemittanceLocation4 `xml:"RltdRmtInf,omitempty"`

	// Information supplied to enable the matching of an entry with the items that the transfer is intended to settle, such as commercial invoices in an accounts' receivable system.
	RemittanceInformation *RemittanceInformation11 `xml:"RmtInf,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (d *DirectDebitTransactionInformation21) AddPaymentIdentification() *PaymentIdentification3 {
	d.PaymentIdentification = new(PaymentIdentification3)
	return d.PaymentIdentification
}

func (d *DirectDebitTransactionInformation21) AddPaymentTypeInformation() *PaymentTypeInformation25 {
	d.PaymentTypeInformation = new(PaymentTypeInformation25)
	return d.PaymentTypeInformation
}

func (d *DirectDebitTransactionInformation21) SetInterbankSettlementAmount(value, currency string) {
	d.InterbankSettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (d *DirectDebitTransactionInformation21) SetInterbankSettlementDate(value string) {
	d.InterbankSettlementDate = (*ISODate)(&value)
}

func (d *DirectDebitTransactionInformation21) SetSettlementPriority(value string) {
	d.SettlementPriority = (*Priority3Code)(&value)
}

func (d *DirectDebitTransactionInformation21) SetInstructedAmount(value, currency string) {
	d.InstructedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (d *DirectDebitTransactionInformation21) SetExchangeRate(value string) {
	d.ExchangeRate = (*BaseOneRate)(&value)
}

func (d *DirectDebitTransactionInformation21) SetChargeBearer(value string) {
	d.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (d *DirectDebitTransactionInformation21) AddChargesInformation() *Charges2 {
	newValue := new(Charges2)
	d.ChargesInformation = append(d.ChargesInformation, newValue)
	return newValue
}

func (d *DirectDebitTransactionInformation21) SetRequestedCollectionDate(value string) {
	d.RequestedCollectionDate = (*ISODate)(&value)
}

func (d *DirectDebitTransactionInformation21) AddDirectDebitTransaction() *DirectDebitTransaction9 {
	d.DirectDebitTransaction = new(DirectDebitTransaction9)
	return d.DirectDebitTransaction
}

func (d *DirectDebitTransactionInformation21) AddCreditor() *PartyIdentification43 {
	d.Creditor = new(PartyIdentification43)
	return d.Creditor
}

func (d *DirectDebitTransactionInformation21) AddCreditorAccount() *CashAccount24 {
	d.CreditorAccount = new(CashAccount24)
	return d.CreditorAccount
}

func (d *DirectDebitTransactionInformation21) AddCreditorAgent() *BranchAndFinancialInstitutionIdentification5 {
	d.CreditorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return d.CreditorAgent
}

func (d *DirectDebitTransactionInformation21) AddCreditorAgentAccount() *CashAccount24 {
	d.CreditorAgentAccount = new(CashAccount24)
	return d.CreditorAgentAccount
}

func (d *DirectDebitTransactionInformation21) AddUltimateCreditor() *PartyIdentification43 {
	d.UltimateCreditor = new(PartyIdentification43)
	return d.UltimateCreditor
}

func (d *DirectDebitTransactionInformation21) AddInitiatingParty() *PartyIdentification43 {
	d.InitiatingParty = new(PartyIdentification43)
	return d.InitiatingParty
}

func (d *DirectDebitTransactionInformation21) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification5 {
	d.InstructingAgent = new(BranchAndFinancialInstitutionIdentification5)
	return d.InstructingAgent
}

func (d *DirectDebitTransactionInformation21) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification5 {
	d.InstructedAgent = new(BranchAndFinancialInstitutionIdentification5)
	return d.InstructedAgent
}

func (d *DirectDebitTransactionInformation21) AddIntermediaryAgent1() *BranchAndFinancialInstitutionIdentification5 {
	d.IntermediaryAgent1 = new(BranchAndFinancialInstitutionIdentification5)
	return d.IntermediaryAgent1
}

func (d *DirectDebitTransactionInformation21) AddIntermediaryAgent1Account() *CashAccount24 {
	d.IntermediaryAgent1Account = new(CashAccount24)
	return d.IntermediaryAgent1Account
}

func (d *DirectDebitTransactionInformation21) AddIntermediaryAgent2() *BranchAndFinancialInstitutionIdentification5 {
	d.IntermediaryAgent2 = new(BranchAndFinancialInstitutionIdentification5)
	return d.IntermediaryAgent2
}

func (d *DirectDebitTransactionInformation21) AddIntermediaryAgent2Account() *CashAccount24 {
	d.IntermediaryAgent2Account = new(CashAccount24)
	return d.IntermediaryAgent2Account
}

func (d *DirectDebitTransactionInformation21) AddIntermediaryAgent3() *BranchAndFinancialInstitutionIdentification5 {
	d.IntermediaryAgent3 = new(BranchAndFinancialInstitutionIdentification5)
	return d.IntermediaryAgent3
}

func (d *DirectDebitTransactionInformation21) AddIntermediaryAgent3Account() *CashAccount24 {
	d.IntermediaryAgent3Account = new(CashAccount24)
	return d.IntermediaryAgent3Account
}

func (d *DirectDebitTransactionInformation21) AddDebtor() *PartyIdentification43 {
	d.Debtor = new(PartyIdentification43)
	return d.Debtor
}

func (d *DirectDebitTransactionInformation21) AddDebtorAccount() *CashAccount24 {
	d.DebtorAccount = new(CashAccount24)
	return d.DebtorAccount
}

func (d *DirectDebitTransactionInformation21) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification5 {
	d.DebtorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return d.DebtorAgent
}

func (d *DirectDebitTransactionInformation21) AddDebtorAgentAccount() *CashAccount24 {
	d.DebtorAgentAccount = new(CashAccount24)
	return d.DebtorAgentAccount
}

func (d *DirectDebitTransactionInformation21) AddUltimateDebtor() *PartyIdentification43 {
	d.UltimateDebtor = new(PartyIdentification43)
	return d.UltimateDebtor
}

func (d *DirectDebitTransactionInformation21) AddPurpose() *Purpose2Choice {
	d.Purpose = new(Purpose2Choice)
	return d.Purpose
}

func (d *DirectDebitTransactionInformation21) AddRegulatoryReporting() *RegulatoryReporting3 {
	newValue := new(RegulatoryReporting3)
	d.RegulatoryReporting = append(d.RegulatoryReporting, newValue)
	return newValue
}

func (d *DirectDebitTransactionInformation21) AddRelatedRemittanceInformation() *RemittanceLocation4 {
	newValue := new(RemittanceLocation4)
	d.RelatedRemittanceInformation = append(d.RelatedRemittanceInformation, newValue)
	return newValue
}

func (d *DirectDebitTransactionInformation21) AddRemittanceInformation() *RemittanceInformation11 {
	d.RemittanceInformation = new(RemittanceInformation11)
	return d.RemittanceInformation
}

func (d *DirectDebitTransactionInformation21) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	d.SupplementaryData = append(d.SupplementaryData, newValue)
	return newValue
}
