package model

// Identifies the underlying (group of) transaction(s) to which the investigation applies.
type UnderlyingTransaction16 struct {

	// Provides information on the original message, to which the cancellation refers.
	OriginalGroupInformationAndCancellation *OriginalGroupHeader6 `xml:"OrgnlGrpInfAndCxl,omitempty"`

	// Provides information on the original transactions to which the cancellation request message refers.
	TransactionInformation []*PaymentTransaction75 `xml:"TxInf,omitempty"`
}

func (u *UnderlyingTransaction16) AddOriginalGroupInformationAndCancellation() *OriginalGroupHeader6 {
	u.OriginalGroupInformationAndCancellation = new(OriginalGroupHeader6)
	return u.OriginalGroupInformationAndCancellation
}

func (u *UnderlyingTransaction16) AddTransactionInformation() *PaymentTransaction75 {
	newValue := new(PaymentTransaction75)
	u.TransactionInformation = append(u.TransactionInformation, newValue)
	return newValue
}
