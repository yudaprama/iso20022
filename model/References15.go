package model

// Reference to the transaction identifier issued by the counterparty. Building block may also be used to reference a previous transaction, or tie a set of messages together.
type References15 struct {

	// Collective reference identifying a set of messages.
	PoolReference *AdditionalReference2 `xml:"PoolRef,omitempty"`

	// Reference of the linked message that was previously sent.
	PreviousReference *AdditionalReference2 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *AdditionalReference2 `xml:"RltdRef,omitempty"`
}

func (r *References15) AddPoolReference() *AdditionalReference2 {
	r.PoolReference = new(AdditionalReference2)
	return r.PoolReference
}

func (r *References15) AddPreviousReference() *AdditionalReference2 {
	r.PreviousReference = new(AdditionalReference2)
	return r.PreviousReference
}

func (r *References15) AddRelatedReference() *AdditionalReference2 {
	r.RelatedReference = new(AdditionalReference2)
	return r.RelatedReference
}
