package model

// Identification of a data set.
type DataSetIdentification5 struct {

	// Name of the data set.
	Name *Max256Text `xml:"Nm"`

	// Category of data set.
	Type *DataSetCategory8Code `xml:"Tp"`

	// Version of the data set.
	Version *Max256Text `xml:"Vrsn,omitempty"`

	// Date and time of creation of the data set.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`
}

func (d *DataSetIdentification5) SetName(value string) {
	d.Name = (*Max256Text)(&value)
}

func (d *DataSetIdentification5) SetType(value string) {
	d.Type = (*DataSetCategory8Code)(&value)
}

func (d *DataSetIdentification5) SetVersion(value string) {
	d.Version = (*Max256Text)(&value)
}

func (d *DataSetIdentification5) SetCreationDateTime(value string) {
	d.CreationDateTime = (*ISODateTime)(&value)
}
