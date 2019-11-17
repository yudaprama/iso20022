package model

// Certificate in which all currency control transactions are registered.
type TransactionCertificate2 struct {

	// Reference of the transaction, that is the underlying payment instruction or statement entry.
	ReferredDocument *CertificateReference1 `xml:"RfrdDoc"`

	// Date of the underlying transaction.
	TransactionDate *ISODate `xml:"TxDt"`

	// Type of the transaction.
	TransactionType *Exact1NumericText `xml:"TxTp"`

	// User community specific instrument.
	//
	// Usage: This element is used to specify a local transaction type to further qualify the transaction type.
	LocalInstrument *Exact5NumericText `xml:"LclInstrm"`

	// Amount as provided in the transaction to be recorded under the contract.
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`
}

func (t *TransactionCertificate2) AddReferredDocument() *CertificateReference1 {
	t.ReferredDocument = new(CertificateReference1)
	return t.ReferredDocument
}

func (t *TransactionCertificate2) SetTransactionDate(value string) {
	t.TransactionDate = (*ISODate)(&value)
}

func (t *TransactionCertificate2) SetTransactionType(value string) {
	t.TransactionType = (*Exact1NumericText)(&value)
}

func (t *TransactionCertificate2) SetLocalInstrument(value string) {
	t.LocalInstrument = (*Exact5NumericText)(&value)
}

func (t *TransactionCertificate2) SetAmount(value, currency string) {
	t.Amount = NewActiveCurrencyAndAmount(value, currency)
}
