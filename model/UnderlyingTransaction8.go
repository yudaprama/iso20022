package model

// Identifies the underlying (group of) transaction(s) to which the investigation applies.
type UnderlyingTransaction8 struct {

	// Provides information on the original message, to which the cancellation refers.
	OriginalGroupInformationAndCancellation *OriginalGroupHeader4 `xml:"OrgnlGrpInfAndCxl,omitempty"`

	// Provides information on the original transactions to which the cancellation request message refers.
	TransactionInformation []*PaymentTransaction48 `xml:"TxInf,omitempty"`
}

func (u *UnderlyingTransaction8) AddOriginalGroupInformationAndCancellation() *OriginalGroupHeader4 {
	u.OriginalGroupInformationAndCancellation = new(OriginalGroupHeader4)
	return u.OriginalGroupInformationAndCancellation
}

func (u *UnderlyingTransaction8) AddTransactionInformation() *PaymentTransaction48 {
	newValue := new(PaymentTransaction48)
	u.TransactionInformation = append(u.TransactionInformation, newValue)
	return newValue
}
