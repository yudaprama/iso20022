package model

// Identifies a document by a unique identification.
type DocumentIdentification9 struct {

	// Identifies the document.
	Identification *Max35Text `xml:"Id"`
}

func (d *DocumentIdentification9) SetIdentification(value string) {
	d.Identification = (*Max35Text)(&value)
}
