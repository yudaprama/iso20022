package model

// Identification of a document and type of link.
type DocumentIdentification15 struct {

	// Identifies the document.
	Identification *Max35Text `xml:"Id"`

	// Specifies when this document is to be processed relative to an other referred document.
	LinkageType *ProcessingPosition1Choice `xml:"LkgTp,omitempty"`
}

func (d *DocumentIdentification15) SetIdentification(value string) {
	d.Identification = (*Max35Text)(&value)
}

func (d *DocumentIdentification15) AddLinkageType() *ProcessingPosition1Choice {
	d.LinkageType = new(ProcessingPosition1Choice)
	return d.LinkageType
}
