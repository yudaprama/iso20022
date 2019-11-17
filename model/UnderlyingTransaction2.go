package model

// Set of elements used to identify the underlying (group of) transaction(s) to which the investigation applies.
type UnderlyingTransaction2 struct {

	// Set of elements used to provide information on the original messsage, to which the cancellation refers.
	OriginalGroupInformationAndCancellation *OriginalGroupInformation23 `xml:"OrgnlGrpInfAndCxl,omitempty"`

	// Set of elements used to provide information on the original transactions to which the cancellation request message refers.
	TransactionInformation []*PaymentTransactionInformation31 `xml:"TxInf,omitempty"`
}

func (u *UnderlyingTransaction2) AddOriginalGroupInformationAndCancellation() *OriginalGroupInformation23 {
	u.OriginalGroupInformationAndCancellation = new(OriginalGroupInformation23)
	return u.OriginalGroupInformationAndCancellation
}

func (u *UnderlyingTransaction2) AddTransactionInformation() *PaymentTransactionInformation31 {
	newValue := new(PaymentTransactionInformation31)
	u.TransactionInformation = append(u.TransactionInformation, newValue)
	return newValue
}
