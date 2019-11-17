package model

// Set of characteristics that apply to the the direct debit transaction(s).
type DirectDebitTransactionInformation1 struct {

	// Set of elements to reference a payment instruction.
	PaymentIdentification *PaymentIdentification1 `xml:"PmtId"`

	// Set of elements used to further specify the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation2 `xml:"PmtTpInf,omitempty"`

	// Amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as ordered by the initiating party.
	InstructedAmount *CurrencyAndAmount `xml:"InstdAmt"`

	// Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr,omitempty"`

	// Set of elements providing information specific to the direct debit mandate.
	DirectDebitTransaction *DirectDebitTransaction1 `xml:"DrctDbtTx,omitempty"`

	// Ultimate party to which an amount of money is due.
	UltimateCreditor *PartyIdentification8 `xml:"UltmtCdtr,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification3 `xml:"DbtrAgt"`

	// Unambiguous identification of the account of the debtor agent at its servicing agent in the payment chain.
	DebtorAgentAccount *CashAccount7 `xml:"DbtrAgtAcct,omitempty"`

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentification8 `xml:"Dbtr"`

	// Identification of the account of the debtor to which a debit entry will be made to execute the transfer.
	DebtorAccount *CashAccount7 `xml:"DbtrAcct"`

	// Ultimate party that owes an amount of money to the (ultimate) creditor.
	UltimateDebtor *PartyIdentification8 `xml:"UltmtDbtr,omitempty"`

	// Further information, related to the processing of the payment instruction, that may need to be acted upon by the creditor agent, depending on agreement between creditor and the creditor agent.
	InstructionForCreditorAgent *Max140Text `xml:"InstrForCdtrAgt,omitempty"`

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

func (d *DirectDebitTransactionInformation1) AddPaymentIdentification() *PaymentIdentification1 {
	d.PaymentIdentification = new(PaymentIdentification1)
	return d.PaymentIdentification
}

func (d *DirectDebitTransactionInformation1) AddPaymentTypeInformation() *PaymentTypeInformation2 {
	d.PaymentTypeInformation = new(PaymentTypeInformation2)
	return d.PaymentTypeInformation
}

func (d *DirectDebitTransactionInformation1) SetInstructedAmount(value, currency string) {
	d.InstructedAmount = NewCurrencyAndAmount(value, currency)
}

func (d *DirectDebitTransactionInformation1) SetChargeBearer(value string) {
	d.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (d *DirectDebitTransactionInformation1) AddDirectDebitTransaction() *DirectDebitTransaction1 {
	d.DirectDebitTransaction = new(DirectDebitTransaction1)
	return d.DirectDebitTransaction
}

func (d *DirectDebitTransactionInformation1) AddUltimateCreditor() *PartyIdentification8 {
	d.UltimateCreditor = new(PartyIdentification8)
	return d.UltimateCreditor
}

func (d *DirectDebitTransactionInformation1) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification3 {
	d.DebtorAgent = new(BranchAndFinancialInstitutionIdentification3)
	return d.DebtorAgent
}

func (d *DirectDebitTransactionInformation1) AddDebtorAgentAccount() *CashAccount7 {
	d.DebtorAgentAccount = new(CashAccount7)
	return d.DebtorAgentAccount
}

func (d *DirectDebitTransactionInformation1) AddDebtor() *PartyIdentification8 {
	d.Debtor = new(PartyIdentification8)
	return d.Debtor
}

func (d *DirectDebitTransactionInformation1) AddDebtorAccount() *CashAccount7 {
	d.DebtorAccount = new(CashAccount7)
	return d.DebtorAccount
}

func (d *DirectDebitTransactionInformation1) AddUltimateDebtor() *PartyIdentification8 {
	d.UltimateDebtor = new(PartyIdentification8)
	return d.UltimateDebtor
}

func (d *DirectDebitTransactionInformation1) SetInstructionForCreditorAgent(value string) {
	d.InstructionForCreditorAgent = (*Max140Text)(&value)
}

func (d *DirectDebitTransactionInformation1) AddPurpose() *Purpose1Choice {
	d.Purpose = new(Purpose1Choice)
	return d.Purpose
}

func (d *DirectDebitTransactionInformation1) AddRegulatoryReporting() *RegulatoryReporting2 {
	newValue := new(RegulatoryReporting2)
	d.RegulatoryReporting = append(d.RegulatoryReporting, newValue)
	return newValue
}

func (d *DirectDebitTransactionInformation1) AddTax() *TaxInformation2 {
	d.Tax = new(TaxInformation2)
	return d.Tax
}

func (d *DirectDebitTransactionInformation1) AddRelatedRemittanceInformation() *RemittanceLocation1 {
	newValue := new(RemittanceLocation1)
	d.RelatedRemittanceInformation = append(d.RelatedRemittanceInformation, newValue)
	return newValue
}

func (d *DirectDebitTransactionInformation1) AddRemittanceInformation() *RemittanceInformation1 {
	d.RemittanceInformation = new(RemittanceInformation1)
	return d.RemittanceInformation
}
