package model

// Choice of references used to reference a previous transaction.
type References61Choice struct {

	// Reference to a linked message that was previously received.
	RelatedReference []*AdditionalReference8 `xml:"RltdRef"`

	// Reference to a linked proprietary message or reference of a system that was previously received.
	OtherReference []*AdditionalReference8 `xml:"OthrRef"`
}

func (r *References61Choice) AddRelatedReference() *AdditionalReference8 {
	newValue := new(AdditionalReference8)
	r.RelatedReference = append(r.RelatedReference, newValue)
	return newValue
}

func (r *References61Choice) AddOtherReference() *AdditionalReference8 {
	newValue := new(AdditionalReference8)
	r.OtherReference = append(r.OtherReference, newValue)
	return newValue
}
