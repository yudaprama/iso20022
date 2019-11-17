package model

// Response to the withdrawal transaction request.
type ATMTransaction2 struct {

	// Identification of the transaction assigned by the ATM.
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the reconciliation period assigned by the ATM.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// True if a withdrawal completion advice has to be sent after the end of the transaction.
	CompletionRequired *TrueFalseIndicator `xml:"CmpltnReqrd,omitempty"`

	// Unprotected account information.
	AccountData *CardAccount4 `xml:"AcctData,omitempty"`

	// Encryption of account information.
	ProtectedAccountData *ContentInformationType10 `xml:"PrtctdAcctData,omitempty"`

	// Total authorised amount.
	TotalAuthorisedAmount *AmountAndCurrency1 `xml:"TtlAuthrsdAmt"`

	// Total requested amount.
	TotalRequestedAmount *ImpliedCurrencyAndAmount `xml:"TtlReqdAmt,omitempty"`

	// Detail of the requested amounts for the withdrawal transaction.
	DetailedRequestedAmount *DetailedAmount12 `xml:"DtldReqdAmt,omitempty"`

	// Currency conversion accepted by the customer, either to convert the amount to dispense in the base currency of the ATM, or to convert the total requested amount in the currency of the customer (so called dynamic currency conversion).
	CurrencyConversion *CurrencyConversion4 `xml:"CcyConvs,omitempty"`

	// Additional charge (for instance tax or fee).
	AdditionalCharge []*DetailedAmount13 `xml:"AddtlChrg,omitempty"`

	// Limit of amounts for the customer.
	Limits *ATMTransactionAmounts2 `xml:"Lmts,omitempty"`

	// Media mix algorithm requested by the ATM Host, the identification of the algorithm is an agreement between the ATM and the ATM manager.
	MixType *Max35Text `xml:"MixTp,omitempty"`

	// Media mix selected requested by the ATM Host.
	Mix []*ATMMediaMix1 `xml:"Mix,omitempty"`

	// Outcome of the withdrawal authorisation.
	AuthorisationResult *AuthorisationResult9 `xml:"AuthstnRslt"`

	// Sequence of one or more TLV data elements from the ATM application, in accordance with ISO 7816-6, not in a specific order. Present if the transaction is performed with an EMV chip card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`

	// Maintenance command to perform on the ATM.
	Command []*ATMCommand1 `xml:"Cmd,omitempty"`
}

func (a *ATMTransaction2) AddTransactionIdentification() *TransactionIdentifier1 {
	a.TransactionIdentification = new(TransactionIdentifier1)
	return a.TransactionIdentification
}

func (a *ATMTransaction2) SetReconciliationIdentification(value string) {
	a.ReconciliationIdentification = (*Max35Text)(&value)
}

func (a *ATMTransaction2) SetCompletionRequired(value string) {
	a.CompletionRequired = (*TrueFalseIndicator)(&value)
}

func (a *ATMTransaction2) AddAccountData() *CardAccount4 {
	a.AccountData = new(CardAccount4)
	return a.AccountData
}

func (a *ATMTransaction2) AddProtectedAccountData() *ContentInformationType10 {
	a.ProtectedAccountData = new(ContentInformationType10)
	return a.ProtectedAccountData
}

func (a *ATMTransaction2) AddTotalAuthorisedAmount() *AmountAndCurrency1 {
	a.TotalAuthorisedAmount = new(AmountAndCurrency1)
	return a.TotalAuthorisedAmount
}

func (a *ATMTransaction2) SetTotalRequestedAmount(value, currency string) {
	a.TotalRequestedAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (a *ATMTransaction2) AddDetailedRequestedAmount() *DetailedAmount12 {
	a.DetailedRequestedAmount = new(DetailedAmount12)
	return a.DetailedRequestedAmount
}

func (a *ATMTransaction2) AddCurrencyConversion() *CurrencyConversion4 {
	a.CurrencyConversion = new(CurrencyConversion4)
	return a.CurrencyConversion
}

func (a *ATMTransaction2) AddAdditionalCharge() *DetailedAmount13 {
	newValue := new(DetailedAmount13)
	a.AdditionalCharge = append(a.AdditionalCharge, newValue)
	return newValue
}

func (a *ATMTransaction2) AddLimits() *ATMTransactionAmounts2 {
	a.Limits = new(ATMTransactionAmounts2)
	return a.Limits
}

func (a *ATMTransaction2) SetMixType(value string) {
	a.MixType = (*Max35Text)(&value)
}

func (a *ATMTransaction2) AddMix() *ATMMediaMix1 {
	newValue := new(ATMMediaMix1)
	a.Mix = append(a.Mix, newValue)
	return newValue
}

func (a *ATMTransaction2) AddAuthorisationResult() *AuthorisationResult9 {
	a.AuthorisationResult = new(AuthorisationResult9)
	return a.AuthorisationResult
}

func (a *ATMTransaction2) SetICCRelatedData(value string) {
	a.ICCRelatedData = (*Max10000Binary)(&value)
}

func (a *ATMTransaction2) AddCommand() *ATMCommand1 {
	newValue := new(ATMCommand1)
	a.Command = append(a.Command, newValue)
	return newValue
}
