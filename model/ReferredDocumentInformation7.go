package model

// Set of elements used to identify the documents referred to in the remittance information.
type ReferredDocumentInformation7 struct {

	// Specifies the type of referred document.
	Type *ReferredDocumentType4 `xml:"Tp,omitempty"`

	// Unique and unambiguous identification of the referred document.
	Number *Max35Text `xml:"Nb,omitempty"`

	// Date associated with the referred document.
	RelatedDate *ISODate `xml:"RltdDt,omitempty"`

	// Set of elements used to provide the content of the referred document line.
	LineDetails []*DocumentLineInformation1 `xml:"LineDtls,omitempty"`
}

func (r *ReferredDocumentInformation7) AddType() *ReferredDocumentType4 {
	r.Type = new(ReferredDocumentType4)
	return r.Type
}

func (r *ReferredDocumentInformation7) SetNumber(value string) {
	r.Number = (*Max35Text)(&value)
}

func (r *ReferredDocumentInformation7) SetRelatedDate(value string) {
	r.RelatedDate = (*ISODate)(&value)
}

func (r *ReferredDocumentInformation7) AddLineDetails() *DocumentLineInformation1 {
	newValue := new(DocumentLineInformation1)
	r.LineDetails = append(r.LineDetails, newValue)
	return newValue
}
