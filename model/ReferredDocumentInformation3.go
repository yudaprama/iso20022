package model

// Set of elements used to identify the documents referred to in the remittance information.
type ReferredDocumentInformation3 struct {

	// Specifies the type of referred document.
	Type *ReferredDocumentType2 `xml:"Tp,omitempty"`

	// Unique and unambiguous identification of the referred document.
	Number *Max35Text `xml:"Nb,omitempty"`

	// Date associated with the referred document.
	RelatedDate *ISODate `xml:"RltdDt,omitempty"`
}

func (r *ReferredDocumentInformation3) AddType() *ReferredDocumentType2 {
	r.Type = new(ReferredDocumentType2)
	return r.Type
}

func (r *ReferredDocumentInformation3) SetNumber(value string) {
	r.Number = (*Max35Text)(&value)
}

func (r *ReferredDocumentInformation3) SetRelatedDate(value string) {
	r.RelatedDate = (*ISODate)(&value)
}
