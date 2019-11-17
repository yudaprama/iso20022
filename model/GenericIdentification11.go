package model

// Information related to an identification, eg, party identification or account identification.
type GenericIdentification11 struct {

	// Name or number assigned by an entity to enable recognition of that entity, eg, account identifier.
	Identification *Max35Text `xml:"Id"`

	// Specifies the nature of the identifier.
	IdentificationType *PersonIdentificationType2Code `xml:"IdTp"`

	// Specifies a type of individual identification.
	ExtendedIdentificationType *Extended350Code `xml:"XtndedIdTp"`

	// Entity that assigns the identifier.
	Issuer *Max35Text `xml:"Issr,omitempty"`
}

func (g *GenericIdentification11) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GenericIdentification11) SetIdentificationType(value string) {
	g.IdentificationType = (*PersonIdentificationType2Code)(&value)
}

func (g *GenericIdentification11) SetExtendedIdentificationType(value string) {
	g.ExtendedIdentificationType = (*Extended350Code)(&value)
}

func (g *GenericIdentification11) SetIssuer(value string) {
	g.Issuer = (*Max35Text)(&value)
}
