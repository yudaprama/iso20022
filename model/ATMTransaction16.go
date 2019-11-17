package model

// Response to the deposit request.
type ATMTransaction16 struct {

	// Identification of the transaction assigned by the ATM.
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the reconciliation period assigned by the ATM.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// True if a completion advice has to be sent after the end of the transaction.
	CompletionRequired *TrueFalseIndicator `xml:"CmpltnReqrd,omitempty"`

	// Unprotected account information.
	AccountData *CardAccount10 `xml:"AcctData,omitempty"`

	// Encryption of account information.
	ProtectedAccountData *ContentInformationType10 `xml:"PrtctdAcctData,omitempty"`

	// Total authorised amount of the deposit transaction.
	TotalAuthorisedAmount *AmountAndCurrency1 `xml:"TtlAuthrsdAmt"`

	// Total requested amount.
	TotalRequestedAmount *ImpliedCurrencyAndAmount `xml:"TtlReqdAmt,omitempty"`

	// Detail of the requested amounts for the deposit transaction.
	DetailedRequestedAmount []*DetailedAmount16 `xml:"DtldReqdAmt,omitempty"`

	// Additional charge (for instance tax or fee).
	AdditionalCharge []*DetailedAmount13 `xml:"AddtlChrg,omitempty"`

	// Outcome of the deposit authorisation.
	AuthorisationResult *AuthorisationResult13 `xml:"AuthstnRslt,omitempty"`

	// Sequence of one or more TLV data elements from the ATM application, in accordance with ISO 7816-6, not in a specific order. Present if the transaction is performed with an EMV chip card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`

	// Maintenance command to perform on the ATM.
	Command []*ATMCommand7 `xml:"Cmd,omitempty"`
}

func (a *ATMTransaction16) AddTransactionIdentification() *TransactionIdentifier1 {
	a.TransactionIdentification = new(TransactionIdentifier1)
	return a.TransactionIdentification
}

func (a *ATMTransaction16) SetReconciliationIdentification(value string) {
	a.ReconciliationIdentification = (*Max35Text)(&value)
}

func (a *ATMTransaction16) SetCompletionRequired(value string) {
	a.CompletionRequired = (*TrueFalseIndicator)(&value)
}

func (a *ATMTransaction16) AddAccountData() *CardAccount10 {
	a.AccountData = new(CardAccount10)
	return a.AccountData
}

func (a *ATMTransaction16) AddProtectedAccountData() *ContentInformationType10 {
	a.ProtectedAccountData = new(ContentInformationType10)
	return a.ProtectedAccountData
}

func (a *ATMTransaction16) AddTotalAuthorisedAmount() *AmountAndCurrency1 {
	a.TotalAuthorisedAmount = new(AmountAndCurrency1)
	return a.TotalAuthorisedAmount
}

func (a *ATMTransaction16) SetTotalRequestedAmount(value, currency string) {
	a.TotalRequestedAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (a *ATMTransaction16) AddDetailedRequestedAmount() *DetailedAmount16 {
	newValue := new(DetailedAmount16)
	a.DetailedRequestedAmount = append(a.DetailedRequestedAmount, newValue)
	return newValue
}

func (a *ATMTransaction16) AddAdditionalCharge() *DetailedAmount13 {
	newValue := new(DetailedAmount13)
	a.AdditionalCharge = append(a.AdditionalCharge, newValue)
	return newValue
}

func (a *ATMTransaction16) AddAuthorisationResult() *AuthorisationResult13 {
	a.AuthorisationResult = new(AuthorisationResult13)
	return a.AuthorisationResult
}

func (a *ATMTransaction16) SetICCRelatedData(value string) {
	a.ICCRelatedData = (*Max10000Binary)(&value)
}

func (a *ATMTransaction16) AddCommand() *ATMCommand7 {
	newValue := new(ATMCommand7)
	a.Command = append(a.Command, newValue)
	return newValue
}
