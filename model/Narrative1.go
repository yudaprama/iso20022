package model

// Narrative information for an undertaking.
type Narrative1 struct {

	// Type of term or condition.
	Type *NarrativeType1Choice `xml:"Tp,omitempty"`

	// Narrative text.
	Text []*Max20000Text `xml:"Txt"`
}

func (n *Narrative1) AddType() *NarrativeType1Choice {
	n.Type = new(NarrativeType1Choice)
	return n.Type
}

func (n *Narrative1) AddText(value string) {
	n.Text = append(n.Text, (*Max20000Text)(&value))
}
