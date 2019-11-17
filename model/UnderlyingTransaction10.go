package model

// Identifies the underlying (group of) transaction(s) to which the investigation applies.
type UnderlyingTransaction10 struct {

	// Provides information on the original message, to which the cancellation refers.
	OriginalGroupInformationAndCancellation *OriginalGroupHeader4 `xml:"OrgnlGrpInfAndCxl,omitempty"`

	// Provides information on the original transactions to which the cancellation request message refers.
	TransactionInformation []*PaymentTransaction55 `xml:"TxInf,omitempty"`
}

func (u *UnderlyingTransaction10) AddOriginalGroupInformationAndCancellation() *OriginalGroupHeader4 {
	u.OriginalGroupInformationAndCancellation = new(OriginalGroupHeader4)
	return u.OriginalGroupInformationAndCancellation
}

func (u *UnderlyingTransaction10) AddTransactionInformation() *PaymentTransaction55 {
	newValue := new(PaymentTransaction55)
	u.TransactionInformation = append(u.TransactionInformation, newValue)
	return newValue
}
