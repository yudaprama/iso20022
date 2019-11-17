package model

// Identification of a data set.
type DataSetIdentification3 struct {

	// Name of the data set.
	Name *Max256Text `xml:"Nm,omitempty"`

	// Category of data set.
	Type *DataSetCategory3Code `xml:"Tp"`

	// Version of the data set.
	Version *Max256Text `xml:"Vrsn,omitempty"`

	// Date and time of creation of the data set.
	CreationDateTime *ISODateTime `xml:"CreDtTm,omitempty"`
}

func (d *DataSetIdentification3) SetName(value string) {
	d.Name = (*Max256Text)(&value)
}

func (d *DataSetIdentification3) SetType(value string) {
	d.Type = (*DataSetCategory3Code)(&value)
}

func (d *DataSetIdentification3) SetVersion(value string) {
	d.Version = (*Max256Text)(&value)
}

func (d *DataSetIdentification3) SetCreationDateTime(value string) {
	d.CreationDateTime = (*ISODateTime)(&value)
}
