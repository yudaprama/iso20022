package model

// Set of elements used to identify the underlying transaction.
type EntryTransaction2 struct {

	// Set of elements used to provide the identification of the underlying transaction.
	References *TransactionReferences2 `xml:"Refs,omitempty"`

	// Set of elements providing detailed information on the original amount.
	//
	// Usage: This component (on transaction level) should be used in case booking is for a single transaction and the original amount is different from the entry amount. It can also be used in case individual original amounts are provided in case of a batch or aggregate booking.
	AmountDetails *AmountAndCurrencyExchange3 `xml:"AmtDtls,omitempty"`

	// Set of elements used to indicate when the booked amount of money will become available, that is can be accessed and starts generating interest.
	//
	// Usage: This type of information is used in the US and is linked to particular instruments such as cheques.
	// Example: When a cheque is deposited, it will be booked on the deposit day, but the amount of money will only be accessible as of the indicated availability day (according to national banking regulations).
	Availability []*CashBalanceAvailability2 `xml:"Avlbty,omitempty"`

	// Set of elements used to fully identify the type of underlying transaction resulting in an entry.
	BankTransactionCode *BankTransactionCodeStructure4 `xml:"BkTxCd,omitempty"`

	// Provides information on the charges included in the entry amount.
	//
	// Usage: This component (on transaction level) can be used in case the booking is for a single transaction, and charges are included in the entry amount. It can also be used in case individual charge amounts are applied to individual transactions in case of a batch or aggregate amount booking.
	Charges []*ChargesInformation6 `xml:"Chrgs,omitempty"`

	// Set of elements used to provide details of the interest amount included in the entry amount.
	//
	// Usage: This component (on transaction level) can be used if the booking is for a single transaction, and interest amount is included in the entry amount.  It can also be used if individual interest amounts are applied to individual transactions in the case of a batch or aggregate amount booking.
	Interest []*TransactionInterest2 `xml:"Intrst,omitempty"`

	// Set of elements used to identify the parties related to the underlying transaction.
	RelatedParties *TransactionParty2 `xml:"RltdPties,omitempty"`

	// Set of elements used to identify the agents related to the underlying transaction.
	RelatedAgents *TransactionAgents2 `xml:"RltdAgts,omitempty"`

	// Underlying reason for the payment transaction.
	// Usage: Purpose is used by the end-customers, that is initiating party, (ultimate) debtor, (ultimate) creditor to provide information concerning the nature of the payment. Purpose is a content element, which is not used for processing by any of the agents involved in the payment chain.
	Purpose *Purpose2Choice `xml:"Purp,omitempty"`

	// Set of elements used to provide information related to the handling of the remittance information by any of the agents in the transaction processing chain.
	RelatedRemittanceInformation []*RemittanceLocation2 `xml:"RltdRmtInf,omitempty"`

	// Structured information that enables the matching, ie, reconciliation, of a payment with the items that the payment is intended to settle, such as commercial invoices in an account receivable system.
	RemittanceInformation *RemittanceInformation5 `xml:"RmtInf,omitempty"`

	// Set of elements used to identify the dates related to the underlying transactions.
	RelatedDates *TransactionDates2 `xml:"RltdDts,omitempty"`

	// Set of elements used to identify the price information related to the underlying transaction.
	RelatedPrice *TransactionPrice2Choice `xml:"RltdPric,omitempty"`

	// Set of elements used to identify the related quantities, such as securities, in the underlying transaction.
	RelatedQuantities []*TransactionQuantities1Choice `xml:"RltdQties,omitempty"`

	// Identification of a security, as assigned under a formal or proprietary identification scheme.
	FinancialInstrumentIdentification *SecurityIdentification4Choice `xml:"FinInstrmId,omitempty"`

	// Set of elements used to provide details on the tax.
	Tax *TaxInformation3 `xml:"Tax,omitempty"`

	// Set of elements used to provide the return information.
	ReturnInformation *ReturnReasonInformation10 `xml:"RtrInf,omitempty"`

	// Set of elements used to identify the underlying corporate action.
	CorporateAction *CorporateAction1 `xml:"CorpActn,omitempty"`

	// Safekeeping or investment account. A safekeeping account is an account on which a securities entry is made. An investment account is an account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
	SafekeepingAccount *CashAccount16 `xml:"SfkpgAcct,omitempty"`

	// Further details of the transaction.
	AdditionalTransactionInformation *Max500Text `xml:"AddtlTxInf,omitempty"`
}

func (e *EntryTransaction2) AddReferences() *TransactionReferences2 {
	e.References = new(TransactionReferences2)
	return e.References
}

func (e *EntryTransaction2) AddAmountDetails() *AmountAndCurrencyExchange3 {
	e.AmountDetails = new(AmountAndCurrencyExchange3)
	return e.AmountDetails
}

func (e *EntryTransaction2) AddAvailability() *CashBalanceAvailability2 {
	newValue := new(CashBalanceAvailability2)
	e.Availability = append(e.Availability, newValue)
	return newValue
}

func (e *EntryTransaction2) AddBankTransactionCode() *BankTransactionCodeStructure4 {
	e.BankTransactionCode = new(BankTransactionCodeStructure4)
	return e.BankTransactionCode
}

func (e *EntryTransaction2) AddCharges() *ChargesInformation6 {
	newValue := new(ChargesInformation6)
	e.Charges = append(e.Charges, newValue)
	return newValue
}

func (e *EntryTransaction2) AddInterest() *TransactionInterest2 {
	newValue := new(TransactionInterest2)
	e.Interest = append(e.Interest, newValue)
	return newValue
}

func (e *EntryTransaction2) AddRelatedParties() *TransactionParty2 {
	e.RelatedParties = new(TransactionParty2)
	return e.RelatedParties
}

func (e *EntryTransaction2) AddRelatedAgents() *TransactionAgents2 {
	e.RelatedAgents = new(TransactionAgents2)
	return e.RelatedAgents
}

func (e *EntryTransaction2) AddPurpose() *Purpose2Choice {
	e.Purpose = new(Purpose2Choice)
	return e.Purpose
}

func (e *EntryTransaction2) AddRelatedRemittanceInformation() *RemittanceLocation2 {
	newValue := new(RemittanceLocation2)
	e.RelatedRemittanceInformation = append(e.RelatedRemittanceInformation, newValue)
	return newValue
}

func (e *EntryTransaction2) AddRemittanceInformation() *RemittanceInformation5 {
	e.RemittanceInformation = new(RemittanceInformation5)
	return e.RemittanceInformation
}

func (e *EntryTransaction2) AddRelatedDates() *TransactionDates2 {
	e.RelatedDates = new(TransactionDates2)
	return e.RelatedDates
}

func (e *EntryTransaction2) AddRelatedPrice() *TransactionPrice2Choice {
	e.RelatedPrice = new(TransactionPrice2Choice)
	return e.RelatedPrice
}

func (e *EntryTransaction2) AddRelatedQuantities() *TransactionQuantities1Choice {
	newValue := new(TransactionQuantities1Choice)
	e.RelatedQuantities = append(e.RelatedQuantities, newValue)
	return newValue
}

func (e *EntryTransaction2) AddFinancialInstrumentIdentification() *SecurityIdentification4Choice {
	e.FinancialInstrumentIdentification = new(SecurityIdentification4Choice)
	return e.FinancialInstrumentIdentification
}

func (e *EntryTransaction2) AddTax() *TaxInformation3 {
	e.Tax = new(TaxInformation3)
	return e.Tax
}

func (e *EntryTransaction2) AddReturnInformation() *ReturnReasonInformation10 {
	e.ReturnInformation = new(ReturnReasonInformation10)
	return e.ReturnInformation
}

func (e *EntryTransaction2) AddCorporateAction() *CorporateAction1 {
	e.CorporateAction = new(CorporateAction1)
	return e.CorporateAction
}

func (e *EntryTransaction2) AddSafekeepingAccount() *CashAccount16 {
	e.SafekeepingAccount = new(CashAccount16)
	return e.SafekeepingAccount
}

func (e *EntryTransaction2) SetAdditionalTransactionInformation(value string) {
	e.AdditionalTransactionInformation = (*Max500Text)(&value)
}
