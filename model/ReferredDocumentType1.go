package model

// Specifies the type of the document referred in the remittance information.
type ReferredDocumentType1 struct {

	// Document type in a coded form.
	Code *DocumentType2Code `xml:"Cd"`

	// Proprietary identification of the type of the remittance document.
	Proprietary *Max35Text `xml:"Prtry"`

	// Identification of the issuer of the reference document type.
	Issuer *Max35Text `xml:"Issr,omitempty"`
}

func (r *ReferredDocumentType1) SetCode(value string) {
	r.Code = (*DocumentType2Code)(&value)
}

func (r *ReferredDocumentType1) SetProprietary(value string) {
	r.Proprietary = (*Max35Text)(&value)
}

func (r *ReferredDocumentType1) SetIssuer(value string) {
	r.Issuer = (*Max35Text)(&value)
}
