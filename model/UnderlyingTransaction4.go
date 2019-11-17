package model

// Identifies the underlying (group of) transaction(s) to which the resolution of investigation applies.
type UnderlyingTransaction4 struct {

	// Provides information on the original cancellation message, to which the resolution refers.
	OriginalGroupInformationAndStatus *OriginalGroupHeader5 `xml:"OrgnlGrpInfAndSts,omitempty"`

	// Provides information on the original (group of) transactions, to which the cancellation status refers.
	OriginalPaymentInformationAndStatus []*OriginalPaymentInstruction3 `xml:"OrgnlPmtInfAndSts,omitempty"`

	// Provides details on the original transactions to which the cancellation request message refers.
	TransactionInformationAndStatus []*PaymentTransaction40 `xml:"TxInfAndSts,omitempty"`
}

func (u *UnderlyingTransaction4) AddOriginalGroupInformationAndStatus() *OriginalGroupHeader5 {
	u.OriginalGroupInformationAndStatus = new(OriginalGroupHeader5)
	return u.OriginalGroupInformationAndStatus
}

func (u *UnderlyingTransaction4) AddOriginalPaymentInformationAndStatus() *OriginalPaymentInstruction3 {
	newValue := new(OriginalPaymentInstruction3)
	u.OriginalPaymentInformationAndStatus = append(u.OriginalPaymentInformationAndStatus, newValue)
	return newValue
}

func (u *UnderlyingTransaction4) AddTransactionInformationAndStatus() *PaymentTransaction40 {
	newValue := new(PaymentTransaction40)
	u.TransactionInformationAndStatus = append(u.TransactionInformationAndStatus, newValue)
	return newValue
}
