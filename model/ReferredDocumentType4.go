package model

// Specifies the type of the document referred in the remittance information.
type ReferredDocumentType4 struct {

	// Provides the type details of the referred document.
	CodeOrProprietary *ReferredDocumentType3Choice `xml:"CdOrPrtry"`

	// Identification of the issuer of the reference document type.
	Issuer *Max35Text `xml:"Issr,omitempty"`
}

func (r *ReferredDocumentType4) AddCodeOrProprietary() *ReferredDocumentType3Choice {
	r.CodeOrProprietary = new(ReferredDocumentType3Choice)
	return r.CodeOrProprietary
}

func (r *ReferredDocumentType4) SetIssuer(value string) {
	r.Issuer = (*Max35Text)(&value)
}
