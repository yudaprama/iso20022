package model

// Structured information supplied to fully identify the documents referred to in the remittance information.
type ReferredDocumentInformation1 struct {

	// Provides the type of the referred document.
	ReferredDocumentType *ReferredDocumentType1 `xml:"RfrdDocTp,omitempty"`

	// Unique and unambiguous identification number of the referred document.
	ReferredDocumentNumber *Max35Text `xml:"RfrdDocNb,omitempty"`
}

func (r *ReferredDocumentInformation1) AddReferredDocumentType() *ReferredDocumentType1 {
	r.ReferredDocumentType = new(ReferredDocumentType1)
	return r.ReferredDocumentType
}

func (r *ReferredDocumentInformation1) SetReferredDocumentNumber(value string) {
	r.ReferredDocumentNumber = (*Max35Text)(&value)
}
