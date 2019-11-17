package model

// Set of elements to identify a proprietary reference.
type ProprietaryReference1 struct {

	// Identifies the type of reference reported.
	Type *Max35Text `xml:"Tp"`

	// Proprietary reference specification related to the underlying transaction.
	Reference *Max35Text `xml:"Ref"`
}

func (p *ProprietaryReference1) SetType(value string) {
	p.Type = (*Max35Text)(&value)
}

func (p *ProprietaryReference1) SetReference(value string) {
	p.Reference = (*Max35Text)(&value)
}
