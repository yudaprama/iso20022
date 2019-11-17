package model

// Identification of a document as well as the document number and type of link.
type DocumentIdentification13 struct {

	// Unique identifier of the document (message) assigned either by the account servicer or the account owner.
	Identification *DocumentIdentification1Choice `xml:"Id"`

	// Identification of the type of document.
	DocumentNumber *DocumentNumber1Choice `xml:"DocNb,omitempty"`

	// Specifies when this document is to be processed relative to an other referred document.
	LinkageType *ProcessingPosition1Choice `xml:"LkgTp,omitempty"`
}

func (d *DocumentIdentification13) AddIdentification() *DocumentIdentification1Choice {
	d.Identification = new(DocumentIdentification1Choice)
	return d.Identification
}

func (d *DocumentIdentification13) AddDocumentNumber() *DocumentNumber1Choice {
	d.DocumentNumber = new(DocumentNumber1Choice)
	return d.DocumentNumber
}

func (d *DocumentIdentification13) AddLinkageType() *ProcessingPosition1Choice {
	d.LinkageType = new(ProcessingPosition1Choice)
	return d.LinkageType
}
