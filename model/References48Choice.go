package model

// Choice of references used to reference a previous transaction.
type References48Choice struct {

	// Reference to a linked message that was previously sent.
	PreviousReference *AdditionalReference6 `xml:"PrvsRef"`

	// Reference to a linked message sent in a proprietary way or the reference of a system.
	OtherReference *AdditionalReference6 `xml:"OthrRef"`
}

func (r *References48Choice) AddPreviousReference() *AdditionalReference6 {
	r.PreviousReference = new(AdditionalReference6)
	return r.PreviousReference
}

func (r *References48Choice) AddOtherReference() *AdditionalReference6 {
	r.OtherReference = new(AdditionalReference6)
	return r.OtherReference
}
