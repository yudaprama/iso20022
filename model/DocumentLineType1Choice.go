package model

// Specifies the type of the document line identification.
type DocumentLineType1Choice struct {

	// Line identification type in a coded form.
	Code *ExternalDocumentLineType1Code `xml:"Cd"`

	// Proprietary identification of the type of the remittance document.
	Proprietary *Max35Text `xml:"Prtry"`
}

func (d *DocumentLineType1Choice) SetCode(value string) {
	d.Code = (*ExternalDocumentLineType1Code)(&value)
}

func (d *DocumentLineType1Choice) SetProprietary(value string) {
	d.Proprietary = (*Max35Text)(&value)
}
