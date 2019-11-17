package model

// Identifies a document by a unique identification and a version together with the identification of the submitter of the document.
type DocumentIdentification1 struct {

	// Identification of a set of data.
	Identification *Max35Text `xml:"Id"`

	// Unambiguous identification of the version of a set of data. Example: Version 1.
	Version *Number `xml:"Vrsn"`

	// Uniquely identifies the financial institution which has submitted the set of data by using a BIC.
	Submitter *BICIdentification1 `xml:"Submitr"`
}

func (d *DocumentIdentification1) SetIdentification(value string) {
	d.Identification = (*Max35Text)(&value)
}

func (d *DocumentIdentification1) SetVersion(value string) {
	d.Version = (*Number)(&value)
}

func (d *DocumentIdentification1) AddSubmitter() *BICIdentification1 {
	d.Submitter = new(BICIdentification1)
	return d.Submitter
}
