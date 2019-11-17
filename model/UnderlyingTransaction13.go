package model

// Identifies the underlying (group of) transaction(s) to which the investigation applies.
type UnderlyingTransaction13 struct {

	// Provides information on the original message, to which the cancellation refers.
	OriginalGroupInformationAndCancellation *OriginalGroupHeader4 `xml:"OrgnlGrpInfAndCxl,omitempty"`

	// Provides information on the original transactions to which the cancellation request message refers.
	TransactionInformation []*PaymentTransaction62 `xml:"TxInf,omitempty"`
}

func (u *UnderlyingTransaction13) AddOriginalGroupInformationAndCancellation() *OriginalGroupHeader4 {
	u.OriginalGroupInformationAndCancellation = new(OriginalGroupHeader4)
	return u.OriginalGroupInformationAndCancellation
}

func (u *UnderlyingTransaction13) AddTransactionInformation() *PaymentTransaction62 {
	newValue := new(PaymentTransaction62)
	u.TransactionInformation = append(u.TransactionInformation, newValue)
	return newValue
}
