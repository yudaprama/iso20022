package model

// Identification of a document as well as the document number.
type DocumentIdentification34 struct {

	// Unique identifier of the document (message) assigned either by the account servicer or the account owner.
	Identification *DocumentIdentification4Choice `xml:"Id"`

	// Identification of the type of document.
	DocumentNumber *DocumentNumber6Choice `xml:"DocNb,omitempty"`
}

func (d *DocumentIdentification34) AddIdentification() *DocumentIdentification4Choice {
	d.Identification = new(DocumentIdentification4Choice)
	return d.Identification
}

func (d *DocumentIdentification34) AddDocumentNumber() *DocumentNumber6Choice {
	d.DocumentNumber = new(DocumentNumber6Choice)
	return d.DocumentNumber
}
