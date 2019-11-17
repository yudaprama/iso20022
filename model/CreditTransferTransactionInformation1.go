package model

// Set of elements providing information specific to the individual transaction(s) included in the message.
type CreditTransferTransactionInformation1 struct {

	// Set of elements to reference a payment instruction.
	PaymentIdentification *PaymentIdentification1 `xml:"PmtId"`

	// Set of elements used to further specify the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation1 `xml:"PmtTpInf,omitempty"`

	// Amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as ordered by the initiating party.
	Amount *AmountType2Choice `xml:"Amt"`

	// Further detailed information on the exchange rate specified in the payment transaction.
	ExchangeRateInformation *ExchangeRateInformation1 `xml:"XchgRateInf,omitempty"`

	// Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr,omitempty"`

	// Information related to the issuance of a cheque.
	ChequeInstruction *Cheque5 `xml:"ChqInstr,omitempty"`

	// Ultimate party that owes an amount of money to the (ultimate) creditor.
	UltimateDebtor *PartyIdentification8 `xml:"UltmtDbtr,omitempty"`

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

	// Financial institution servicing an account for the creditor.
	CreditorAgent *BranchAndFinancialInstitutionIdentification3 `xml:"CdtrAgt,omitempty"`

	// Unambiguous identification of the account of the creditor agent at its servicing agent to which a credit entry will be made as a result of the payment transaction.
	CreditorAgentAccount *CashAccount7 `xml:"CdtrAgtAcct,omitempty"`

	// Party to which an amount of money is due.
	Creditor *PartyIdentification8 `xml:"Cdtr,omitempty"`

	// Unambiguous identification of the account of the creditor to which a credit entry will be posted as a result of the payment transaction.
	CreditorAccount *CashAccount7 `xml:"CdtrAcct,omitempty"`

	// Ultimate party to which an amount of money is due.
	UltimateCreditor *PartyIdentification8 `xml:"UltmtCdtr,omitempty"`

	// Further information related to the processing of the payment instruction, provided by the initiating party, and intended for the creditor agent.
	InstructionForCreditorAgent []*InstructionForCreditorAgent1 `xml:"InstrForCdtrAgt,omitempty"`

	// Further information related to the processing of the payment instruction that may need to be acted upon by the debtor agent, depending on agreement between debtor and the debtor agent.
	//
	// Usage: The instruction can relate to a level of service, can be an instruction to be executed by the debtor's agent, or can be information required by the debtor's agent to process the instruction.
	InstructionForDebtorAgent *Max140Text `xml:"InstrForDbtrAgt,omitempty"`

	// Underlying reason for the payment transaction, eg, a charity payment, or a commercial agreement between the creditor and the debtor.
	//
	// Usage: purpose is used by the end-customers, ie originating party, initiating party, debtor, creditor, final party, to provide information concerning the nature of the payment transaction. Purpose is a content element, which is not used for processing by any of the agents involved in the payment chain.
	Purpose *Purpose1Choice `xml:"Purp,omitempty"`

	// Information needed due to regulatory and statutory requirements.
	RegulatoryReporting []*RegulatoryReporting2 `xml:"RgltryRptg,omitempty"`

	// Amount of money due to the government or tax authority, according to various pre-defined parameters such as thresholds or income.
	Tax *TaxInformation2 `xml:"Tax,omitempty"`

	// Information related to the handling of the remittance information by any of the agents in the transaction processing chain.
	RelatedRemittanceInformation []*RemittanceLocation1 `xml:"RltdRmtInf,omitempty"`

	// Information that enables the matching, ie, reconciliation, of a payment with the items that the payment is intended to settle, eg, commercial invoices in an account receivable system.
	RemittanceInformation *RemittanceInformation1 `xml:"RmtInf,omitempty"`
}

func (c *CreditTransferTransactionInformation1) AddPaymentIdentification() *PaymentIdentification1 {
	c.PaymentIdentification = new(PaymentIdentification1)
	return c.PaymentIdentification
}

func (c *CreditTransferTransactionInformation1) AddPaymentTypeInformation() *PaymentTypeInformation1 {
	c.PaymentTypeInformation = new(PaymentTypeInformation1)
	return c.PaymentTypeInformation
}

func (c *CreditTransferTransactionInformation1) AddAmount() *AmountType2Choice {
	c.Amount = new(AmountType2Choice)
	return c.Amount
}

func (c *CreditTransferTransactionInformation1) AddExchangeRateInformation() *ExchangeRateInformation1 {
	c.ExchangeRateInformation = new(ExchangeRateInformation1)
	return c.ExchangeRateInformation
}

func (c *CreditTransferTransactionInformation1) SetChargeBearer(value string) {
	c.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (c *CreditTransferTransactionInformation1) AddChequeInstruction() *Cheque5 {
	c.ChequeInstruction = new(Cheque5)
	return c.ChequeInstruction
}

func (c *CreditTransferTransactionInformation1) AddUltimateDebtor() *PartyIdentification8 {
	c.UltimateDebtor = new(PartyIdentification8)
	return c.UltimateDebtor
}

func (c *CreditTransferTransactionInformation1) AddIntermediaryAgent1() *BranchAndFinancialInstitutionIdentification3 {
	c.IntermediaryAgent1 = new(BranchAndFinancialInstitutionIdentification3)
	return c.IntermediaryAgent1
}

func (c *CreditTransferTransactionInformation1) AddIntermediaryAgent1Account() *CashAccount7 {
	c.IntermediaryAgent1Account = new(CashAccount7)
	return c.IntermediaryAgent1Account
}

func (c *CreditTransferTransactionInformation1) AddIntermediaryAgent2() *BranchAndFinancialInstitutionIdentification3 {
	c.IntermediaryAgent2 = new(BranchAndFinancialInstitutionIdentification3)
	return c.IntermediaryAgent2
}

func (c *CreditTransferTransactionInformation1) AddIntermediaryAgent2Account() *CashAccount7 {
	c.IntermediaryAgent2Account = new(CashAccount7)
	return c.IntermediaryAgent2Account
}

func (c *CreditTransferTransactionInformation1) AddIntermediaryAgent3() *BranchAndFinancialInstitutionIdentification3 {
	c.IntermediaryAgent3 = new(BranchAndFinancialInstitutionIdentification3)
	return c.IntermediaryAgent3
}

func (c *CreditTransferTransactionInformation1) AddIntermediaryAgent3Account() *CashAccount7 {
	c.IntermediaryAgent3Account = new(CashAccount7)
	return c.IntermediaryAgent3Account
}

func (c *CreditTransferTransactionInformation1) AddCreditorAgent() *BranchAndFinancialInstitutionIdentification3 {
	c.CreditorAgent = new(BranchAndFinancialInstitutionIdentification3)
	return c.CreditorAgent
}

func (c *CreditTransferTransactionInformation1) AddCreditorAgentAccount() *CashAccount7 {
	c.CreditorAgentAccount = new(CashAccount7)
	return c.CreditorAgentAccount
}

func (c *CreditTransferTransactionInformation1) AddCreditor() *PartyIdentification8 {
	c.Creditor = new(PartyIdentification8)
	return c.Creditor
}

func (c *CreditTransferTransactionInformation1) AddCreditorAccount() *CashAccount7 {
	c.CreditorAccount = new(CashAccount7)
	return c.CreditorAccount
}

func (c *CreditTransferTransactionInformation1) AddUltimateCreditor() *PartyIdentification8 {
	c.UltimateCreditor = new(PartyIdentification8)
	return c.UltimateCreditor
}

func (c *CreditTransferTransactionInformation1) AddInstructionForCreditorAgent() *InstructionForCreditorAgent1 {
	newValue := new(InstructionForCreditorAgent1)
	c.InstructionForCreditorAgent = append(c.InstructionForCreditorAgent, newValue)
	return newValue
}

func (c *CreditTransferTransactionInformation1) SetInstructionForDebtorAgent(value string) {
	c.InstructionForDebtorAgent = (*Max140Text)(&value)
}

func (c *CreditTransferTransactionInformation1) AddPurpose() *Purpose1Choice {
	c.Purpose = new(Purpose1Choice)
	return c.Purpose
}

func (c *CreditTransferTransactionInformation1) AddRegulatoryReporting() *RegulatoryReporting2 {
	newValue := new(RegulatoryReporting2)
	c.RegulatoryReporting = append(c.RegulatoryReporting, newValue)
	return newValue
}

func (c *CreditTransferTransactionInformation1) AddTax() *TaxInformation2 {
	c.Tax = new(TaxInformation2)
	return c.Tax
}

func (c *CreditTransferTransactionInformation1) AddRelatedRemittanceInformation() *RemittanceLocation1 {
	newValue := new(RemittanceLocation1)
	c.RelatedRemittanceInformation = append(c.RelatedRemittanceInformation, newValue)
	return newValue
}

func (c *CreditTransferTransactionInformation1) AddRemittanceInformation() *RemittanceInformation1 {
	c.RemittanceInformation = new(RemittanceInformation1)
	return c.RemittanceInformation
}
