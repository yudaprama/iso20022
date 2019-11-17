package model

// Provides further details specific to the individual direct debit transaction(s) included in the message.
type DirectDebitTransactionInformation15 struct {

	// References used for  a payment instruction.
	PaymentIdentification *PaymentIdentification3 `xml:"PmtId"`

	// Specifies the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation21 `xml:"PmtTpInf,omitempty"`

	// Amount of money moved between the instructing agent and the instructed agent.
	InterbankSettlementAmount *ActiveCurrencyAndAmount `xml:"IntrBkSttlmAmt"`

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	InterbankSettlementDate *ISODate `xml:"IntrBkSttlmDt,omitempty"`

	// Indicator of the urgency or order of importance that the instructing party would like the instructed party to apply to the processing of the settlement instruction.
	SettlementPriority *Priority3Code `xml:"SttlmPrty,omitempty"`

	// Provides information on the requested settlement time(s) of the payment instruction.
	SettlementTimeRequest *SettlementTimeRequest2 `xml:"SttlmTmReq,omitempty"`

	// Ultimate financial institution that owes an amount of money to the (ultimate) institutional creditor.
	UltimateDebtor *BranchAndFinancialInstitutionIdentification5 `xml:"UltmtDbtr,omitempty"`

	// Financial institution that owes an amount of money to the (ultimate) financial institutional creditor.
	Debtor *BranchAndFinancialInstitutionIdentification5 `xml:"Dbtr"`

	// Unambiguous identification of the account of the debtor to which a debit entry will be made as a result of the transaction.
	DebtorAccount *CashAccount24 `xml:"DbtrAcct,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"DbtrAgt,omitempty"`

	// Unambiguous identification of the account of the debtor agent at its servicing agent in the payment chain.
	DebtorAgentAccount *CashAccount24 `xml:"DbtrAgtAcct,omitempty"`

	// Further information related to the processing of the payment instruction, that may need to be acted upon by the debtor agent, depending on agreement between debtor and the debtor agent.
	//
	InstructionForDebtorAgent *Max210Text `xml:"InstrForDbtrAgt,omitempty"`

	// Information supplied to enable the matching of an entry with the items that the transfer is intended to settle, such as commercial invoices in an accounts' receivable system.
	RemittanceInformation *RemittanceInformation2 `xml:"RmtInf,omitempty"`
}

func (d *DirectDebitTransactionInformation15) AddPaymentIdentification() *PaymentIdentification3 {
	d.PaymentIdentification = new(PaymentIdentification3)
	return d.PaymentIdentification
}

func (d *DirectDebitTransactionInformation15) AddPaymentTypeInformation() *PaymentTypeInformation21 {
	d.PaymentTypeInformation = new(PaymentTypeInformation21)
	return d.PaymentTypeInformation
}

func (d *DirectDebitTransactionInformation15) SetInterbankSettlementAmount(value, currency string) {
	d.InterbankSettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (d *DirectDebitTransactionInformation15) SetInterbankSettlementDate(value string) {
	d.InterbankSettlementDate = (*ISODate)(&value)
}

func (d *DirectDebitTransactionInformation15) SetSettlementPriority(value string) {
	d.SettlementPriority = (*Priority3Code)(&value)
}

func (d *DirectDebitTransactionInformation15) AddSettlementTimeRequest() *SettlementTimeRequest2 {
	d.SettlementTimeRequest = new(SettlementTimeRequest2)
	return d.SettlementTimeRequest
}

func (d *DirectDebitTransactionInformation15) AddUltimateDebtor() *BranchAndFinancialInstitutionIdentification5 {
	d.UltimateDebtor = new(BranchAndFinancialInstitutionIdentification5)
	return d.UltimateDebtor
}

func (d *DirectDebitTransactionInformation15) AddDebtor() *BranchAndFinancialInstitutionIdentification5 {
	d.Debtor = new(BranchAndFinancialInstitutionIdentification5)
	return d.Debtor
}

func (d *DirectDebitTransactionInformation15) AddDebtorAccount() *CashAccount24 {
	d.DebtorAccount = new(CashAccount24)
	return d.DebtorAccount
}

func (d *DirectDebitTransactionInformation15) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification5 {
	d.DebtorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return d.DebtorAgent
}

func (d *DirectDebitTransactionInformation15) AddDebtorAgentAccount() *CashAccount24 {
	d.DebtorAgentAccount = new(CashAccount24)
	return d.DebtorAgentAccount
}

func (d *DirectDebitTransactionInformation15) SetInstructionForDebtorAgent(value string) {
	d.InstructionForDebtorAgent = (*Max210Text)(&value)
}

func (d *DirectDebitTransactionInformation15) AddRemittanceInformation() *RemittanceInformation2 {
	d.RemittanceInformation = new(RemittanceInformation2)
	return d.RemittanceInformation
}
