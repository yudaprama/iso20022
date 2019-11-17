package model

// Choice of format for the type of document format.
type DocumentFormat1Choice struct {

	// Document format.
	//
	Code *ExternalDocumentFormat1Code `xml:"Cd"`

	// Document format expressed as a proprietary code.
	Proprietary *GenericIdentification1 `xml:"Prtry"`
}

func (d *DocumentFormat1Choice) SetCode(value string) {
	d.Code = (*ExternalDocumentFormat1Code)(&value)
}

func (d *DocumentFormat1Choice) AddProprietary() *GenericIdentification1 {
	d.Proprietary = new(GenericIdentification1)
	return d.Proprietary
}
