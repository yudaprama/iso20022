package model

// Identifies the documents referred to in the remittance information.
type ReferredDocumentInformation4 struct {

	// Specifies the type of referred document.
	Type *ReferredDocumentType2 `xml:"Tp,omitempty"`

	// Unique and unambiguous identification of the referred document.
	Number *Max35Text `xml:"Nb,omitempty"`

	// Date associated with the referred document.
	RelatedDate *ISODate `xml:"RltdDt,omitempty"`

	// Set of elements used to provide the content of the referred document line.
	LineDetails []*DocumentLineInformation1 `xml:"LineDtls,omitempty"`
}

func (r *ReferredDocumentInformation4) AddType() *ReferredDocumentType2 {
	r.Type = new(ReferredDocumentType2)
	return r.Type
}

func (r *ReferredDocumentInformation4) SetNumber(value string) {
	r.Number = (*Max35Text)(&value)
}

func (r *ReferredDocumentInformation4) SetRelatedDate(value string) {
	r.RelatedDate = (*ISODate)(&value)
}

func (r *ReferredDocumentInformation4) AddLineDetails() *DocumentLineInformation1 {
	newValue := new(DocumentLineInformation1)
	r.LineDetails = append(r.LineDetails, newValue)
	return newValue
}
