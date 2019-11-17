package model

// Transfer information for the transaction.
type ATMTransaction24 struct {

	// Identification of the transaction assigned by the ATM.
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the reconciliation period assigned by the ATM.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Description of the transfer for the creditor.
	CreditorLabel *Max35Text `xml:"CdtrLabl,omitempty"`

	// Description of the transfer for the debtor.
	DebtorLabel *Max35Text `xml:"DbtrLabl,omitempty"`

	// Identifier of the approved transfer transaction for the bank.
	TransferIdentifier *Max70Text `xml:"TrfIdr,omitempty"`

	// Reference of the payment.
	PaymentReference *Max35Text `xml:"PmtRef,omitempty"`

	// Result of the fund transfer request.
	TransactionResponse *ResponseType7 `xml:"TxRspn"`

	// Sequence of actions to be performed by the ATM to complete the transaction.
	Action []*Action7 `xml:"Actn,omitempty"`

	// Information about the source account of the transfer.
	AccountFrom *CardAccount13 `xml:"AcctFr,omitempty"`

	// Encryption of the source account information.
	ProtectedAccountFrom *ContentInformationType10 `xml:"PrtctdAcctFr,omitempty"`

	// Information about the destination account of the transfer.
	AccountTo []*CardAccount13 `xml:"AcctTo,omitempty"`

	// Encryption of the destination account information.
	ProtectedAccountTo *ContentInformationType10 `xml:"PrtctdAcctTo,omitempty"`

	// Total authorised amount.
	TotalAuthorisedAmount *AmountAndCurrency1 `xml:"TtlAuthrsdAmt"`

	// Total requested amount.
	TotalRequestedAmount *ImpliedCurrencyAndAmount `xml:"TtlReqdAmt,omitempty"`

	// Details of the transfer transaction amounts.
	DetailedRequestedAmount *DetailedAmount17 `xml:"DtldReqdAmt,omitempty"`

	// Additional charge (for instance tax or fee).
	AdditionalCharge []*DetailedAmount18 `xml:"AddtlChrg,omitempty"`

	// Limit of amounts for the customer.
	Limits *ATMTransactionAmounts6 `xml:"Lmts,omitempty"`

	// Requested date of the execution of the transfer.
	RequestedExecutionDate *ISODate `xml:"ReqdExctnDt,omitempty"`

	// Proposed date of the execution of the transfer.
	ProposedExecutionDate *ISODate `xml:"PropsdExctnDt,omitempty"`

	// Identifies the instant transfer program.
	InstantTransferProgram *Max35Text `xml:"InstntTrfPrgm,omitempty"`

	// Information for reccurring transfer or standing orders.
	RecurringTransfer *RecurringTransaction3 `xml:"RcrngTrf,omitempty"`

	// Outcome of the transfer authorisation.
	AuthorisationResult *AuthorisationResult13 `xml:"AuthstnRslt,omitempty"`

	// Sequence of one or more TLV data elements from the ATM application, in accordance with ISO 7816-6, not in a specific order. Present if the transaction is performed with an EMV chip card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`

	// Maintenance command to perform on the ATM.
	Command []*ATMCommand7 `xml:"Cmd,omitempty"`
}

func (a *ATMTransaction24) AddTransactionIdentification() *TransactionIdentifier1 {
	a.TransactionIdentification = new(TransactionIdentifier1)
	return a.TransactionIdentification
}

func (a *ATMTransaction24) SetReconciliationIdentification(value string) {
	a.ReconciliationIdentification = (*Max35Text)(&value)
}

func (a *ATMTransaction24) SetCreditorLabel(value string) {
	a.CreditorLabel = (*Max35Text)(&value)
}

func (a *ATMTransaction24) SetDebtorLabel(value string) {
	a.DebtorLabel = (*Max35Text)(&value)
}

func (a *ATMTransaction24) SetTransferIdentifier(value string) {
	a.TransferIdentifier = (*Max70Text)(&value)
}

func (a *ATMTransaction24) SetPaymentReference(value string) {
	a.PaymentReference = (*Max35Text)(&value)
}

func (a *ATMTransaction24) AddTransactionResponse() *ResponseType7 {
	a.TransactionResponse = new(ResponseType7)
	return a.TransactionResponse
}

func (a *ATMTransaction24) AddAction() *Action7 {
	newValue := new(Action7)
	a.Action = append(a.Action, newValue)
	return newValue
}

func (a *ATMTransaction24) AddAccountFrom() *CardAccount13 {
	a.AccountFrom = new(CardAccount13)
	return a.AccountFrom
}

func (a *ATMTransaction24) AddProtectedAccountFrom() *ContentInformationType10 {
	a.ProtectedAccountFrom = new(ContentInformationType10)
	return a.ProtectedAccountFrom
}

func (a *ATMTransaction24) AddAccountTo() *CardAccount13 {
	newValue := new(CardAccount13)
	a.AccountTo = append(a.AccountTo, newValue)
	return newValue
}

func (a *ATMTransaction24) AddProtectedAccountTo() *ContentInformationType10 {
	a.ProtectedAccountTo = new(ContentInformationType10)
	return a.ProtectedAccountTo
}

func (a *ATMTransaction24) AddTotalAuthorisedAmount() *AmountAndCurrency1 {
	a.TotalAuthorisedAmount = new(AmountAndCurrency1)
	return a.TotalAuthorisedAmount
}

func (a *ATMTransaction24) SetTotalRequestedAmount(value, currency string) {
	a.TotalRequestedAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (a *ATMTransaction24) AddDetailedRequestedAmount() *DetailedAmount17 {
	a.DetailedRequestedAmount = new(DetailedAmount17)
	return a.DetailedRequestedAmount
}

func (a *ATMTransaction24) AddAdditionalCharge() *DetailedAmount18 {
	newValue := new(DetailedAmount18)
	a.AdditionalCharge = append(a.AdditionalCharge, newValue)
	return newValue
}

func (a *ATMTransaction24) AddLimits() *ATMTransactionAmounts6 {
	a.Limits = new(ATMTransactionAmounts6)
	return a.Limits
}

func (a *ATMTransaction24) SetRequestedExecutionDate(value string) {
	a.RequestedExecutionDate = (*ISODate)(&value)
}

func (a *ATMTransaction24) SetProposedExecutionDate(value string) {
	a.ProposedExecutionDate = (*ISODate)(&value)
}

func (a *ATMTransaction24) SetInstantTransferProgram(value string) {
	a.InstantTransferProgram = (*Max35Text)(&value)
}

func (a *ATMTransaction24) AddRecurringTransfer() *RecurringTransaction3 {
	a.RecurringTransfer = new(RecurringTransaction3)
	return a.RecurringTransfer
}

func (a *ATMTransaction24) AddAuthorisationResult() *AuthorisationResult13 {
	a.AuthorisationResult = new(AuthorisationResult13)
	return a.AuthorisationResult
}

func (a *ATMTransaction24) SetICCRelatedData(value string) {
	a.ICCRelatedData = (*Max10000Binary)(&value)
}

func (a *ATMTransaction24) AddCommand() *ATMCommand7 {
	newValue := new(ATMCommand7)
	a.Command = append(a.Command, newValue)
	return newValue
}
