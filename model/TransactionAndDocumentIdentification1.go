package model

// Transaction and document identification details.
type TransactionAndDocumentIdentification1 struct {

	// Unambiguous identification of the transaction as know by the instructing party.
	TransactionIdentification *Max35Text `xml:"TxId"`

	// Unique identifier of the document (message) assigned by the sender of the document.
	DocumentIdentification *Max35Text `xml:"DocId,omitempty"`

	// Date and time at which the transaction was created by the instructing party in its business application.
	CreationDateTime *DateAndDateTimeChoice `xml:"CreDtTm,omitempty"`

	// Specifies if this document is a copy, a duplicate, or a duplicate of a copy.
	CopyDuplicate *CopyDuplicate1Code `xml:"CpyDplct,omitempty"`
}

func (t *TransactionAndDocumentIdentification1) SetTransactionIdentification(value string) {
	t.TransactionIdentification = (*Max35Text)(&value)
}

func (t *TransactionAndDocumentIdentification1) SetDocumentIdentification(value string) {
	t.DocumentIdentification = (*Max35Text)(&value)
}

func (t *TransactionAndDocumentIdentification1) AddCreationDateTime() *DateAndDateTimeChoice {
	t.CreationDateTime = new(DateAndDateTimeChoice)
	return t.CreationDateTime
}

func (t *TransactionAndDocumentIdentification1) SetCopyDuplicate(value string) {
	t.CopyDuplicate = (*CopyDuplicate1Code)(&value)
}
