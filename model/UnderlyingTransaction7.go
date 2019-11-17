package model

// Identifies the underlying (group of) transaction(s) to which the investigation applies.
type UnderlyingTransaction7 struct {

	// Provides information on the original message, to which the cancellation refers.
	OriginalGroupInformationAndCancellation *OriginalGroupHeader4 `xml:"OrgnlGrpInfAndCxl,omitempty"`

	// Provides information on the original (group of) transactions, to which the cancellation request refers.
	OriginalPaymentInformationAndCancellation []*OriginalPaymentInstruction8 `xml:"OrgnlPmtInfAndCxl,omitempty"`
}

func (u *UnderlyingTransaction7) AddOriginalGroupInformationAndCancellation() *OriginalGroupHeader4 {
	u.OriginalGroupInformationAndCancellation = new(OriginalGroupHeader4)
	return u.OriginalGroupInformationAndCancellation
}

func (u *UnderlyingTransaction7) AddOriginalPaymentInformationAndCancellation() *OriginalPaymentInstruction8 {
	newValue := new(OriginalPaymentInstruction8)
	u.OriginalPaymentInformationAndCancellation = append(u.OriginalPaymentInformationAndCancellation, newValue)
	return newValue
}
