package model

// Withdrawal transaction for which an authorisation is requested.
type ATMTransaction13 struct {

	// True if cash has to be dispensed by the ATM for the transaction.
	CashDispensed *TrueFalseIndicator `xml:"CshDspnsd,omitempty"`

	// Identification of the transaction assigned by the ATM.
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the reconciliation period assigned by the ATM.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Unprotected account information.
	AccountData *CardAccount7 `xml:"AcctData,omitempty"`

	// Encryption of account information.
	ProtectedAccountData *ContentInformationType10 `xml:"PrtctdAcctData,omitempty"`

	// Amount to be authorised by the issuer.
	TotalRequestedAmount *AmountAndCurrency1 `xml:"TtlReqdAmt,omitempty"`

	// Amounts of the withdrawal transaction.
	DetailedRequestedAmount *DetailedAmount12 `xml:"DtldReqdAmt,omitempty"`

	// Currency conversion accepted by the customer, either to convert the amount to dispense in the base currency of the ATM, or to convert the total requested amount in the currency of the customer (so called dynamic currency conversion).
	CurrencyConversionResult *CurrencyConversion9 `xml:"CcyConvsRslt,omitempty"`

	// Media mix algorithm, the identification of the algorithm is an agreement between the ATM and the ATM manager.
	SelectedMixType *Max35Text `xml:"SelctdMixTp,omitempty"`

	// Media mix selected.
	SelectedMix []*ATMMediaMix1 `xml:"SelctdMix,omitempty"`

	// True if a receipt has be requested by the customer.
	RequestedReceipt *TrueFalseIndicator `xml:"ReqdRct,omitempty"`

	// Sequence of one or more TLV data elements from the ATM application, in accordance with ISO 7816-6, not in a specific order. Present if the transaction is performed with an EMV chip card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`
}

func (a *ATMTransaction13) SetCashDispensed(value string) {
	a.CashDispensed = (*TrueFalseIndicator)(&value)
}

func (a *ATMTransaction13) AddTransactionIdentification() *TransactionIdentifier1 {
	a.TransactionIdentification = new(TransactionIdentifier1)
	return a.TransactionIdentification
}

func (a *ATMTransaction13) SetReconciliationIdentification(value string) {
	a.ReconciliationIdentification = (*Max35Text)(&value)
}

func (a *ATMTransaction13) AddAccountData() *CardAccount7 {
	a.AccountData = new(CardAccount7)
	return a.AccountData
}

func (a *ATMTransaction13) AddProtectedAccountData() *ContentInformationType10 {
	a.ProtectedAccountData = new(ContentInformationType10)
	return a.ProtectedAccountData
}

func (a *ATMTransaction13) AddTotalRequestedAmount() *AmountAndCurrency1 {
	a.TotalRequestedAmount = new(AmountAndCurrency1)
	return a.TotalRequestedAmount
}

func (a *ATMTransaction13) AddDetailedRequestedAmount() *DetailedAmount12 {
	a.DetailedRequestedAmount = new(DetailedAmount12)
	return a.DetailedRequestedAmount
}

func (a *ATMTransaction13) AddCurrencyConversionResult() *CurrencyConversion9 {
	a.CurrencyConversionResult = new(CurrencyConversion9)
	return a.CurrencyConversionResult
}

func (a *ATMTransaction13) SetSelectedMixType(value string) {
	a.SelectedMixType = (*Max35Text)(&value)
}

func (a *ATMTransaction13) AddSelectedMix() *ATMMediaMix1 {
	newValue := new(ATMMediaMix1)
	a.SelectedMix = append(a.SelectedMix, newValue)
	return newValue
}

func (a *ATMTransaction13) SetRequestedReceipt(value string) {
	a.RequestedReceipt = (*TrueFalseIndicator)(&value)
}

func (a *ATMTransaction13) SetICCRelatedData(value string) {
	a.ICCRelatedData = (*Max10000Binary)(&value)
}
