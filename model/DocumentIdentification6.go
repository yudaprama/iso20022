package model

// Identifies a document by a unique identification and a version.
// Also provides reference to a baseline amendment number.
type DocumentIdentification6 struct {

	// Identification of a set of data.
	Identification *Max35Text `xml:"Id"`

	// Unambiguous identification of the version of a set of data. Example: Version 1.
	Version *Number `xml:"Vrsn"`

	// Number that is assigned sequentially by the TSU to a baseline amendment.
	AmendmentSequenceNumber *Max3NumericText `xml:"AmdmntSeqNb,omitempty"`
}

func (d *DocumentIdentification6) SetIdentification(value string) {
	d.Identification = (*Max35Text)(&value)
}

func (d *DocumentIdentification6) SetVersion(value string) {
	d.Version = (*Number)(&value)
}

func (d *DocumentIdentification6) SetAmendmentSequenceNumber(value string) {
	d.AmendmentSequenceNumber = (*Max3NumericText)(&value)
}
