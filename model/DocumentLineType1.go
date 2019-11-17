package model

// Specifies the type of the document line identification.
type DocumentLineType1 struct {

	// Provides the type details of the referred document line identification.
	CodeOrProprietary *DocumentLineType1Choice `xml:"CdOrPrtry"`

	// Identification of the issuer of the reference document line identificationtype.
	Issuer *Max35Text `xml:"Issr,omitempty"`
}

func (d *DocumentLineType1) AddCodeOrProprietary() *DocumentLineType1Choice {
	d.CodeOrProprietary = new(DocumentLineType1Choice)
	return d.CodeOrProprietary
}

func (d *DocumentLineType1) SetIssuer(value string) {
	d.Issuer = (*Max35Text)(&value)
}
