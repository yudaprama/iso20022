package model

// Identifies the documents referred to in the remittance information.
type ReferredMandateDocument1 struct {

	// Specifies the type of referred document.
	Type *ReferredDocumentType4 `xml:"Tp,omitempty"`

	// Unique and unambiguous identification of the referred document.
	Number *Max35Text `xml:"Nb,omitempty"`

	// Unique and unambiguous identification as assigned by the creditor to the referred document shared with the debtor for its own reference.
	CreditorReference *Max35Text `xml:"CdtrRef,omitempty"`

	// Date associated with the referred document.
	RelatedDate *ISODate `xml:"RltdDt,omitempty"`
}

func (r *ReferredMandateDocument1) AddType() *ReferredDocumentType4 {
	r.Type = new(ReferredDocumentType4)
	return r.Type
}

func (r *ReferredMandateDocument1) SetNumber(value string) {
	r.Number = (*Max35Text)(&value)
}

func (r *ReferredMandateDocument1) SetCreditorReference(value string) {
	r.CreditorReference = (*Max35Text)(&value)
}

func (r *ReferredMandateDocument1) SetRelatedDate(value string) {
	r.RelatedDate = (*ISODate)(&value)
}
