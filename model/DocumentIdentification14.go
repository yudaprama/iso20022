package model

// Identification of a document as well as the document number.
type DocumentIdentification14 struct {

	// Unique identifier of the document (message) assigned either by the account servicer or the account owner.
	Identification *DocumentIdentification1Choice `xml:"Id"`

	// Identification of the type of document.
	DocumentNumber *DocumentNumber1Choice `xml:"DocNb,omitempty"`
}

func (d *DocumentIdentification14) AddIdentification() *DocumentIdentification1Choice {
	d.Identification = new(DocumentIdentification1Choice)
	return d.Identification
}

func (d *DocumentIdentification14) AddDocumentNumber() *DocumentNumber1Choice {
	d.DocumentNumber = new(DocumentNumber1Choice)
	return d.DocumentNumber
}
