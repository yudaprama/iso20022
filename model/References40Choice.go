package model

// Choice of references used to reference a previous transaction.
type References40Choice struct {

	// Reference to a linked message that was previously received.
	RelatedReference []*AdditionalReference3 `xml:"RltdRef"`

	// Reference to a linked proprietary message or reference of a system that was previously received.
	OtherReference []*AdditionalReference3 `xml:"OthrRef"`
}

func (r *References40Choice) AddRelatedReference() *AdditionalReference3 {
	newValue := new(AdditionalReference3)
	r.RelatedReference = append(r.RelatedReference, newValue)
	return newValue
}

func (r *References40Choice) AddOtherReference() *AdditionalReference3 {
	newValue := new(AdditionalReference3)
	r.OtherReference = append(r.OtherReference, newValue)
	return newValue
}
