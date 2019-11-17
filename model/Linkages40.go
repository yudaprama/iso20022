package model

// Information related to a linked transaction.
type Linkages40 struct {

	// Reference to the linked transaction.
	Reference *References47Choice `xml:"Ref"`
}

func (l *Linkages40) AddReference() *References47Choice {
	l.Reference = new(References47Choice)
	return l.Reference
}
