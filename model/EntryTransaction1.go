package model

// Posting to an account that results in an increase or decrease to a balance.
type EntryTransaction1 struct {

	// Set of elements providing the identification of the underlying transaction.
	References *TransactionReferences1 `xml:"Refs,omitempty"`

	// Set of elements providing details information on the original amount.
	//
	// Usage: This component (on transaction level) should be used in case booking is for a single transaction and the original amount is different from the entry amount.  It can also be used in case individual original amounts are provided in case of a batch or aggregate booking.
	AmountDetails *AmountAndCurrencyExchange2 `xml:"AmtDtls,omitempty"`

	// Set of elements used to indicate when the booked funds will become available, ie can be accessed and start generating interest.
	//
	// Usage : this type of info is eg used in US, and is linked to particular instruments, such as cheques.
	// Example : When a cheque is deposited, it will be booked on the deposit day, but the funds will only be accessible as of the indicated availability day (according to national banking regulations).
	Availability []*CashBalanceAvailability1 `xml:"Avlbty,omitempty"`

	// Set of elements to fully identify the type of underlying transaction resulting in an entry.
	BankTransactionCode *BankTransactionCodeStructure1 `xml:"BkTxCd,omitempty"`

	// Provides information on the charges included in the entry amount.
	//
	// Usage : This component (on transaction level) can be used in case the booking is for a single transaction, and charges are included in the entry amount.  It can also be used in case individual charge amounts are applied to individual transactions in case of a batch or aggregate amount booking.
	Charges []*ChargesInformation3 `xml:"Chrgs,omitempty"`

	// Set of elements providing details on the interest amount included in the entry amount.
	//
	// Usage : This component (on transaction level) can be used in case the booking is for a single transaction, and interest amount is included in the entry amount.  It can also be used in case individual interest amounts are applied to individual transactions in case of a batch or aggregate amount booking.
	Interest []*TransactionInterest1 `xml:"Intrst,omitempty"`

	// Set of elements identifying the parties related to the underlying transaction.
	RelatedParties *TransactionParty1 `xml:"RltdPties,omitempty"`

	// Set of elements identifying the agents related to the underlying transaction.
	RelatedAgents *TransactionAgents1 `xml:"RltdAgts,omitempty"`

	// Underlying reason for the payment transaction, eg, a charity payment, or a commercial agreement between the creditor and the debtor.
	//
	// Usage: purpose is used by the end-customers, ie originating party, initiating party, debtor, creditor, final party, to provide information concerning the nature of the payment transaction. Purpose is a content element, which is not used for processing by any of the agents involved in the payment chain.
	Purpose *Purpose1Choice `xml:"Purp,omitempty"`

	// Information related to the handling of the remittance information by any of the agents in the transaction processing chain.
	RelatedRemittanceInformation []*RemittanceLocation1 `xml:"RltdRmtInf,omitempty"`

	// Information that enables the matching, ie, reconciliation, of a payment with the items that the payment is intended to settle, eg, commercial invoices in an account receivable system.
	RemittanceInformation *RemittanceInformation1 `xml:"RmtInf,omitempty"`

	// Set of elements identifying the dates related to the underlying transactions.
	RelatedDates *TransactionDates1 `xml:"RltdDts,omitempty"`

	// Set of elements identifying the price information related to the underlying transaction.
	RelatedPrice *TransactionPrice1Choice `xml:"RltdPric,omitempty"`

	// Identifies related quantities (eg of securities) in the underlying transaction.
	RelatedQuantities []*TransactionQuantities1Choice `xml:"RltdQties,omitempty"`

	// Identification of a security, as assigned under a formal or proprietary identification scheme.
	FinancialInstrumentIdentification *SecurityIdentification4Choice `xml:"FinInstrmId,omitempty"`

	// Amount of money due to the government or tax authority, according to various pre-defined parameters such as thresholds or income.
	Tax *TaxInformation2 `xml:"Tax,omitempty"`

	// Set of elements specifying the return information.
	ReturnInformation *ReturnReasonInformation5 `xml:"RtrInf,omitempty"`

	// Set of elements identifying the underlying corporate action.
	CorporateAction *CorporateAction1 `xml:"CorpActn,omitempty"`

	// Safekeeping or investment account. A safekeeping account is an account on which a securities entry is made. An investment account is an account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
	SafekeepingAccount *CashAccount7 `xml:"SfkpgAcct,omitempty"`

	// Further details on the transaction details.
	AdditionalTransactionInformation *Max500Text `xml:"AddtlTxInf,omitempty"`
}

func (e *EntryTransaction1) AddReferences() *TransactionReferences1 {
	e.References = new(TransactionReferences1)
	return e.References
}

func (e *EntryTransaction1) AddAmountDetails() *AmountAndCurrencyExchange2 {
	e.AmountDetails = new(AmountAndCurrencyExchange2)
	return e.AmountDetails
}

func (e *EntryTransaction1) AddAvailability() *CashBalanceAvailability1 {
	newValue := new(CashBalanceAvailability1)
	e.Availability = append(e.Availability, newValue)
	return newValue
}

func (e *EntryTransaction1) AddBankTransactionCode() *BankTransactionCodeStructure1 {
	e.BankTransactionCode = new(BankTransactionCodeStructure1)
	return e.BankTransactionCode
}

func (e *EntryTransaction1) AddCharges() *ChargesInformation3 {
	newValue := new(ChargesInformation3)
	e.Charges = append(e.Charges, newValue)
	return newValue
}

func (e *EntryTransaction1) AddInterest() *TransactionInterest1 {
	newValue := new(TransactionInterest1)
	e.Interest = append(e.Interest, newValue)
	return newValue
}

func (e *EntryTransaction1) AddRelatedParties() *TransactionParty1 {
	e.RelatedParties = new(TransactionParty1)
	return e.RelatedParties
}

func (e *EntryTransaction1) AddRelatedAgents() *TransactionAgents1 {
	e.RelatedAgents = new(TransactionAgents1)
	return e.RelatedAgents
}

func (e *EntryTransaction1) AddPurpose() *Purpose1Choice {
	e.Purpose = new(Purpose1Choice)
	return e.Purpose
}

func (e *EntryTransaction1) AddRelatedRemittanceInformation() *RemittanceLocation1 {
	newValue := new(RemittanceLocation1)
	e.RelatedRemittanceInformation = append(e.RelatedRemittanceInformation, newValue)
	return newValue
}

func (e *EntryTransaction1) AddRemittanceInformation() *RemittanceInformation1 {
	e.RemittanceInformation = new(RemittanceInformation1)
	return e.RemittanceInformation
}

func (e *EntryTransaction1) AddRelatedDates() *TransactionDates1 {
	e.RelatedDates = new(TransactionDates1)
	return e.RelatedDates
}

func (e *EntryTransaction1) AddRelatedPrice() *TransactionPrice1Choice {
	e.RelatedPrice = new(TransactionPrice1Choice)
	return e.RelatedPrice
}

func (e *EntryTransaction1) AddRelatedQuantities() *TransactionQuantities1Choice {
	newValue := new(TransactionQuantities1Choice)
	e.RelatedQuantities = append(e.RelatedQuantities, newValue)
	return newValue
}

func (e *EntryTransaction1) AddFinancialInstrumentIdentification() *SecurityIdentification4Choice {
	e.FinancialInstrumentIdentification = new(SecurityIdentification4Choice)
	return e.FinancialInstrumentIdentification
}

func (e *EntryTransaction1) AddTax() *TaxInformation2 {
	e.Tax = new(TaxInformation2)
	return e.Tax
}

func (e *EntryTransaction1) AddReturnInformation() *ReturnReasonInformation5 {
	e.ReturnInformation = new(ReturnReasonInformation5)
	return e.ReturnInformation
}

func (e *EntryTransaction1) AddCorporateAction() *CorporateAction1 {
	e.CorporateAction = new(CorporateAction1)
	return e.CorporateAction
}

func (e *EntryTransaction1) AddSafekeepingAccount() *CashAccount7 {
	e.SafekeepingAccount = new(CashAccount7)
	return e.SafekeepingAccount
}

func (e *EntryTransaction1) SetAdditionalTransactionInformation(value string) {
	e.AdditionalTransactionInformation = (*Max500Text)(&value)
}
