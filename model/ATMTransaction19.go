package model

// Response to the deposit request.
type ATMTransaction19 struct {

	// Identification of the transaction assigned by the ATM.
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Outcome of the financial transaction for the customer.
	TransactionStatus *ATMTransactionStatus1Code `xml:"TxSts"`

	// Identification of the reconciliation period assigned by the ATM.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Incident occurring during the processing of the transaction.
	Incident []*FailureReason7Code `xml:"Incdnt,omitempty"`

	// Explanation of the incident.
	IncidentDetail []*Max70Text `xml:"IncdntDtl,omitempty"`

	// Unprotected account information.
	AccountData *CardAccount8 `xml:"AcctData,omitempty"`

	// Encryption of account information.
	ProtectedAccountData *ContentInformationType10 `xml:"PrtctdAcctData,omitempty"`

	// Total amount deposed by the customer.
	TotalDepositedAmount *AmountAndCurrency1 `xml:"TtlDpstdAmt"`

	// Total authorised amount of the deposit transaction.
	TotalAuthorisedAmount *ImpliedCurrencyAndAmount `xml:"TtlAuthrsdAmt"`

	// Total requested amount.
	TotalRequestedAmount *ImpliedCurrencyAndAmount `xml:"TtlReqdAmt,omitempty"`

	// Amounts of the deposit transaction.
	DetailedRequestedAmount *DetailedAmount16 `xml:"DtldReqdAmt,omitempty"`

	// Additional charge (for instance tax or fee).
	AdditionalCharge []*DetailedAmount13 `xml:"AddtlChrg,omitempty"`

	// True if the customer has requested a receipt.
	RequestedReceipt *TrueFalseIndicator `xml:"ReqdRct,omitempty"`

	// True if a receipt has been printed and presented to the customer.
	ReceiptPrinted *TrueFalseIndicator `xml:"RctPrtd,omitempty"`

	// Outcome of the deposit authorisation.
	AuthorisationResult *AuthorisationResult13 `xml:"AuthstnRslt,omitempty"`

	// Deposited media put in the safe.
	DepositedMedia []*ATMDepositedMedia1 `xml:"DpstdMdia,omitempty"`

	// Media unit not put in the safe. These deposits have to be reconciliated.
	ToBeReconciledMediaCounters []*ATMDepositedMedia3 `xml:"ToBeRcncldMdiaCntrs,omitempty"`

	// Current totals of the ATM.
	ATMTotals []*ATMTotals1 `xml:"ATMTtls,omitempty"`

	// Information on the cassettes of the ATM.
	Cassette []*ATMCassette2 `xml:"Csstt,omitempty"`

	// Sequence of one or more TLV data elements from the ATM application, in accordance with ISO 7816-6, not in a specific order. Present if the transaction is performed with an EMV chip card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`
}

func (a *ATMTransaction19) AddTransactionIdentification() *TransactionIdentifier1 {
	a.TransactionIdentification = new(TransactionIdentifier1)
	return a.TransactionIdentification
}

func (a *ATMTransaction19) SetTransactionStatus(value string) {
	a.TransactionStatus = (*ATMTransactionStatus1Code)(&value)
}

func (a *ATMTransaction19) SetReconciliationIdentification(value string) {
	a.ReconciliationIdentification = (*Max35Text)(&value)
}

func (a *ATMTransaction19) AddIncident(value string) {
	a.Incident = append(a.Incident, (*FailureReason7Code)(&value))
}

func (a *ATMTransaction19) AddIncidentDetail(value string) {
	a.IncidentDetail = append(a.IncidentDetail, (*Max70Text)(&value))
}

func (a *ATMTransaction19) AddAccountData() *CardAccount8 {
	a.AccountData = new(CardAccount8)
	return a.AccountData
}

func (a *ATMTransaction19) AddProtectedAccountData() *ContentInformationType10 {
	a.ProtectedAccountData = new(ContentInformationType10)
	return a.ProtectedAccountData
}

func (a *ATMTransaction19) AddTotalDepositedAmount() *AmountAndCurrency1 {
	a.TotalDepositedAmount = new(AmountAndCurrency1)
	return a.TotalDepositedAmount
}

func (a *ATMTransaction19) SetTotalAuthorisedAmount(value, currency string) {
	a.TotalAuthorisedAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (a *ATMTransaction19) SetTotalRequestedAmount(value, currency string) {
	a.TotalRequestedAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (a *ATMTransaction19) AddDetailedRequestedAmount() *DetailedAmount16 {
	a.DetailedRequestedAmount = new(DetailedAmount16)
	return a.DetailedRequestedAmount
}

func (a *ATMTransaction19) AddAdditionalCharge() *DetailedAmount13 {
	newValue := new(DetailedAmount13)
	a.AdditionalCharge = append(a.AdditionalCharge, newValue)
	return newValue
}

func (a *ATMTransaction19) SetRequestedReceipt(value string) {
	a.RequestedReceipt = (*TrueFalseIndicator)(&value)
}

func (a *ATMTransaction19) SetReceiptPrinted(value string) {
	a.ReceiptPrinted = (*TrueFalseIndicator)(&value)
}

func (a *ATMTransaction19) AddAuthorisationResult() *AuthorisationResult13 {
	a.AuthorisationResult = new(AuthorisationResult13)
	return a.AuthorisationResult
}

func (a *ATMTransaction19) AddDepositedMedia() *ATMDepositedMedia1 {
	newValue := new(ATMDepositedMedia1)
	a.DepositedMedia = append(a.DepositedMedia, newValue)
	return newValue
}

func (a *ATMTransaction19) AddToBeReconciledMediaCounters() *ATMDepositedMedia3 {
	newValue := new(ATMDepositedMedia3)
	a.ToBeReconciledMediaCounters = append(a.ToBeReconciledMediaCounters, newValue)
	return newValue
}

func (a *ATMTransaction19) AddATMTotals() *ATMTotals1 {
	newValue := new(ATMTotals1)
	a.ATMTotals = append(a.ATMTotals, newValue)
	return newValue
}

func (a *ATMTransaction19) AddCassette() *ATMCassette2 {
	newValue := new(ATMCassette2)
	a.Cassette = append(a.Cassette, newValue)
	return newValue
}

func (a *ATMTransaction19) SetICCRelatedData(value string) {
	a.ICCRelatedData = (*Max10000Binary)(&value)
}
