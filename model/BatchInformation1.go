package model

// Identifies the underlying batches.
type BatchInformation1 struct {

	// Point to point reference assigned by the sending party to unambiguously identify the batch of transactions.
	MessageIdentification *Max35Text `xml:"MsgId,omitempty"`

	// Reference assigned by a sending party to unambiguously identify a payment information block within a payment message.
	PaymentInformationIdentification *Max35Text `xml:"PmtInfId,omitempty"`

	// Number of individual transactions included in the batch.
	NumberOfTransactions *Max15NumericText `xml:"NbOfTxs,omitempty"`
}

func (b *BatchInformation1) SetMessageIdentification(value string) {
	b.MessageIdentification = (*Max35Text)(&value)
}

func (b *BatchInformation1) SetPaymentInformationIdentification(value string) {
	b.PaymentInformationIdentification = (*Max35Text)(&value)
}

func (b *BatchInformation1) SetNumberOfTransactions(value string) {
	b.NumberOfTransactions = (*Max15NumericText)(&value)
}
