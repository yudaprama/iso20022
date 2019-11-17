package model

// Replacement of an existing content by a different one.
type Replacement2 struct {

	// Content of the current element.
	CurrentValue *Max350Text `xml:"CurVal"`

	// Content of the new element.
	ProposedValue *Max350Text `xml:"PropsdVal"`
}

func (r *Replacement2) SetCurrentValue(value string) {
	r.CurrentValue = (*Max350Text)(&value)
}

func (r *Replacement2) SetProposedValue(value string) {
	r.ProposedValue = (*Max350Text)(&value)
}
