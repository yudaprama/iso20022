package model

// Identification of a data set.
type DataSetIdentification4 struct {

	// Name of the data set.
	Name *Max256Text `xml:"Nm,omitempty"`

	// Category of data set.
	Type *DataSetCategory4Code `xml:"Tp"`

	// Version of the data set.
	Version *Max256Text `xml:"Vrsn,omitempty"`

	// Date and time of creation of the data set.
	CreationDateTime *ISODateTime `xml:"CreDtTm,omitempty"`
}

func (d *DataSetIdentification4) SetName(value string) {
	d.Name = (*Max256Text)(&value)
}

func (d *DataSetIdentification4) SetType(value string) {
	d.Type = (*DataSetCategory4Code)(&value)
}

func (d *DataSetIdentification4) SetVersion(value string) {
	d.Version = (*Max256Text)(&value)
}

func (d *DataSetIdentification4) SetCreationDateTime(value string) {
	d.CreationDateTime = (*ISODateTime)(&value)
}
