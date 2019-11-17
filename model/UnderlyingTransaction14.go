package model

// Identifies the underlying (group of) transaction(s) to which the resolution of investigation applies.
type UnderlyingTransaction14 struct {

	// Provides information on the original cancellation message, to which the resolution refers.
	OriginalGroupInformationAndStatus *OriginalGroupHeader5 `xml:"OrgnlGrpInfAndSts,omitempty"`

	// Provides information on the original (group of) transactions, to which the cancellation status refers.
	OriginalPaymentInformationAndStatus []*OriginalPaymentInstruction17 `xml:"OrgnlPmtInfAndSts,omitempty"`

	// Provides details on the original transactions to which the cancellation request message refers.
	TransactionInformationAndStatus []*PaymentTransaction67 `xml:"TxInfAndSts,omitempty"`
}

func (u *UnderlyingTransaction14) AddOriginalGroupInformationAndStatus() *OriginalGroupHeader5 {
	u.OriginalGroupInformationAndStatus = new(OriginalGroupHeader5)
	return u.OriginalGroupInformationAndStatus
}

func (u *UnderlyingTransaction14) AddOriginalPaymentInformationAndStatus() *OriginalPaymentInstruction17 {
	newValue := new(OriginalPaymentInstruction17)
	u.OriginalPaymentInformationAndStatus = append(u.OriginalPaymentInformationAndStatus, newValue)
	return newValue
}

func (u *UnderlyingTransaction14) AddTransactionInformationAndStatus() *PaymentTransaction67 {
	newValue := new(PaymentTransaction67)
	u.TransactionInformationAndStatus = append(u.TransactionInformationAndStatus, newValue)
	return newValue
}
