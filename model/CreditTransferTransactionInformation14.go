package model

// Set of elements used to provide information specific to the individual transaction(s) included in the message.
type CreditTransferTransactionInformation14 struct {

	// Set of elements used to reference a payment instruction.
	PaymentIdentification *PaymentIdentification1 `xml:"PmtId"`

	// Set of elements used to further specify the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation19 `xml:"PmtTpInf,omitempty"`

	// Amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as ordered by the initiating party.
	Amount *AmountType3Choice `xml:"Amt"`

	// Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr"`

	// Set of elements needed to issue a cheque.
	ChequeInstruction *Cheque6 `xml:"ChqInstr,omitempty"`

	// Ultimate party that owes an amount of money to the (ultimate) creditor.
	UltimateDebtor *PartyIdentification43 `xml:"UltmtDbtr,omitempty"`

	// Agent between the debtor's agent and the creditor's agent.
	//
	// Usage: If more than one intermediary agent is present, then IntermediaryAgent1 identifies the agent between the DebtorAgent and the IntermediaryAgent2.
	IntermediaryAgent1 *BranchAndFinancialInstitutionIdentification5 `xml:"IntrmyAgt1,omitempty"`

	// Agent between the debtor's agent and the creditor's agent.
	//
	// Usage: If more than two intermediary agents are present, then IntermediaryAgent2 identifies the agent between the IntermediaryAgent1 and the IntermediaryAgent3.
	IntermediaryAgent2 *BranchAndFinancialInstitutionIdentification5 `xml:"IntrmyAgt2,omitempty"`

	// Agent between the debtor's agent and the creditor's agent.
	//
	// Usage: If IntermediaryAgent3 is present, then it identifies the agent between the IntermediaryAgent 2 and the CreditorAgent.
	IntermediaryAgent3 *BranchAndFinancialInstitutionIdentification5 `xml:"IntrmyAgt3,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"CdtrAgt"`

	// Party to which an amount of money is due.
	Creditor *PartyIdentification43 `xml:"Cdtr"`

	// Unambiguous identification of the account of the creditor to which a credit entry will be posted as a result of the payment transaction.
	CreditorAccount *CashAccount16 `xml:"CdtrAcct,omitempty"`

	// Ultimate party to which an amount of money is due.
	UltimateCreditor *PartyIdentification43 `xml:"UltmtCdtr,omitempty"`

	// Further information related to the processing of the payment instruction, provided by the initiating party, and intended for the creditor agent.
	InstructionForCreditorAgent []*InstructionForCreditorAgent1 `xml:"InstrForCdtrAgt,omitempty"`

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
	RemittanceInformation *RemittanceInformation6 `xml:"RmtInf,omitempty"`
}

func (c *CreditTransferTransactionInformation14) AddPaymentIdentification() *PaymentIdentification1 {
	c.PaymentIdentification = new(PaymentIdentification1)
	return c.PaymentIdentification
}

func (c *CreditTransferTransactionInformation14) AddPaymentTypeInformation() *PaymentTypeInformation19 {
	c.PaymentTypeInformation = new(PaymentTypeInformation19)
	return c.PaymentTypeInformation
}

func (c *CreditTransferTransactionInformation14) AddAmount() *AmountType3Choice {
	c.Amount = new(AmountType3Choice)
	return c.Amount
}

func (c *CreditTransferTransactionInformation14) SetChargeBearer(value string) {
	c.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (c *CreditTransferTransactionInformation14) AddChequeInstruction() *Cheque6 {
	c.ChequeInstruction = new(Cheque6)
	return c.ChequeInstruction
}

func (c *CreditTransferTransactionInformation14) AddUltimateDebtor() *PartyIdentification43 {
	c.UltimateDebtor = new(PartyIdentification43)
	return c.UltimateDebtor
}

func (c *CreditTransferTransactionInformation14) AddIntermediaryAgent1() *BranchAndFinancialInstitutionIdentification5 {
	c.IntermediaryAgent1 = new(BranchAndFinancialInstitutionIdentification5)
	return c.IntermediaryAgent1
}

func (c *CreditTransferTransactionInformation14) AddIntermediaryAgent2() *BranchAndFinancialInstitutionIdentification5 {
	c.IntermediaryAgent2 = new(BranchAndFinancialInstitutionIdentification5)
	return c.IntermediaryAgent2
}

func (c *CreditTransferTransactionInformation14) AddIntermediaryAgent3() *BranchAndFinancialInstitutionIdentification5 {
	c.IntermediaryAgent3 = new(BranchAndFinancialInstitutionIdentification5)
	return c.IntermediaryAgent3
}

func (c *CreditTransferTransactionInformation14) AddCreditorAgent() *BranchAndFinancialInstitutionIdentification5 {
	c.CreditorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return c.CreditorAgent
}

func (c *CreditTransferTransactionInformation14) AddCreditor() *PartyIdentification43 {
	c.Creditor = new(PartyIdentification43)
	return c.Creditor
}

func (c *CreditTransferTransactionInformation14) AddCreditorAccount() *CashAccount16 {
	c.CreditorAccount = new(CashAccount16)
	return c.CreditorAccount
}

func (c *CreditTransferTransactionInformation14) AddUltimateCreditor() *PartyIdentification43 {
	c.UltimateCreditor = new(PartyIdentification43)
	return c.UltimateCreditor
}

func (c *CreditTransferTransactionInformation14) AddInstructionForCreditorAgent() *InstructionForCreditorAgent1 {
	newValue := new(InstructionForCreditorAgent1)
	c.InstructionForCreditorAgent = append(c.InstructionForCreditorAgent, newValue)
	return newValue
}

func (c *CreditTransferTransactionInformation14) AddPurpose() *Purpose2Choice {
	c.Purpose = new(Purpose2Choice)
	return c.Purpose
}

func (c *CreditTransferTransactionInformation14) AddRegulatoryReporting() *RegulatoryReporting3 {
	newValue := new(RegulatoryReporting3)
	c.RegulatoryReporting = append(c.RegulatoryReporting, newValue)
	return newValue
}

func (c *CreditTransferTransactionInformation14) AddTax() *TaxInformation3 {
	c.Tax = new(TaxInformation3)
	return c.Tax
}

func (c *CreditTransferTransactionInformation14) AddRelatedRemittanceInformation() *RemittanceLocation2 {
	newValue := new(RemittanceLocation2)
	c.RelatedRemittanceInformation = append(c.RelatedRemittanceInformation, newValue)
	return newValue
}

func (c *CreditTransferTransactionInformation14) AddRemittanceInformation() *RemittanceInformation6 {
	c.RemittanceInformation = new(RemittanceInformation6)
	return c.RemittanceInformation
}
