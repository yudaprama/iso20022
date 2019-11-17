package model

// Identifies the underlying transaction.
type EntryTransaction4 struct {

	// Provides the identification of the underlying transaction.
	References *TransactionReferences3 `xml:"Refs,omitempty"`

	// Amount of money in the cash transaction.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Indicates whether the transaction is a credit or a debit transaction.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Provides detailed information on the original amount.
	//
	// Usage: This component (on transaction level) should be used in case booking is for a single transaction and the original amount is different from the entry amount. It can also be used in case individual original amounts are provided in case of a batch or aggregate booking.
	AmountDetails *AmountAndCurrencyExchange3 `xml:"AmtDtls,omitempty"`

	// Indicates when the booked amount of money will become available, that is can be accessed and starts generating interest.
	//
	// Usage: This type of information is used in the US and is linked to particular instruments such as cheques.
	// Example: When a cheque is deposited, it will be booked on the deposit day, but the amount of money will only be accessible as of the indicated availability day (according to national banking regulations).
	Availability []*CashBalanceAvailability2 `xml:"Avlbty,omitempty"`

	// Set of elements used to fully identify the type of underlying transaction resulting in an entry.
	BankTransactionCode *BankTransactionCodeStructure4 `xml:"BkTxCd,omitempty"`

	// Provides information on the charges, pre-advised or included in the entry amount.
	//
	// Usage: This component (on transaction level) can be used in case the booking is for a single transaction, and charges are included in the entry amount. It can also be used in case individual charge amounts are applied to individual transactions in case of a batch or aggregate amount booking.
	Charges *Charges4 `xml:"Chrgs,omitempty"`

	// Provides details of the interest amount included in the entry amount.
	//
	// Usage: This component (on transaction level) can be used if the booking is for a single transaction, and interest amount is included in the entry amount.  It can also be used if individual interest amounts are applied to individual transactions in the case of a batch or aggregate amount booking.
	Interest *TransactionInterest3 `xml:"Intrst,omitempty"`

	// Set of elements used to identify the parties related to the underlying transaction.
	RelatedParties *TransactionParties3 `xml:"RltdPties,omitempty"`

	// Set of elements used to identify the agents related to the underlying transaction.
	RelatedAgents *TransactionAgents3 `xml:"RltdAgts,omitempty"`

	// Underlying reason for the payment transaction.
	// Usage: Purpose is used by the end-customers, that is initiating party, (ultimate) debtor, (ultimate) creditor to provide information concerning the nature of the payment. Purpose is a content element, which is not used for processing by any of the agents involved in the payment chain.
	Purpose *Purpose2Choice `xml:"Purp,omitempty"`

	// Provides information related to the handling of the remittance information by any of the agents in the transaction processing chain.
	RelatedRemittanceInformation []*RemittanceLocation2 `xml:"RltdRmtInf,omitempty"`

	// Structured information that enables the matching, that is reconciliation, of a payment with the items that the payment is intended to settle, such as commercial invoices in an account receivable system.
	RemittanceInformation *RemittanceInformation7 `xml:"RmtInf,omitempty"`

	// Set of elements used to identify the dates related to the underlying transactions.
	RelatedDates *TransactionDates2 `xml:"RltdDts,omitempty"`

	// Set of elements used to identify the price information related to the underlying transaction.
	RelatedPrice *TransactionPrice3Choice `xml:"RltdPric,omitempty"`

	// Set of elements used to identify the related quantities, such as securities, in the underlying transaction.
	RelatedQuantities []*TransactionQuantities2Choice `xml:"RltdQties,omitempty"`

	// Identification of a security, as assigned under a formal or proprietary identification scheme.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId,omitempty"`

	// Provides details on the tax.
	Tax *TaxInformation3 `xml:"Tax,omitempty"`

	// Provides the return information.
	ReturnInformation *PaymentReturnReason2 `xml:"RtrInf,omitempty"`

	// Set of elements used to identify the underlying corporate action.
	CorporateAction *CorporateAction9 `xml:"CorpActn,omitempty"`

	// Safekeeping or investment account. A safekeeping account is an account on which a securities entry is made. An investment account is an account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
	SafekeepingAccount *SecuritiesAccount13 `xml:"SfkpgAcct,omitempty"`

	// Provides the details of a cash deposit for an amount of money in cash notes and/or coins.
	CashDeposit []*CashDeposit1 `xml:"CshDpst,omitempty"`

	// Provides the data related to the card (number, scheme), terminal (number, identification) and transactional data used to uniquely identify a card transaction.
	CardTransaction *CardTransaction1 `xml:"CardTx,omitempty"`

	// Further details of the transaction.
	AdditionalTransactionInformation *Max500Text `xml:"AddtlTxInf,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (e *EntryTransaction4) AddReferences() *TransactionReferences3 {
	e.References = new(TransactionReferences3)
	return e.References
}

func (e *EntryTransaction4) SetAmount(value, currency string) {
	e.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (e *EntryTransaction4) SetCreditDebitIndicator(value string) {
	e.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (e *EntryTransaction4) AddAmountDetails() *AmountAndCurrencyExchange3 {
	e.AmountDetails = new(AmountAndCurrencyExchange3)
	return e.AmountDetails
}

func (e *EntryTransaction4) AddAvailability() *CashBalanceAvailability2 {
	newValue := new(CashBalanceAvailability2)
	e.Availability = append(e.Availability, newValue)
	return newValue
}

func (e *EntryTransaction4) AddBankTransactionCode() *BankTransactionCodeStructure4 {
	e.BankTransactionCode = new(BankTransactionCodeStructure4)
	return e.BankTransactionCode
}

func (e *EntryTransaction4) AddCharges() *Charges4 {
	e.Charges = new(Charges4)
	return e.Charges
}

func (e *EntryTransaction4) AddInterest() *TransactionInterest3 {
	e.Interest = new(TransactionInterest3)
	return e.Interest
}

func (e *EntryTransaction4) AddRelatedParties() *TransactionParties3 {
	e.RelatedParties = new(TransactionParties3)
	return e.RelatedParties
}

func (e *EntryTransaction4) AddRelatedAgents() *TransactionAgents3 {
	e.RelatedAgents = new(TransactionAgents3)
	return e.RelatedAgents
}

func (e *EntryTransaction4) AddPurpose() *Purpose2Choice {
	e.Purpose = new(Purpose2Choice)
	return e.Purpose
}

func (e *EntryTransaction4) AddRelatedRemittanceInformation() *RemittanceLocation2 {
	newValue := new(RemittanceLocation2)
	e.RelatedRemittanceInformation = append(e.RelatedRemittanceInformation, newValue)
	return newValue
}

func (e *EntryTransaction4) AddRemittanceInformation() *RemittanceInformation7 {
	e.RemittanceInformation = new(RemittanceInformation7)
	return e.RemittanceInformation
}

func (e *EntryTransaction4) AddRelatedDates() *TransactionDates2 {
	e.RelatedDates = new(TransactionDates2)
	return e.RelatedDates
}

func (e *EntryTransaction4) AddRelatedPrice() *TransactionPrice3Choice {
	e.RelatedPrice = new(TransactionPrice3Choice)
	return e.RelatedPrice
}

func (e *EntryTransaction4) AddRelatedQuantities() *TransactionQuantities2Choice {
	newValue := new(TransactionQuantities2Choice)
	e.RelatedQuantities = append(e.RelatedQuantities, newValue)
	return newValue
}

func (e *EntryTransaction4) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	e.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return e.FinancialInstrumentIdentification
}

func (e *EntryTransaction4) AddTax() *TaxInformation3 {
	e.Tax = new(TaxInformation3)
	return e.Tax
}

func (e *EntryTransaction4) AddReturnInformation() *PaymentReturnReason2 {
	e.ReturnInformation = new(PaymentReturnReason2)
	return e.ReturnInformation
}

func (e *EntryTransaction4) AddCorporateAction() *CorporateAction9 {
	e.CorporateAction = new(CorporateAction9)
	return e.CorporateAction
}

func (e *EntryTransaction4) AddSafekeepingAccount() *SecuritiesAccount13 {
	e.SafekeepingAccount = new(SecuritiesAccount13)
	return e.SafekeepingAccount
}

func (e *EntryTransaction4) AddCashDeposit() *CashDeposit1 {
	newValue := new(CashDeposit1)
	e.CashDeposit = append(e.CashDeposit, newValue)
	return newValue
}

func (e *EntryTransaction4) AddCardTransaction() *CardTransaction1 {
	e.CardTransaction = new(CardTransaction1)
	return e.CardTransaction
}

func (e *EntryTransaction4) SetAdditionalTransactionInformation(value string) {
	e.AdditionalTransactionInformation = (*Max500Text)(&value)
}

func (e *EntryTransaction4) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	e.SupplementaryData = append(e.SupplementaryData, newValue)
	return newValue
}
