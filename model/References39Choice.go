package model

// Choice of references used to reference a previous transaction.
type References39Choice struct {

	// Reference to a linked message that was previously sent.
	PreviousReference *AdditionalReference3 `xml:"PrvsRef"`

	// Reference to a linked message sent in a proprietary way or the reference of a system.
	OtherReference *AdditionalReference3 `xml:"OthrRef"`
}

func (r *References39Choice) AddPreviousReference() *AdditionalReference3 {
	r.PreviousReference = new(AdditionalReference3)
	return r.PreviousReference
}

func (r *References39Choice) AddOtherReference() *AdditionalReference3 {
	r.OtherReference = new(AdditionalReference3)
	return r.OtherReference
}
