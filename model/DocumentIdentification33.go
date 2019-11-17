package model

// Identification of a document as well as the document number.
type DocumentIdentification33 struct {

	// Unique identifier of the document (message) assigned either by the account servicer or the account owner.
	Identification *DocumentIdentification3Choice `xml:"Id"`

	// Identification of the type of document.
	DocumentNumber *DocumentNumber5Choice `xml:"DocNb,omitempty"`
}

func (d *DocumentIdentification33) AddIdentification() *DocumentIdentification3Choice {
	d.Identification = new(DocumentIdentification3Choice)
	return d.Identification
}

func (d *DocumentIdentification33) AddDocumentNumber() *DocumentNumber5Choice {
	d.DocumentNumber = new(DocumentNumber5Choice)
	return d.DocumentNumber
}
