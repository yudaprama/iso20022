package model

// Choice of references used to reference a previous transaction.
type References62Choice struct {

	// Reference to a linked message that was previously sent.
	PreviousReference []*AdditionalReference8 `xml:"PrvsRef"`

	// Reference to a linked proprietary message or reference of a system that was previously received.
	OtherReference []*AdditionalReference8 `xml:"OthrRef"`
}

func (r *References62Choice) AddPreviousReference() *AdditionalReference8 {
	newValue := new(AdditionalReference8)
	r.PreviousReference = append(r.PreviousReference, newValue)
	return newValue
}

func (r *References62Choice) AddOtherReference() *AdditionalReference8 {
	newValue := new(AdditionalReference8)
	r.OtherReference = append(r.OtherReference, newValue)
	return newValue
}
