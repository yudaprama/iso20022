package model

// Identifies the underlying (group of) transaction(s) to which the investigation applies.
type UnderlyingTransaction11 struct {

	// Provides information on the original message, to which the cancellation refers.
	OriginalGroupInformationAndCancellation *OriginalGroupHeader4 `xml:"OrgnlGrpInfAndCxl,omitempty"`

	// Provides information on the original (group of) transactions, to which the cancellation request refers.
	OriginalPaymentInformationAndCancellation []*OriginalPaymentInstruction13 `xml:"OrgnlPmtInfAndCxl,omitempty"`
}

func (u *UnderlyingTransaction11) AddOriginalGroupInformationAndCancellation() *OriginalGroupHeader4 {
	u.OriginalGroupInformationAndCancellation = new(OriginalGroupHeader4)
	return u.OriginalGroupInformationAndCancellation
}

func (u *UnderlyingTransaction11) AddOriginalPaymentInformationAndCancellation() *OriginalPaymentInstruction13 {
	newValue := new(OriginalPaymentInstruction13)
	u.OriginalPaymentInformationAndCancellation = append(u.OriginalPaymentInformationAndCancellation, newValue)
	return newValue
}
