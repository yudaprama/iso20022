package model

// Identification of a document and type of link.
type DocumentIdentification49 struct {

	// Identifies the document.
	Identification *RestrictedFINXMax16Text `xml:"Id"`

	// Specifies when this document is to be processed relative to an other referred document.
	LinkageType *ProcessingPosition22Choice `xml:"LkgTp,omitempty"`
}

func (d *DocumentIdentification49) SetIdentification(value string) {
	d.Identification = (*RestrictedFINXMax16Text)(&value)
}

func (d *DocumentIdentification49) AddLinkageType() *ProcessingPosition22Choice {
	d.LinkageType = new(ProcessingPosition22Choice)
	return d.LinkageType
}
