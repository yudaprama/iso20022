package model

// Identifies the underlying (group of) transaction(s) to which the investigation applies.
type UnderlyingTransaction15 struct {

	// Provides information on the original message, to which the cancellation refers.
	OriginalGroupInformationAndCancellation *OriginalGroupHeader6 `xml:"OrgnlGrpInfAndCxl,omitempty"`

	// Provides information on the original (group of) transactions, to which the cancellation request refers.
	OriginalPaymentInformationAndCancellation []*OriginalPaymentInstruction20 `xml:"OrgnlPmtInfAndCxl,omitempty"`
}

func (u *UnderlyingTransaction15) AddOriginalGroupInformationAndCancellation() *OriginalGroupHeader6 {
	u.OriginalGroupInformationAndCancellation = new(OriginalGroupHeader6)
	return u.OriginalGroupInformationAndCancellation
}

func (u *UnderlyingTransaction15) AddOriginalPaymentInformationAndCancellation() *OriginalPaymentInstruction20 {
	newValue := new(OriginalPaymentInstruction20)
	u.OriginalPaymentInformationAndCancellation = append(u.OriginalPaymentInformationAndCancellation, newValue)
	return newValue
}
