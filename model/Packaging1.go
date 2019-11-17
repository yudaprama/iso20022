package model

// Physical packaging of goods for transport
type Packaging1 struct {

	// Specifies the type of packaging as a code.
	Type *ExternalPackagingType1Code `xml:"Tp,omitempty"`

	// Specifies the type of packaging as text. For instance, halogenated resin (PVC).
	Name *Max35Text `xml:"Nm,omitempty"`
}

func (p *Packaging1) SetType(value string) {
	p.Type = (*ExternalPackagingType1Code)(&value)
}

func (p *Packaging1) SetName(value string) {
	p.Name = (*Max35Text)(&value)
}
