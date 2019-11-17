package model

// Identifies the underlying (group of) transaction(s) to which the investigation applies.
type UnderlyingTransaction5 struct {

	// Provides information on the original message, to which the cancellation refers.
	OriginalGroupInformationAndCancellation *OriginalGroupHeader4 `xml:"OrgnlGrpInfAndCxl,omitempty"`

	// Provides information on the original transactions to which the cancellation request message refers.
	TransactionInformation []*PaymentTransaction38 `xml:"TxInf,omitempty"`
}

func (u *UnderlyingTransaction5) AddOriginalGroupInformationAndCancellation() *OriginalGroupHeader4 {
	u.OriginalGroupInformationAndCancellation = new(OriginalGroupHeader4)
	return u.OriginalGroupInformationAndCancellation
}

func (u *UnderlyingTransaction5) AddTransactionInformation() *PaymentTransaction38 {
	newValue := new(PaymentTransaction38)
	u.TransactionInformation = append(u.TransactionInformation, newValue)
	return newValue
}
