package model

// Specifies the type of the document referred in the remittance information.
type ReferredDocumentType1Choice struct {

	// Document type in a coded form.
	Code *DocumentType5Code `xml:"Cd"`

	// Proprietary identification of the type of the remittance document.
	Proprietary *Max35Text `xml:"Prtry"`
}

func (r *ReferredDocumentType1Choice) SetCode(value string) {
	r.Code = (*DocumentType5Code)(&value)
}

func (r *ReferredDocumentType1Choice) SetProprietary(value string) {
	r.Proprietary = (*Max35Text)(&value)
}
