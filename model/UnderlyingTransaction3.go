package model

// Set of elements used to identify the underlying (group of) transaction(s) to which the resolution of investigation applies.
type UnderlyingTransaction3 struct {

	// Set of elements used to provide information on the original cancellation message, to which the resolution refers.
	OriginalGroupInformationAndStatus *OriginalGroupInformation24 `xml:"OrgnlGrpInfAndSts,omitempty"`

	// Set of elements used to provide information on the original (group of) transactions, to which the cancellation status refers.
	OriginalPaymentInformationAndStatus []*OriginalPaymentInformation3 `xml:"OrgnlPmtInfAndSts,omitempty"`

	// Set of elements used to provide information on the original transactions to which the cancellation request message refers.
	TransactionInformationAndStatus []*PaymentTransactionInformation33 `xml:"TxInfAndSts,omitempty"`
}

func (u *UnderlyingTransaction3) AddOriginalGroupInformationAndStatus() *OriginalGroupInformation24 {
	u.OriginalGroupInformationAndStatus = new(OriginalGroupInformation24)
	return u.OriginalGroupInformationAndStatus
}

func (u *UnderlyingTransaction3) AddOriginalPaymentInformationAndStatus() *OriginalPaymentInformation3 {
	newValue := new(OriginalPaymentInformation3)
	u.OriginalPaymentInformationAndStatus = append(u.OriginalPaymentInformationAndStatus, newValue)
	return newValue
}

func (u *UnderlyingTransaction3) AddTransactionInformationAndStatus() *PaymentTransactionInformation33 {
	newValue := new(PaymentTransactionInformation33)
	u.TransactionInformationAndStatus = append(u.TransactionInformationAndStatus, newValue)
	return newValue
}
