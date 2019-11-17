package model

// Transfer information for the transaction.
type ATMTransaction23 struct {

	// Identification of the transaction assigned by the ATM.
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the reconciliation period assigned by the ATM.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Description of the transfer for the creditor.
	CreditorLabel *Max35Text `xml:"CdtrLabl,omitempty"`

	// Description of the transfer for the debtor.
	DebtorLabel *Max35Text `xml:"DbtrLabl,omitempty"`

	// Reference of the payment.
	PaymentReference *Max35Text `xml:"PmtRef,omitempty"`

	// Information about the source account of the transfer.
	AccountFrom *CardAccount7 `xml:"AcctFr,omitempty"`

	// Encryption of the source account information.
	ProtectedAccountFrom *ContentInformationType10 `xml:"PrtctdAcctFr,omitempty"`

	// Information about the destination account of the transfer.
	AccountTo []*CardAccount7 `xml:"AcctTo,omitempty"`

	// Encryption of the destination account information.
	ProtectedAccountTo *ContentInformationType10 `xml:"PrtctdAcctTo,omitempty"`

	// Amount of the transaction to be authorised.
	TotalRequestedAmount *ImpliedCurrencyAndAmount `xml:"TtlReqdAmt,omitempty"`

	// Details of the transfer transaction amounts.
	DetailedRequestedAmount *DetailedAmount17 `xml:"DtldReqdAmt,omitempty"`

	// Requested date of the execution of the transfer.
	RequestedExecutionDate *ISODate `xml:"ReqdExctnDt,omitempty"`

	// Identifies the instant transfer program.
	InstantTransferProgram *Max35Text `xml:"InstntTrfPrgm,omitempty"`

	// Information for reccurring transfer or standing orders.
	RecurringTransfer *RecurringTransaction3 `xml:"RcrngTrf,omitempty"`

	// True if the customer has requested a receipt.
	RequestedReceipt *TrueFalseIndicator `xml:"ReqdRct,omitempty"`

	// Sequence of one or more TLV data elements from the ATM application, in accordance with ISO 7816-6, not in a specific order. Present if the transaction is performed with an EMV chip card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`
}

func (a *ATMTransaction23) AddTransactionIdentification() *TransactionIdentifier1 {
	a.TransactionIdentification = new(TransactionIdentifier1)
	return a.TransactionIdentification
}

func (a *ATMTransaction23) SetReconciliationIdentification(value string) {
	a.ReconciliationIdentification = (*Max35Text)(&value)
}

func (a *ATMTransaction23) SetCreditorLabel(value string) {
	a.CreditorLabel = (*Max35Text)(&value)
}

func (a *ATMTransaction23) SetDebtorLabel(value string) {
	a.DebtorLabel = (*Max35Text)(&value)
}

func (a *ATMTransaction23) SetPaymentReference(value string) {
	a.PaymentReference = (*Max35Text)(&value)
}

func (a *ATMTransaction23) AddAccountFrom() *CardAccount7 {
	a.AccountFrom = new(CardAccount7)
	return a.AccountFrom
}

func (a *ATMTransaction23) AddProtectedAccountFrom() *ContentInformationType10 {
	a.ProtectedAccountFrom = new(ContentInformationType10)
	return a.ProtectedAccountFrom
}

func (a *ATMTransaction23) AddAccountTo() *CardAccount7 {
	newValue := new(CardAccount7)
	a.AccountTo = append(a.AccountTo, newValue)
	return newValue
}

func (a *ATMTransaction23) AddProtectedAccountTo() *ContentInformationType10 {
	a.ProtectedAccountTo = new(ContentInformationType10)
	return a.ProtectedAccountTo
}

func (a *ATMTransaction23) SetTotalRequestedAmount(value, currency string) {
	a.TotalRequestedAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (a *ATMTransaction23) AddDetailedRequestedAmount() *DetailedAmount17 {
	a.DetailedRequestedAmount = new(DetailedAmount17)
	return a.DetailedRequestedAmount
}

func (a *ATMTransaction23) SetRequestedExecutionDate(value string) {
	a.RequestedExecutionDate = (*ISODate)(&value)
}

func (a *ATMTransaction23) SetInstantTransferProgram(value string) {
	a.InstantTransferProgram = (*Max35Text)(&value)
}

func (a *ATMTransaction23) AddRecurringTransfer() *RecurringTransaction3 {
	a.RecurringTransfer = new(RecurringTransaction3)
	return a.RecurringTransfer
}

func (a *ATMTransaction23) SetRequestedReceipt(value string) {
	a.RequestedReceipt = (*TrueFalseIndicator)(&value)
}

func (a *ATMTransaction23) SetICCRelatedData(value string) {
	a.ICCRelatedData = (*Max10000Binary)(&value)
}
