package model

// Identifies a document by a unique identification and a version together with the submitter of the document.
// Also specifies the type of document and an index for easy referencing.
type DocumentIdentification10 struct {

	// Identification of a set of data.
	Identification *Max35Text `xml:"Id"`

	// Unambiguous identification of the version of a set of data. Example: Version 1.
	Version *Number `xml:"Vrsn"`

	// Identifies the type of data set.
	Type *DataSetType2Code `xml:"Tp"`

	// Uniquely identifies the financial institution which has submitted the set of data by using a BIC.
	Submitter *BICIdentification1 `xml:"Submitr"`

	// Index assigned to the document, to allow easy referencing.
	DocumentIndex *Max3NumericText `xml:"DocIndx"`
}

func (d *DocumentIdentification10) SetIdentification(value string) {
	d.Identification = (*Max35Text)(&value)
}

func (d *DocumentIdentification10) SetVersion(value string) {
	d.Version = (*Number)(&value)
}

func (d *DocumentIdentification10) SetType(value string) {
	d.Type = (*DataSetType2Code)(&value)
}

func (d *DocumentIdentification10) AddSubmitter() *BICIdentification1 {
	d.Submitter = new(BICIdentification1)
	return d.Submitter
}

func (d *DocumentIdentification10) SetDocumentIndex(value string) {
	d.DocumentIndex = (*Max3NumericText)(&value)
}
