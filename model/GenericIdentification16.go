package model

// Specifies a generic identification.
type GenericIdentification16 struct {

	// The identifier.
	Identification *Max35Text `xml:"Id"`

	// Specifies the nature of the identifier.
	IdentificationType *PersonIdentificationType3Choice `xml:"IdTp"`

	// Party that assigns the identification.
	Issuer *Max35Text `xml:"Issr,omitempty"`
}

func (g *GenericIdentification16) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GenericIdentification16) AddIdentificationType() *PersonIdentificationType3Choice {
	g.IdentificationType = new(PersonIdentificationType3Choice)
	return g.IdentificationType
}

func (g *GenericIdentification16) SetIssuer(value string) {
	g.Issuer = (*Max35Text)(&value)
}
