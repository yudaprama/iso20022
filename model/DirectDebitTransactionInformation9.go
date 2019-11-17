package model

// Set of elements used to provide information specific to the individual direct debit transaction(s) included in the message.
type DirectDebitTransactionInformation9 struct {

	// Set of elements used to reference a payment instruction.
	PaymentIdentification *PaymentIdentification1 `xml:"PmtId"`

	// Set of elements used to further specify the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation20 `xml:"PmtTpInf,omitempty"`

	// Amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as ordered by the initiating party.
	InstructedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"InstdAmt"`

	// Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr,omitempty"`

	// Set of elements providing information specific to the direct debit mandate.
	DirectDebitTransaction *DirectDebitTransaction6 `xml:"DrctDbtTx,omitempty"`

	// Ultimate party to which an amount of money is due.
	UltimateCreditor *PartyIdentification32 `xml:"UltmtCdtr,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification4 `xml:"DbtrAgt"`

	// Unambiguous identification of the account of the debtor agent at its servicing agent in the payment chain.
	DebtorAgentAccount *CashAccount16 `xml:"DbtrAgtAcct,omitempty"`

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentification32 `xml:"Dbtr"`

	// Unambiguous identification of the account of the debtor to which a debit entry will be made as a result of the transaction.
	DebtorAccount *CashAccount16 `xml:"DbtrAcct"`

	// Ultimate party that owes an amount of money to the (ultimate) creditor.
	UltimateDebtor *PartyIdentification32 `xml:"UltmtDbtr,omitempty"`

	// Further information, related to the processing of the payment instruction, that may need to be acted upon by the creditor agent, depending on agreement between creditor and the creditor agent.
	InstructionForCreditorAgent *Max140Text `xml:"InstrForCdtrAgt,omitempty"`

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

func (d *DirectDebitTransactionInformation9) AddPaymentIdentification() *PaymentIdentification1 {
	d.PaymentIdentification = new(PaymentIdentification1)
	return d.PaymentIdentification
}

func (d *DirectDebitTransactionInformation9) AddPaymentTypeInformation() *PaymentTypeInformation20 {
	d.PaymentTypeInformation = new(PaymentTypeInformation20)
	return d.PaymentTypeInformation
}

func (d *DirectDebitTransactionInformation9) SetInstructedAmount(value, currency string) {
	d.InstructedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (d *DirectDebitTransactionInformation9) SetChargeBearer(value string) {
	d.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (d *DirectDebitTransactionInformation9) AddDirectDebitTransaction() *DirectDebitTransaction6 {
	d.DirectDebitTransaction = new(DirectDebitTransaction6)
	return d.DirectDebitTransaction
}

func (d *DirectDebitTransactionInformation9) AddUltimateCreditor() *PartyIdentification32 {
	d.UltimateCreditor = new(PartyIdentification32)
	return d.UltimateCreditor
}

func (d *DirectDebitTransactionInformation9) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification4 {
	d.DebtorAgent = new(BranchAndFinancialInstitutionIdentification4)
	return d.DebtorAgent
}

func (d *DirectDebitTransactionInformation9) AddDebtorAgentAccount() *CashAccount16 {
	d.DebtorAgentAccount = new(CashAccount16)
	return d.DebtorAgentAccount
}

func (d *DirectDebitTransactionInformation9) AddDebtor() *PartyIdentification32 {
	d.Debtor = new(PartyIdentification32)
	return d.Debtor
}

func (d *DirectDebitTransactionInformation9) AddDebtorAccount() *CashAccount16 {
	d.DebtorAccount = new(CashAccount16)
	return d.DebtorAccount
}

func (d *DirectDebitTransactionInformation9) AddUltimateDebtor() *PartyIdentification32 {
	d.UltimateDebtor = new(PartyIdentification32)
	return d.UltimateDebtor
}

func (d *DirectDebitTransactionInformation9) SetInstructionForCreditorAgent(value string) {
	d.InstructionForCreditorAgent = (*Max140Text)(&value)
}

func (d *DirectDebitTransactionInformation9) AddPurpose() *Purpose2Choice {
	d.Purpose = new(Purpose2Choice)
	return d.Purpose
}

func (d *DirectDebitTransactionInformation9) AddRegulatoryReporting() *RegulatoryReporting3 {
	newValue := new(RegulatoryReporting3)
	d.RegulatoryReporting = append(d.RegulatoryReporting, newValue)
	return newValue
}

func (d *DirectDebitTransactionInformation9) AddTax() *TaxInformation3 {
	d.Tax = new(TaxInformation3)
	return d.Tax
}

func (d *DirectDebitTransactionInformation9) AddRelatedRemittanceInformation() *RemittanceLocation2 {
	newValue := new(RemittanceLocation2)
	d.RelatedRemittanceInformation = append(d.RelatedRemittanceInformation, newValue)
	return newValue
}

func (d *DirectDebitTransactionInformation9) AddRemittanceInformation() *RemittanceInformation5 {
	d.RemittanceInformation = new(RemittanceInformation5)
	return d.RemittanceInformation
}
