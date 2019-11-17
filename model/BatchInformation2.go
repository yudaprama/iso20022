package model

// Set of elements used to identify the underlying batches.
type BatchInformation2 struct {

	// Point to point reference, as assigned by the sending party, to unambiguously identify the batch of transactions.
	MessageIdentification *Max35Text `xml:"MsgId,omitempty"`

	// Unique identification, as assigned by a sending party, to unambiguously identify the payment information group within the message.
	PaymentInformationIdentification *Max35Text `xml:"PmtInfId,omitempty"`

	// Number of individual transactions included in the batch.
	NumberOfTransactions *Max15NumericText `xml:"NbOfTxs,omitempty"`

	// Total amount of money reported in the batch entry.
	TotalAmount *ActiveOrHistoricCurrencyAndAmount `xml:"TtlAmt,omitempty"`

	// Indicates whether the batch entry is a credit or a debit entry.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd,omitempty"`
}

func (b *BatchInformation2) SetMessageIdentification(value string) {
	b.MessageIdentification = (*Max35Text)(&value)
}

func (b *BatchInformation2) SetPaymentInformationIdentification(value string) {
	b.PaymentInformationIdentification = (*Max35Text)(&value)
}

func (b *BatchInformation2) SetNumberOfTransactions(value string) {
	b.NumberOfTransactions = (*Max15NumericText)(&value)
}

func (b *BatchInformation2) SetTotalAmount(value, currency string) {
	b.TotalAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (b *BatchInformation2) SetCreditDebitIndicator(value string) {
	b.CreditDebitIndicator = (*CreditDebitCode)(&value)
}
