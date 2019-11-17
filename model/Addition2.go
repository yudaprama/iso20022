package model

// Addition of a new element.
type Addition2 struct {

	// Content of the new element.
	ProposedValue *Max350Text `xml:"PropsdVal,omitempty"`
}

func (a *Addition2) SetProposedValue(value string) {
	a.ProposedValue = (*Max350Text)(&value)
}
