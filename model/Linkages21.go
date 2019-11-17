package model

// Information related to a linked transaction.
type Linkages21 struct {

	// Reference to the linked transaction.
	Reference *References1Choice `xml:"Ref"`
}

func (l *Linkages21) AddReference() *References1Choice {
	l.Reference = new(References1Choice)
	return l.Reference
}
