package model

// Reference to the transaction identifier issued by the counterparty. Building block may also be used to reference a previous transaction, or tie a set of messages together.
type References20 struct {

	// Collective reference identifying a set of messages.
	PoolReference *AdditionalReference6 `xml:"PoolRef,omitempty"`

	// Reference of the linked message that was previously sent.
	PreviousReference *AdditionalReference6 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *AdditionalReference6 `xml:"RltdRef,omitempty"`
}

func (r *References20) AddPoolReference() *AdditionalReference6 {
	r.PoolReference = new(AdditionalReference6)
	return r.PoolReference
}

func (r *References20) AddPreviousReference() *AdditionalReference6 {
	r.PreviousReference = new(AdditionalReference6)
	return r.PreviousReference
}

func (r *References20) AddRelatedReference() *AdditionalReference6 {
	r.RelatedReference = new(AdditionalReference6)
	return r.RelatedReference
}
