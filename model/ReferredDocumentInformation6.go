package model

// Set of elements used to identify the documents referred to in the remittance information.
type ReferredDocumentInformation6 struct {

	// Specifies the type of referred document.
	Type *ReferredDocumentType4 `xml:"Tp,omitempty"`

	// Unique and unambiguous identification of the referred document.
	Number *Max35Text `xml:"Nb,omitempty"`

	// Date associated with the referred document.
	RelatedDate *ISODate `xml:"RltdDt,omitempty"`
}

func (r *ReferredDocumentInformation6) AddType() *ReferredDocumentType4 {
	r.Type = new(ReferredDocumentType4)
	return r.Type
}

func (r *ReferredDocumentInformation6) SetNumber(value string) {
	r.Number = (*Max35Text)(&value)
}

func (r *ReferredDocumentInformation6) SetRelatedDate(value string) {
	r.RelatedDate = (*ISODate)(&value)
}
