package model

// Specifies the type of the document referred in the remittance information.
type ReferredDocumentType2 struct {

	// Provides the type details of the referred document.
	CodeOrProprietary *ReferredDocumentType1Choice `xml:"CdOrPrtry"`

	// Identification of the issuer of the reference document type.
	Issuer *Max35Text `xml:"Issr,omitempty"`
}

func (r *ReferredDocumentType2) AddCodeOrProprietary() *ReferredDocumentType1Choice {
	r.CodeOrProprietary = new(ReferredDocumentType1Choice)
	return r.CodeOrProprietary
}

func (r *ReferredDocumentType2) SetIssuer(value string) {
	r.Issuer = (*Max35Text)(&value)
}
