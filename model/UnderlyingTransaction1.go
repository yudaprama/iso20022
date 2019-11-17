package model

// Set of elements used to identify the underlying (group of) transaction(s) to which the investigation applies.
type UnderlyingTransaction1 struct {

	// Set of elements used to provide information on the original messsage, to which the cancellation refers.
	OriginalGroupInformationAndCancellation *OriginalGroupInformation23 `xml:"OrgnlGrpInfAndCxl,omitempty"`

	// Set of elements used to provide information on the original (group of) transactions, to which the cancellation request refers.
	OriginalPaymentInformationAndCancellation []*OriginalPaymentInformation4 `xml:"OrgnlPmtInfAndCxl,omitempty"`
}

func (u *UnderlyingTransaction1) AddOriginalGroupInformationAndCancellation() *OriginalGroupInformation23 {
	u.OriginalGroupInformationAndCancellation = new(OriginalGroupInformation23)
	return u.OriginalGroupInformationAndCancellation
}

func (u *UnderlyingTransaction1) AddOriginalPaymentInformationAndCancellation() *OriginalPaymentInformation4 {
	newValue := new(OriginalPaymentInformation4)
	u.OriginalPaymentInformationAndCancellation = append(u.OriginalPaymentInformationAndCancellation, newValue)
	return newValue
}
