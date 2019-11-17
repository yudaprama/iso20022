package model

// Identifies a document by a unique identification.
type DocumentIdentification17 struct {

	// Identifies the document.
	Identification *RestrictedFINXMax16Text `xml:"Id"`
}

func (d *DocumentIdentification17) SetIdentification(value string) {
	d.Identification = (*RestrictedFINXMax16Text)(&value)
}
