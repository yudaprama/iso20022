package model

// Unique and unambiguous way to identify a person.
type PersonIdentification6 struct {

	// Entity that assigns the identifier.
	Issuer *Max35Text `xml:"Issr,omitempty"`

	// Personal identification type.
	PersonIdentificationType *PersonIdentificationType1Choice `xml:"PrsnIdTp"`
}

func (p *PersonIdentification6) SetIssuer(value string) {
	p.Issuer = (*Max35Text)(&value)
}

func (p *PersonIdentification6) AddPersonIdentificationType() *PersonIdentificationType1Choice {
	p.PersonIdentificationType = new(PersonIdentificationType1Choice)
	return p.PersonIdentificationType
}
