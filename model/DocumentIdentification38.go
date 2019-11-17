package model

// Identification of a document as well as the document number and type of link.
type DocumentIdentification38 struct {

	// Unique identifier of the document (message) assigned either by the account servicer or the account owner.
	Identification *DocumentIdentification4Choice `xml:"Id"`

	// Identification of the type of document.
	DocumentNumber *DocumentNumber6Choice `xml:"DocNb,omitempty"`

	// Specifies when this document is to be processed relative to an other referred document.
	LinkageType *ProcessingPosition10Choice `xml:"LkgTp,omitempty"`
}

func (d *DocumentIdentification38) AddIdentification() *DocumentIdentification4Choice {
	d.Identification = new(DocumentIdentification4Choice)
	return d.Identification
}

func (d *DocumentIdentification38) AddDocumentNumber() *DocumentNumber6Choice {
	d.DocumentNumber = new(DocumentNumber6Choice)
	return d.DocumentNumber
}

func (d *DocumentIdentification38) AddLinkageType() *ProcessingPosition10Choice {
	d.LinkageType = new(ProcessingPosition10Choice)
	return d.LinkageType
}
