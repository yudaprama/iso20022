package model

// Identifies the documents referred to in the remittance information.
type DocumentLineIdentification1 struct {

	// Specifies the type of referred document line identification.
	Type *DocumentLineType1 `xml:"Tp,omitempty"`

	// Identification of the type specified for the referred document line.
	Number *Max35Text `xml:"Nb,omitempty"`

	// Date associated with the referred document line.
	RelatedDate *ISODate `xml:"RltdDt,omitempty"`
}

func (d *DocumentLineIdentification1) AddType() *DocumentLineType1 {
	d.Type = new(DocumentLineType1)
	return d.Type
}

func (d *DocumentLineIdentification1) SetNumber(value string) {
	d.Number = (*Max35Text)(&value)
}

func (d *DocumentLineIdentification1) SetRelatedDate(value string) {
	d.RelatedDate = (*ISODate)(&value)
}
