package model

// Identification of a document and type of link.
type DocumentIdentification37 struct {

	// Identifies the document.
	Identification *RestrictedFINXMax16Text `xml:"Id"`

	// Specifies when this document is to be processed relative to an other referred document.
	LinkageType *ProcessingPosition10Choice `xml:"LkgTp,omitempty"`
}

func (d *DocumentIdentification37) SetIdentification(value string) {
	d.Identification = (*RestrictedFINXMax16Text)(&value)
}

func (d *DocumentIdentification37) AddLinkageType() *ProcessingPosition10Choice {
	d.LinkageType = new(ProcessingPosition10Choice)
	return d.LinkageType
}
