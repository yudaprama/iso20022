package model

// Identification of a data set.
type DataSetIdentification1 struct {

	// Name of the data set.
	Name *Max256Text `xml:"Nm"`

	// Category of data set.
	Type *DataSetCategory1Code `xml:"Tp"`

	// Version of the data set.
	Version *Max256Text `xml:"Vrsn,omitempty"`

	// Date and time of creation of the data set.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`
}

func (d *DataSetIdentification1) SetName(value string) {
	d.Name = (*Max256Text)(&value)
}

func (d *DataSetIdentification1) SetType(value string) {
	d.Type = (*DataSetCategory1Code)(&value)
}

func (d *DataSetIdentification1) SetVersion(value string) {
	d.Version = (*Max256Text)(&value)
}

func (d *DataSetIdentification1) SetCreationDateTime(value string) {
	d.CreationDateTime = (*ISODateTime)(&value)
}
