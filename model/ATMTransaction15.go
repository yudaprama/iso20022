package model

// Deposit transaction for which the service is requested.
type ATMTransaction15 struct {

	// Identification of the transaction assigned by the ATM.
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the reconciliation period assigned by the ATM.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Unprotected account information.
	AccountData *CardAccount9 `xml:"AcctData,omitempty"`

	// Encryption of account information.
	ProtectedAccountData *ContentInformationType10 `xml:"PrtctdAcctData,omitempty"`

	// Total amount of the deposit transaction.
	TotalAmount *AmountAndCurrency1 `xml:"TtlAmt,omitempty"`

	// Amounts of the deposit transaction.
	DetailedRequestedAmount []*DetailedAmount16 `xml:"DtldReqdAmt,omitempty"`

	// Deposited media put in the safe.
	DepositedMedia []*ATMDepositedMedia1 `xml:"DpstdMdia,omitempty"`

	// True if a receipt has be requested by the customer.
	RequestedReceipt *TrueFalseIndicator `xml:"ReqdRct,omitempty"`

	// Sequence of one or more TLV data elements from the ATM application, in accordance with ISO 7816-6, not in a specific order. Present if the transaction is performed with an EMV chip card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`
}

func (a *ATMTransaction15) AddTransactionIdentification() *TransactionIdentifier1 {
	a.TransactionIdentification = new(TransactionIdentifier1)
	return a.TransactionIdentification
}

func (a *ATMTransaction15) SetReconciliationIdentification(value string) {
	a.ReconciliationIdentification = (*Max35Text)(&value)
}

func (a *ATMTransaction15) AddAccountData() *CardAccount9 {
	a.AccountData = new(CardAccount9)
	return a.AccountData
}

func (a *ATMTransaction15) AddProtectedAccountData() *ContentInformationType10 {
	a.ProtectedAccountData = new(ContentInformationType10)
	return a.ProtectedAccountData
}

func (a *ATMTransaction15) AddTotalAmount() *AmountAndCurrency1 {
	a.TotalAmount = new(AmountAndCurrency1)
	return a.TotalAmount
}

func (a *ATMTransaction15) AddDetailedRequestedAmount() *DetailedAmount16 {
	newValue := new(DetailedAmount16)
	a.DetailedRequestedAmount = append(a.DetailedRequestedAmount, newValue)
	return newValue
}

func (a *ATMTransaction15) AddDepositedMedia() *ATMDepositedMedia1 {
	newValue := new(ATMDepositedMedia1)
	a.DepositedMedia = append(a.DepositedMedia, newValue)
	return newValue
}

func (a *ATMTransaction15) SetRequestedReceipt(value string) {
	a.RequestedReceipt = (*TrueFalseIndicator)(&value)
}

func (a *ATMTransaction15) SetICCRelatedData(value string) {
	a.ICCRelatedData = (*Max10000Binary)(&value)
}
