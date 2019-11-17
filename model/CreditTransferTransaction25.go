package model

// Provides further details specific to the individual transaction(s) included in the message.
type CreditTransferTransaction25 struct {

	// Set of elements used to reference a payment instruction.
	PaymentIdentification *PaymentIdentification3 `xml:"PmtId"`

	// Set of elements used to further specify the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation21 `xml:"PmtTpInf,omitempty"`

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

	// Point in time when the payment order from the initiating party meets the processing conditions of the account servicing agent. This means that the account servicing agent has received the payment order and has applied checks such as authorisation, availability of funds.
	AcceptanceDateTime *ISODateTime `xml:"AccptncDtTm,omitempty"`

	// Date used for the correction of the value date of a cash pool movement that has been posted with a different value date.
	PoolingAdjustmentDate *ISODate `xml:"PoolgAdjstmntDt,omitempty"`

	// Amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as ordered by the initiating party.
	// Usage: This amount has to be transported unchanged through the transaction chain.
	InstructedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"InstdAmt,omitempty"`

	// Factor used to convert an amount from one currency into another. This reflects the price at which one currency was bought with another currency.
	ExchangeRate *BaseOneRate `xml:"XchgRate,omitempty"`

	// Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr"`

	// Provides information on the charges to be paid by the charge bearer(s) related to the payment transaction.
	ChargesInformation []*Charges2 `xml:"ChrgsInf,omitempty"`

	// Agent immediately prior to the instructing agent.
	PreviousInstructingAgent *BranchAndFinancialInstitutionIdentification5 `xml:"PrvsInstgAgt,omitempty"`

	// Unambiguous identification of the account of the previous instructing agent at its servicing agent in the payment chain.
	PreviousInstructingAgentAccount *CashAccount24 `xml:"PrvsInstgAgtAcct,omitempty"`

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

	// Further information related to the processing of the payment instruction, provided by the initiating party, and intended for the creditor agent.
	InstructionForCreditorAgent []*InstructionForCreditorAgent1 `xml:"InstrForCdtrAgt,omitempty"`

	// Further information related to the processing of the payment instruction that may need to be acted upon by the next agent.
	//
	// Usage: The next agent may not be the creditor agent.
	// The instruction can relate to a level of service, can be an instruction that has to be executed by the agent, or can be information required by the next agent.
	InstructionForNextAgent []*InstructionForNextAgent1 `xml:"InstrForNxtAgt,omitempty"`

	// Underlying reason for the payment transaction.
	// Usage: Purpose is used by the end-customers, that is initiating party, (ultimate) debtor, (ultimate) creditor to provide information concerning the nature of the payment. Purpose is a content element, which is not used for processing by any of the agents involved in the payment chain.
	Purpose *Purpose2Choice `xml:"Purp,omitempty"`

	// Information needed due to regulatory and statutory requirements.
	RegulatoryReporting []*RegulatoryReporting3 `xml:"RgltryRptg,omitempty"`

	// Provides details on the tax.
	Tax *TaxInformation3 `xml:"Tax,omitempty"`

	// Provides information related to the handling of the remittance information by any of the agents in the transaction processing chain.
	RelatedRemittanceInformation []*RemittanceLocation4 `xml:"RltdRmtInf,omitempty"`

	// Information supplied to enable the matching of an entry with the items that the transfer is intended to settle, such as commercial invoices in an accounts' receivable system.
	RemittanceInformation *RemittanceInformation11 `xml:"RmtInf,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *CreditTransferTransaction25) AddPaymentIdentification() *PaymentIdentification3 {
	c.PaymentIdentification = new(PaymentIdentification3)
	return c.PaymentIdentification
}

func (c *CreditTransferTransaction25) AddPaymentTypeInformation() *PaymentTypeInformation21 {
	c.PaymentTypeInformation = new(PaymentTypeInformation21)
	return c.PaymentTypeInformation
}

func (c *CreditTransferTransaction25) SetInterbankSettlementAmount(value, currency string) {
	c.InterbankSettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CreditTransferTransaction25) SetInterbankSettlementDate(value string) {
	c.InterbankSettlementDate = (*ISODate)(&value)
}

func (c *CreditTransferTransaction25) SetSettlementPriority(value string) {
	c.SettlementPriority = (*Priority3Code)(&value)
}

func (c *CreditTransferTransaction25) AddSettlementTimeIndication() *SettlementDateTimeIndication1 {
	c.SettlementTimeIndication = new(SettlementDateTimeIndication1)
	return c.SettlementTimeIndication
}

func (c *CreditTransferTransaction25) AddSettlementTimeRequest() *SettlementTimeRequest2 {
	c.SettlementTimeRequest = new(SettlementTimeRequest2)
	return c.SettlementTimeRequest
}

func (c *CreditTransferTransaction25) SetAcceptanceDateTime(value string) {
	c.AcceptanceDateTime = (*ISODateTime)(&value)
}

func (c *CreditTransferTransaction25) SetPoolingAdjustmentDate(value string) {
	c.PoolingAdjustmentDate = (*ISODate)(&value)
}

func (c *CreditTransferTransaction25) SetInstructedAmount(value, currency string) {
	c.InstructedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (c *CreditTransferTransaction25) SetExchangeRate(value string) {
	c.ExchangeRate = (*BaseOneRate)(&value)
}

func (c *CreditTransferTransaction25) SetChargeBearer(value string) {
	c.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (c *CreditTransferTransaction25) AddChargesInformation() *Charges2 {
	newValue := new(Charges2)
	c.ChargesInformation = append(c.ChargesInformation, newValue)
	return newValue
}

func (c *CreditTransferTransaction25) AddPreviousInstructingAgent() *BranchAndFinancialInstitutionIdentification5 {
	c.PreviousInstructingAgent = new(BranchAndFinancialInstitutionIdentification5)
	return c.PreviousInstructingAgent
}

func (c *CreditTransferTransaction25) AddPreviousInstructingAgentAccount() *CashAccount24 {
	c.PreviousInstructingAgentAccount = new(CashAccount24)
	return c.PreviousInstructingAgentAccount
}

func (c *CreditTransferTransaction25) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification5 {
	c.InstructingAgent = new(BranchAndFinancialInstitutionIdentification5)
	return c.InstructingAgent
}

func (c *CreditTransferTransaction25) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification5 {
	c.InstructedAgent = new(BranchAndFinancialInstitutionIdentification5)
	return c.InstructedAgent
}

func (c *CreditTransferTransaction25) AddIntermediaryAgent1() *BranchAndFinancialInstitutionIdentification5 {
	c.IntermediaryAgent1 = new(BranchAndFinancialInstitutionIdentification5)
	return c.IntermediaryAgent1
}

func (c *CreditTransferTransaction25) AddIntermediaryAgent1Account() *CashAccount24 {
	c.IntermediaryAgent1Account = new(CashAccount24)
	return c.IntermediaryAgent1Account
}

func (c *CreditTransferTransaction25) AddIntermediaryAgent2() *BranchAndFinancialInstitutionIdentification5 {
	c.IntermediaryAgent2 = new(BranchAndFinancialInstitutionIdentification5)
	return c.IntermediaryAgent2
}

func (c *CreditTransferTransaction25) AddIntermediaryAgent2Account() *CashAccount24 {
	c.IntermediaryAgent2Account = new(CashAccount24)
	return c.IntermediaryAgent2Account
}

func (c *CreditTransferTransaction25) AddIntermediaryAgent3() *BranchAndFinancialInstitutionIdentification5 {
	c.IntermediaryAgent3 = new(BranchAndFinancialInstitutionIdentification5)
	return c.IntermediaryAgent3
}

func (c *CreditTransferTransaction25) AddIntermediaryAgent3Account() *CashAccount24 {
	c.IntermediaryAgent3Account = new(CashAccount24)
	return c.IntermediaryAgent3Account
}

func (c *CreditTransferTransaction25) AddUltimateDebtor() *PartyIdentification43 {
	c.UltimateDebtor = new(PartyIdentification43)
	return c.UltimateDebtor
}

func (c *CreditTransferTransaction25) AddInitiatingParty() *PartyIdentification43 {
	c.InitiatingParty = new(PartyIdentification43)
	return c.InitiatingParty
}

func (c *CreditTransferTransaction25) AddDebtor() *PartyIdentification43 {
	c.Debtor = new(PartyIdentification43)
	return c.Debtor
}

func (c *CreditTransferTransaction25) AddDebtorAccount() *CashAccount24 {
	c.DebtorAccount = new(CashAccount24)
	return c.DebtorAccount
}

func (c *CreditTransferTransaction25) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification5 {
	c.DebtorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return c.DebtorAgent
}

func (c *CreditTransferTransaction25) AddDebtorAgentAccount() *CashAccount24 {
	c.DebtorAgentAccount = new(CashAccount24)
	return c.DebtorAgentAccount
}

func (c *CreditTransferTransaction25) AddCreditorAgent() *BranchAndFinancialInstitutionIdentification5 {
	c.CreditorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return c.CreditorAgent
}

func (c *CreditTransferTransaction25) AddCreditorAgentAccount() *CashAccount24 {
	c.CreditorAgentAccount = new(CashAccount24)
	return c.CreditorAgentAccount
}

func (c *CreditTransferTransaction25) AddCreditor() *PartyIdentification43 {
	c.Creditor = new(PartyIdentification43)
	return c.Creditor
}

func (c *CreditTransferTransaction25) AddCreditorAccount() *CashAccount24 {
	c.CreditorAccount = new(CashAccount24)
	return c.CreditorAccount
}

func (c *CreditTransferTransaction25) AddUltimateCreditor() *PartyIdentification43 {
	c.UltimateCreditor = new(PartyIdentification43)
	return c.UltimateCreditor
}

func (c *CreditTransferTransaction25) AddInstructionForCreditorAgent() *InstructionForCreditorAgent1 {
	newValue := new(InstructionForCreditorAgent1)
	c.InstructionForCreditorAgent = append(c.InstructionForCreditorAgent, newValue)
	return newValue
}

func (c *CreditTransferTransaction25) AddInstructionForNextAgent() *InstructionForNextAgent1 {
	newValue := new(InstructionForNextAgent1)
	c.InstructionForNextAgent = append(c.InstructionForNextAgent, newValue)
	return newValue
}

func (c *CreditTransferTransaction25) AddPurpose() *Purpose2Choice {
	c.Purpose = new(Purpose2Choice)
	return c.Purpose
}

func (c *CreditTransferTransaction25) AddRegulatoryReporting() *RegulatoryReporting3 {
	newValue := new(RegulatoryReporting3)
	c.RegulatoryReporting = append(c.RegulatoryReporting, newValue)
	return newValue
}

func (c *CreditTransferTransaction25) AddTax() *TaxInformation3 {
	c.Tax = new(TaxInformation3)
	return c.Tax
}

func (c *CreditTransferTransaction25) AddRelatedRemittanceInformation() *RemittanceLocation4 {
	newValue := new(RemittanceLocation4)
	c.RelatedRemittanceInformation = append(c.RelatedRemittanceInformation, newValue)
	return newValue
}

func (c *CreditTransferTransaction25) AddRemittanceInformation() *RemittanceInformation11 {
	c.RemittanceInformation = new(RemittanceInformation11)
	return c.RemittanceInformation
}

func (c *CreditTransferTransaction25) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
