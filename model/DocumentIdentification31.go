package model

// Identification of a document and type of link.
type DocumentIdentification31 struct {

	// Identifies the document.
	Identification *Max35Text `xml:"Id"`

	// Specifies when this document is to be processed relative to an other referred document.
	LinkageType *ProcessingPosition7Choice `xml:"LkgTp,omitempty"`
}

func (d *DocumentIdentification31) SetIdentification(value string) {
	d.Identification = (*Max35Text)(&value)
}

func (d *DocumentIdentification31) AddLinkageType() *ProcessingPosition7Choice {
	d.LinkageType = new(ProcessingPosition7Choice)
	return d.LinkageType
}
