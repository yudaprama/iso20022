package model

// Certificate record in which all currency control transactions are registered.
type TransactionCertificateRecord1 struct {

	// Unique and unambiguous identification of the certificate record.
	CertificateRecordIdentification *Max35Text `xml:"CertRcrdId"`

	// Details of the transaction for which the record has been generated.
	Transaction *TransactionCertificate2 `xml:"Tx"`

	// Contract registration details related to the certificate record.
	Contract *TransactionCertificateContract1 `xml:"Ctrct,omitempty"`

	// Documents provided as attachments to the registered contract.
	Attachment []*DocumentGeneralInformation3 `xml:"Attchmnt,omitempty"`
}

func (t *TransactionCertificateRecord1) SetCertificateRecordIdentification(value string) {
	t.CertificateRecordIdentification = (*Max35Text)(&value)
}

func (t *TransactionCertificateRecord1) AddTransaction() *TransactionCertificate2 {
	t.Transaction = new(TransactionCertificate2)
	return t.Transaction
}

func (t *TransactionCertificateRecord1) AddContract() *TransactionCertificateContract1 {
	t.Contract = new(TransactionCertificateContract1)
	return t.Contract
}

func (t *TransactionCertificateRecord1) AddAttachment() *DocumentGeneralInformation3 {
	newValue := new(DocumentGeneralInformation3)
	t.Attachment = append(t.Attachment, newValue)
	return newValue
}
