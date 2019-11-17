package model

// Set of elements used to provide information specific to the individual transaction(s) included in the message.
type CreditTransferTransactionInformation10 struct {

	// Set of elements used to reference a payment instruction.
	PaymentIdentification *PaymentIdentification1 `xml:"PmtId"`

	// Set of elements used to further specify the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation19 `xml:"PmtTpInf,omitempty"`

	// Amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as ordered by the initiating party.
	Amount *AmountType3Choice `xml:"Amt"`

	// Set of elements used to provide details on the currency exchange rate and contract.
	ExchangeRateInformation *ExchangeRateInformation1 `xml:"XchgRateInf,omitempty"`

	// Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr,omitempty"`

	// Set of elements needed to issue a cheque.
	ChequeInstruction *Cheque6 `xml:"ChqInstr,omitempty"`

	// Ultimate party that owes an amount of money to the (ultimate) creditor.
	UltimateDebtor *PartyIdentification32 `xml:"UltmtDbtr,omitempty"`

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
	CreditorAgent *BranchAndFinancialInstitutionIdentification4 `xml:"CdtrAgt,omitempty"`

	// Unambiguous identification of the account of the creditor agent at its servicing agent to which a credit entry will be made as a result of the payment transaction.
	CreditorAgentAccount *CashAccount16 `xml:"CdtrAgtAcct,omitempty"`

	// Party to which an amount of money is due.
	Creditor *PartyIdentification32 `xml:"Cdtr,omitempty"`

	// Unambiguous identification of the account of the creditor to which a credit entry will be posted as a result of the payment transaction.
	CreditorAccount *CashAccount16 `xml:"CdtrAcct,omitempty"`

	// Ultimate party to which an amount of money is due.
	UltimateCreditor *PartyIdentification32 `xml:"UltmtCdtr,omitempty"`

	// Further information related to the processing of the payment instruction, provided by the initiating party, and intended for the creditor agent.
	InstructionForCreditorAgent []*InstructionForCreditorAgent1 `xml:"InstrForCdtrAgt,omitempty"`

	// Further information related to the processing of the payment instruction, that may need to be acted upon by the debtor agent, depending on agreement between debtor and the debtor agent.
	InstructionForDebtorAgent *Max140Text `xml:"InstrForDbtrAgt,omitempty"`

	// Underlying reason for the payment transaction.
	// Usage: Purpose is used by the end-customers, that is initiating party, (ultimate) debtor, (ultimate) creditor to provide information concerning the nature of the payment. Purpose is a content element, which is not used for processing by any of the agents involved in the payment chain.
	Purpose *Purpose2Choice `xml:"Purp,omitempty"`

	// Information needed due to regulatory and statutory requirements.
	RegulatoryReporting []*RegulatoryReporting3 `xml:"RgltryRptg,omitempty"`

	// Set of elements used to provide details on the tax.
	Tax *TaxInformation3 `xml:"Tax,omitempty"`

	// Set of elements used to provide information related to the handling of the remittance information by any of the agents in the transaction processing chain.
	RelatedRemittanceInformation []*RemittanceLocation2 `xml:"RltdRmtInf,omitempty"`

	// Information supplied to enable the matching of an entry with the items that the transfer is intended to settle, such as commercial invoices in an accounts' receivable system.
	RemittanceInformation *RemittanceInformation5 `xml:"RmtInf,omitempty"`
}

func (c *CreditTransferTransactionInformation10) AddPaymentIdentification() *PaymentIdentification1 {
	c.PaymentIdentification = new(PaymentIdentification1)
	return c.PaymentIdentification
}

func (c *CreditTransferTransactionInformation10) AddPaymentTypeInformation() *PaymentTypeInformation19 {
	c.PaymentTypeInformation = new(PaymentTypeInformation19)
	return c.PaymentTypeInformation
}

func (c *CreditTransferTransactionInformation10) AddAmount() *AmountType3Choice {
	c.Amount = new(AmountType3Choice)
	return c.Amount
}

func (c *CreditTransferTransactionInformation10) AddExchangeRateInformation() *ExchangeRateInformation1 {
	c.ExchangeRateInformation = new(ExchangeRateInformation1)
	return c.ExchangeRateInformation
}

func (c *CreditTransferTransactionInformation10) SetChargeBearer(value string) {
	c.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (c *CreditTransferTransactionInformation10) AddChequeInstruction() *Cheque6 {
	c.ChequeInstruction = new(Cheque6)
	return c.ChequeInstruction
}

func (c *CreditTransferTransactionInformation10) AddUltimateDebtor() *PartyIdentification32 {
	c.UltimateDebtor = new(PartyIdentification32)
	return c.UltimateDebtor
}

func (c *CreditTransferTransactionInformation10) AddIntermediaryAgent1() *BranchAndFinancialInstitutionIdentification4 {
	c.IntermediaryAgent1 = new(BranchAndFinancialInstitutionIdentification4)
	return c.IntermediaryAgent1
}

func (c *CreditTransferTransactionInformation10) AddIntermediaryAgent1Account() *CashAccount16 {
	c.IntermediaryAgent1Account = new(CashAccount16)
	return c.IntermediaryAgent1Account
}

func (c *CreditTransferTransactionInformation10) AddIntermediaryAgent2() *BranchAndFinancialInstitutionIdentification4 {
	c.IntermediaryAgent2 = new(BranchAndFinancialInstitutionIdentification4)
	return c.IntermediaryAgent2
}

func (c *CreditTransferTransactionInformation10) AddIntermediaryAgent2Account() *CashAccount16 {
	c.IntermediaryAgent2Account = new(CashAccount16)
	return c.IntermediaryAgent2Account
}

func (c *CreditTransferTransactionInformation10) AddIntermediaryAgent3() *BranchAndFinancialInstitutionIdentification4 {
	c.IntermediaryAgent3 = new(BranchAndFinancialInstitutionIdentification4)
	return c.IntermediaryAgent3
}

func (c *CreditTransferTransactionInformation10) AddIntermediaryAgent3Account() *CashAccount16 {
	c.IntermediaryAgent3Account = new(CashAccount16)
	return c.IntermediaryAgent3Account
}

func (c *CreditTransferTransactionInformation10) AddCreditorAgent() *BranchAndFinancialInstitutionIdentification4 {
	c.CreditorAgent = new(BranchAndFinancialInstitutionIdentification4)
	return c.CreditorAgent
}

func (c *CreditTransferTransactionInformation10) AddCreditorAgentAccount() *CashAccount16 {
	c.CreditorAgentAccount = new(CashAccount16)
	return c.CreditorAgentAccount
}

func (c *CreditTransferTransactionInformation10) AddCreditor() *PartyIdentification32 {
	c.Creditor = new(PartyIdentification32)
	return c.Creditor
}

func (c *CreditTransferTransactionInformation10) AddCreditorAccount() *CashAccount16 {
	c.CreditorAccount = new(CashAccount16)
	return c.CreditorAccount
}

func (c *CreditTransferTransactionInformation10) AddUltimateCreditor() *PartyIdentification32 {
	c.UltimateCreditor = new(PartyIdentification32)
	return c.UltimateCreditor
}

func (c *CreditTransferTransactionInformation10) AddInstructionForCreditorAgent() *InstructionForCreditorAgent1 {
	newValue := new(InstructionForCreditorAgent1)
	c.InstructionForCreditorAgent = append(c.InstructionForCreditorAgent, newValue)
	return newValue
}

func (c *CreditTransferTransactionInformation10) SetInstructionForDebtorAgent(value string) {
	c.InstructionForDebtorAgent = (*Max140Text)(&value)
}

func (c *CreditTransferTransactionInformation10) AddPurpose() *Purpose2Choice {
	c.Purpose = new(Purpose2Choice)
	return c.Purpose
}

func (c *CreditTransferTransactionInformation10) AddRegulatoryReporting() *RegulatoryReporting3 {
	newValue := new(RegulatoryReporting3)
	c.RegulatoryReporting = append(c.RegulatoryReporting, newValue)
	return newValue
}

func (c *CreditTransferTransactionInformation10) AddTax() *TaxInformation3 {
	c.Tax = new(TaxInformation3)
	return c.Tax
}

func (c *CreditTransferTransactionInformation10) AddRelatedRemittanceInformation() *RemittanceLocation2 {
	newValue := new(RemittanceLocation2)
	c.RelatedRemittanceInformation = append(c.RelatedRemittanceInformation, newValue)
	return newValue
}

func (c *CreditTransferTransactionInformation10) AddRemittanceInformation() *RemittanceInformation5 {
	c.RemittanceInformation = new(RemittanceInformation5)
	return c.RemittanceInformation
}
