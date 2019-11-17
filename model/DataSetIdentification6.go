package model

// Identification of a data set.
type DataSetIdentification6 struct {

	// Name of the data set.
	Name *Max256Text `xml:"Nm,omitempty"`

	// Category of data set.
	Type *DataSetCategory9Code `xml:"Tp"`

	// Version of the data set.
	Version *Max256Text `xml:"Vrsn,omitempty"`

	// Date and time of creation of the data set.
	CreationDateTime *ISODateTime `xml:"CreDtTm,omitempty"`
}

func (d *DataSetIdentification6) SetName(value string) {
	d.Name = (*Max256Text)(&value)
}

func (d *DataSetIdentification6) SetType(value string) {
	d.Type = (*DataSetCategory9Code)(&value)
}

func (d *DataSetIdentification6) SetVersion(value string) {
	d.Version = (*Max256Text)(&value)
}

func (d *DataSetIdentification6) SetCreationDateTime(value string) {
	d.CreationDateTime = (*ISODateTime)(&value)
}
