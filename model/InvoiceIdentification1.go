package model

// Identifies a document by a unique identification and a date of issue.
type InvoiceIdentification1 struct {

	// Identifies the document.
	InvoiceNumber *Max35Text `xml:"InvcNb"`

	// Date of issuance of the document.
	IssueDate *ISODate `xml:"IsseDt"`
}

func (i *InvoiceIdentification1) SetInvoiceNumber(value string) {
	i.InvoiceNumber = (*Max35Text)(&value)
}

func (i *InvoiceIdentification1) SetIssueDate(value string) {
	i.IssueDate = (*ISODate)(&value)
}
