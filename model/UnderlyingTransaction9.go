package model

// Identifies the underlying (group of) transaction(s) to which the resolution of investigation applies.
type UnderlyingTransaction9 struct {

	// Provides information on the original cancellation message, to which the resolution refers.
	OriginalGroupInformationAndStatus *OriginalGroupHeader5 `xml:"OrgnlGrpInfAndSts,omitempty"`

	// Provides information on the original (group of) transactions, to which the cancellation status refers.
	OriginalPaymentInformationAndStatus []*OriginalPaymentInstruction10 `xml:"OrgnlPmtInfAndSts,omitempty"`

	// Provides details on the original transactions to which the cancellation request message refers.
	TransactionInformationAndStatus []*PaymentTransaction53 `xml:"TxInfAndSts,omitempty"`
}

func (u *UnderlyingTransaction9) AddOriginalGroupInformationAndStatus() *OriginalGroupHeader5 {
	u.OriginalGroupInformationAndStatus = new(OriginalGroupHeader5)
	return u.OriginalGroupInformationAndStatus
}

func (u *UnderlyingTransaction9) AddOriginalPaymentInformationAndStatus() *OriginalPaymentInstruction10 {
	newValue := new(OriginalPaymentInstruction10)
	u.OriginalPaymentInformationAndStatus = append(u.OriginalPaymentInformationAndStatus, newValue)
	return newValue
}

func (u *UnderlyingTransaction9) AddTransactionInformationAndStatus() *PaymentTransaction53 {
	newValue := new(PaymentTransaction53)
	u.TransactionInformationAndStatus = append(u.TransactionInformationAndStatus, newValue)
	return newValue
}
