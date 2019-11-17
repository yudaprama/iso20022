package model

// Specifies the type of the document referred in the remittance information.
type ReferredDocumentType3Choice struct {

	// Document type in a coded form.
	Code *DocumentType6Code `xml:"Cd"`

	// Proprietary identification of the type of the remittance document.
	Proprietary *Max35Text `xml:"Prtry"`
}

func (r *ReferredDocumentType3Choice) SetCode(value string) {
	r.Code = (*DocumentType6Code)(&value)
}

func (r *ReferredDocumentType3Choice) SetProprietary(value string) {
	r.Proprietary = (*Max35Text)(&value)
}
