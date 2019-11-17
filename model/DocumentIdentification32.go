package model

// Identification of a document as well as the document number and type of link.
type DocumentIdentification32 struct {

	// Unique identifier of the document (message) assigned either by the account servicer or the account owner.
	Identification *DocumentIdentification3Choice `xml:"Id"`

	// Identification of the type of document.
	DocumentNumber *DocumentNumber5Choice `xml:"DocNb,omitempty"`

	// Specifies when this document is to be processed relative to an other referred document.
	LinkageType *ProcessingPosition7Choice `xml:"LkgTp,omitempty"`
}

func (d *DocumentIdentification32) AddIdentification() *DocumentIdentification3Choice {
	d.Identification = new(DocumentIdentification3Choice)
	return d.Identification
}

func (d *DocumentIdentification32) AddDocumentNumber() *DocumentNumber5Choice {
	d.DocumentNumber = new(DocumentNumber5Choice)
	return d.DocumentNumber
}

func (d *DocumentIdentification32) AddLinkageType() *ProcessingPosition7Choice {
	d.LinkageType = new(ProcessingPosition7Choice)
	return d.LinkageType
}
