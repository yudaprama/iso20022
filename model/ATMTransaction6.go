package model

// Inquiry information for the transaction.
type ATMTransaction6 struct {

	// Identification of the transaction assigned by the ATM.
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Unprotected account information.
	AccountData *CardAccount3 `xml:"AcctData,omitempty"`

	// Encryption of account information.
	ProtectedAccountData *ContentInformationType10 `xml:"PrtctdAcctData,omitempty"`

	// Amount to be authorised by the issuer.
	TotalRequestedAmount *AmountAndCurrency1 `xml:"TtlReqdAmt,omitempty"`

	// Amounts of the withdrawal transaction.
	DetailedRequestedAmount *DetailedAmount12 `xml:"DtldReqdAmt,omitempty"`

	// Sequence of one or more TLV data elements from the ATM application, in accordance with ISO 7816-6, not in a specific order. Present if the transaction is performed with an EMV chip card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`
}

func (a *ATMTransaction6) AddTransactionIdentification() *TransactionIdentifier1 {
	a.TransactionIdentification = new(TransactionIdentifier1)
	return a.TransactionIdentification
}

func (a *ATMTransaction6) AddAccountData() *CardAccount3 {
	a.AccountData = new(CardAccount3)
	return a.AccountData
}

func (a *ATMTransaction6) AddProtectedAccountData() *ContentInformationType10 {
	a.ProtectedAccountData = new(ContentInformationType10)
	return a.ProtectedAccountData
}

func (a *ATMTransaction6) AddTotalRequestedAmount() *AmountAndCurrency1 {
	a.TotalRequestedAmount = new(AmountAndCurrency1)
	return a.TotalRequestedAmount
}

func (a *ATMTransaction6) AddDetailedRequestedAmount() *DetailedAmount12 {
	a.DetailedRequestedAmount = new(DetailedAmount12)
	return a.DetailedRequestedAmount
}

func (a *ATMTransaction6) SetICCRelatedData(value string) {
	a.ICCRelatedData = (*Max10000Binary)(&value)
}
