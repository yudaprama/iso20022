package model

// General information that unambiguously identifies a document, such as identification number and issue date time.
type DocumentGeneralInformation1 struct {

	// Specifies the type of the document, for example commercial invoice.
	DocumentType *DocumentType4Code `xml:"DocTp"`

	// Unique identifier of the document.
	DocumentNumber *Max35Text `xml:"DocNb"`

	// Specifies the identification sequence number for a specific couple sender/receiver.
	SenderReceiverSequenceIdentification *Max140Text `xml:"SndrRcvrSeqId,omitempty"`

	// Issue date of the document.
	IssueDate *ISODate `xml:"IsseDt"`

	// URL (Uniform Resource Locator) related to the document.
	URL *Max256Text `xml:"URL,omitempty"`
}

func (d *DocumentGeneralInformation1) SetDocumentType(value string) {
	d.DocumentType = (*DocumentType4Code)(&value)
}

func (d *DocumentGeneralInformation1) SetDocumentNumber(value string) {
	d.DocumentNumber = (*Max35Text)(&value)
}

func (d *DocumentGeneralInformation1) SetSenderReceiverSequenceIdentification(value string) {
	d.SenderReceiverSequenceIdentification = (*Max140Text)(&value)
}

func (d *DocumentGeneralInformation1) SetIssueDate(value string) {
	d.IssueDate = (*ISODate)(&value)
}

func (d *DocumentGeneralInformation1) SetURL(value string) {
	d.URL = (*Max256Text)(&value)
}
