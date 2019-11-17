package model

// Withdrawal transaction for which the completion is sent.
type ATMTransaction3 struct {

	// Identification of the transaction assigned by the ATM.
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Outcome of the financial transaction for the customer.
	TransactionStatus *ATMTransactionStatus1Code `xml:"TxSts"`

	// Multi bundle dispense.
	MultiBundle *TrueFalseIndicator `xml:"MultiBndl,omitempty"`

	// Amount per bundle in the currency of the total presented amount, in the order of the presentation.
	BundlePresentedAmount []*ImpliedCurrencyAndAmount `xml:"BndlPresntdAmt,omitempty"`

	// Status of the amount presented to the customer in the last bundle.
	PresentedAmountStatus *ATMTransactionStatus2Code `xml:"PresntdAmtSts"`

	// Incident occurring during the processing of the transaction.
	Incident []*FailureReason4Code `xml:"Incdnt,omitempty"`

	// Explanation of the incident.
	IncidentDetail []*Max70Text `xml:"IncdntDtl,omitempty"`

	// Identification of the reconciliation period assigned by the ATM.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Unprotected account information.
	AccountData *CardAccount5 `xml:"AcctData,omitempty"`

	// Encryption of account information.
	ProtectedAccountData *ContentInformationType10 `xml:"PrtctdAcctData,omitempty"`

	// Total amount presented to the customer.
	TotalPresentedAmount *AmountAndCurrency1 `xml:"TtlPresntdAmt"`

	// Total authorised amount.
	TotalAuthorisedAmount *ImpliedCurrencyAndAmount `xml:"TtlAuthrsdAmt,omitempty"`

	// Total requested amount.
	TotalRequestedAmount *ImpliedCurrencyAndAmount `xml:"TtlReqdAmt,omitempty"`

	// Detail of the requested amounts for the withdrawal transaction.
	DetailedRequestedAmount *DetailedAmount12 `xml:"DtldReqdAmt,omitempty"`

	// Currency conversion accepted by the customer, either to convert the amount to dispense in the base currency of the ATM, or to convert the total requested amount in the currency of the customer (so called dynamic currency conversion).
	CurrencyConversion *CurrencyConversion4 `xml:"CcyConvs,omitempty"`

	// Additional charge (for instance tax or fee).
	AdditionalCharge []*DetailedAmount13 `xml:"AddtlChrg,omitempty"`

	// True if the customer has requested a receipt.
	RequestedReceipt *TrueFalseIndicator `xml:"ReqdRct,omitempty"`

	// True if a receipt has been printed and presented to the customer.
	ReceiptPrinted *TrueFalseIndicator `xml:"RctPrtd,omitempty"`

	// True when the card was captured by the ATM.
	CapturedCard *TrueFalseIndicator `xml:"CaptrdCard,omitempty"`

	// Limit of amounts for the customer.
	Limits *ATMTransactionAmounts2 `xml:"Lmts,omitempty"`

	// Outcome of the withdrawal authorisation.
	AuthorisationResult *AuthorisationResult9 `xml:"AuthstnRslt,omitempty"`

	// Sequence of one or more TLV data elements from the ATM application, in accordance with ISO 7816-6, not in a specific order. Present if the transaction is performed with an EMV chip card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`

	// Current totals of the ATM.
	ATMTotals []*ATMTotals1 `xml:"ATMTtls,omitempty"`

	// Information on the cassettes of the ATM.
	Cassette []*ATMCassette1 `xml:"Csstt,omitempty"`
}

func (a *ATMTransaction3) AddTransactionIdentification() *TransactionIdentifier1 {
	a.TransactionIdentification = new(TransactionIdentifier1)
	return a.TransactionIdentification
}

func (a *ATMTransaction3) SetTransactionStatus(value string) {
	a.TransactionStatus = (*ATMTransactionStatus1Code)(&value)
}

func (a *ATMTransaction3) SetMultiBundle(value string) {
	a.MultiBundle = (*TrueFalseIndicator)(&value)
}

func (a *ATMTransaction3) AddBundlePresentedAmount(value, currency string) {
	a.BundlePresentedAmount = append(a.BundlePresentedAmount, NewImpliedCurrencyAndAmount(value, currency))
}

func (a *ATMTransaction3) SetPresentedAmountStatus(value string) {
	a.PresentedAmountStatus = (*ATMTransactionStatus2Code)(&value)
}

func (a *ATMTransaction3) AddIncident(value string) {
	a.Incident = append(a.Incident, (*FailureReason4Code)(&value))
}

func (a *ATMTransaction3) AddIncidentDetail(value string) {
	a.IncidentDetail = append(a.IncidentDetail, (*Max70Text)(&value))
}

func (a *ATMTransaction3) SetReconciliationIdentification(value string) {
	a.ReconciliationIdentification = (*Max35Text)(&value)
}

func (a *ATMTransaction3) AddAccountData() *CardAccount5 {
	a.AccountData = new(CardAccount5)
	return a.AccountData
}

func (a *ATMTransaction3) AddProtectedAccountData() *ContentInformationType10 {
	a.ProtectedAccountData = new(ContentInformationType10)
	return a.ProtectedAccountData
}

func (a *ATMTransaction3) AddTotalPresentedAmount() *AmountAndCurrency1 {
	a.TotalPresentedAmount = new(AmountAndCurrency1)
	return a.TotalPresentedAmount
}

func (a *ATMTransaction3) SetTotalAuthorisedAmount(value, currency string) {
	a.TotalAuthorisedAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (a *ATMTransaction3) SetTotalRequestedAmount(value, currency string) {
	a.TotalRequestedAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (a *ATMTransaction3) AddDetailedRequestedAmount() *DetailedAmount12 {
	a.DetailedRequestedAmount = new(DetailedAmount12)
	return a.DetailedRequestedAmount
}

func (a *ATMTransaction3) AddCurrencyConversion() *CurrencyConversion4 {
	a.CurrencyConversion = new(CurrencyConversion4)
	return a.CurrencyConversion
}

func (a *ATMTransaction3) AddAdditionalCharge() *DetailedAmount13 {
	newValue := new(DetailedAmount13)
	a.AdditionalCharge = append(a.AdditionalCharge, newValue)
	return newValue
}

func (a *ATMTransaction3) SetRequestedReceipt(value string) {
	a.RequestedReceipt = (*TrueFalseIndicator)(&value)
}

func (a *ATMTransaction3) SetReceiptPrinted(value string) {
	a.ReceiptPrinted = (*TrueFalseIndicator)(&value)
}

func (a *ATMTransaction3) SetCapturedCard(value string) {
	a.CapturedCard = (*TrueFalseIndicator)(&value)
}

func (a *ATMTransaction3) AddLimits() *ATMTransactionAmounts2 {
	a.Limits = new(ATMTransactionAmounts2)
	return a.Limits
}

func (a *ATMTransaction3) AddAuthorisationResult() *AuthorisationResult9 {
	a.AuthorisationResult = new(AuthorisationResult9)
	return a.AuthorisationResult
}

func (a *ATMTransaction3) SetICCRelatedData(value string) {
	a.ICCRelatedData = (*Max10000Binary)(&value)
}

func (a *ATMTransaction3) AddATMTotals() *ATMTotals1 {
	newValue := new(ATMTotals1)
	a.ATMTotals = append(a.ATMTotals, newValue)
	return newValue
}

func (a *ATMTransaction3) AddCassette() *ATMCassette1 {
	newValue := new(ATMCassette1)
	a.Cassette = append(a.Cassette, newValue)
	return newValue
}
