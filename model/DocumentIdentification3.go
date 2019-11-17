package model

// Identifies a document by a unique identification and a version.
type DocumentIdentification3 struct {

	// Identification of a set of data.
	Identification *Max35Text `xml:"Id"`

	// Unambiguous identification of the version of a set of data. Example: Version 1.
	Version *Number `xml:"Vrsn"`
}

func (d *DocumentIdentification3) SetIdentification(value string) {
	d.Identification = (*Max35Text)(&value)
}

func (d *DocumentIdentification3) SetVersion(value string) {
	d.Version = (*Number)(&value)
}
