package model

// Identifies the underlying (group of) transaction(s) to which the resolution of investigation applies.
type UnderlyingTransaction17 struct {

	// Provides information on the original cancellation message, to which the resolution refers.
	OriginalGroupInformationAndStatus *OriginalGroupHeader5 `xml:"OrgnlGrpInfAndSts,omitempty"`

	// Provides information on the original (group of) transactions, to which the cancellation status refers.
	OriginalPaymentInformationAndStatus []*OriginalPaymentInstruction22 `xml:"OrgnlPmtInfAndSts,omitempty"`

	// Provides details on the original transactions to which the cancellation request message refers.
	TransactionInformationAndStatus []*PaymentTransaction79 `xml:"TxInfAndSts,omitempty"`
}

func (u *UnderlyingTransaction17) AddOriginalGroupInformationAndStatus() *OriginalGroupHeader5 {
	u.OriginalGroupInformationAndStatus = new(OriginalGroupHeader5)
	return u.OriginalGroupInformationAndStatus
}

func (u *UnderlyingTransaction17) AddOriginalPaymentInformationAndStatus() *OriginalPaymentInstruction22 {
	newValue := new(OriginalPaymentInstruction22)
	u.OriginalPaymentInformationAndStatus = append(u.OriginalPaymentInformationAndStatus, newValue)
	return newValue
}

func (u *UnderlyingTransaction17) AddTransactionInformationAndStatus() *PaymentTransaction79 {
	newValue := new(PaymentTransaction79)
	u.TransactionInformationAndStatus = append(u.TransactionInformationAndStatus, newValue)
	return newValue
}
