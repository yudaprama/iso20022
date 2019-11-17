package model

// Information related to an identification, for example, party identification or account identification.
type GenericIdentification49 struct {

	// Identifier issued to a person or an institution for which no other specific identifier has been defined.
	Identification *Max35Text `xml:"Id"`

	// Type of identifier.
	IdentificationType *Max35Text `xml:"IdTp"`
}

func (g *GenericIdentification49) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GenericIdentification49) SetIdentificationType(value string) {
	g.IdentificationType = (*Max35Text)(&value)
}
