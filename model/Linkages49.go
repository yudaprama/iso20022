package model

// Information related to a linked transaction.
type Linkages49 struct {

	// Reference to the linked transaction.
	Reference *References58Choice `xml:"Ref"`
}

func (l *Linkages49) AddReference() *References58Choice {
	l.Reference = new(References58Choice)
	return l.Reference
}
