package model

// Choice of references used to reference a previous transaction.
type References49Choice struct {

	// Reference to a linked message that was previously received.
	RelatedReference []*AdditionalReference6 `xml:"RltdRef"`

	// Reference to a linked proprietary message or reference of a system that was previously received.
	OtherReference []*AdditionalReference6 `xml:"OthrRef"`
}

func (r *References49Choice) AddRelatedReference() *AdditionalReference6 {
	newValue := new(AdditionalReference6)
	r.RelatedReference = append(r.RelatedReference, newValue)
	return newValue
}

func (r *References49Choice) AddOtherReference() *AdditionalReference6 {
	newValue := new(AdditionalReference6)
	r.OtherReference = append(r.OtherReference, newValue)
	return newValue
}
