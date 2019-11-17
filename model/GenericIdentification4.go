package model

// Information related to an identification, eg, party identification or account identification.
type GenericIdentification4 struct {

	// Identifier issued to a person for which no specific identifier has been defined.
	Identification *Max35Text `xml:"Id"`

	// Specifies the nature of the identifier.
	// Usage: IdentificationType is used to specify what kind of identifier is used. It should be used in case the identifier is different from the identifiers listed in the pre-defined identifier list.
	IdentificationType *Max35Text `xml:"IdTp"`
}

func (g *GenericIdentification4) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GenericIdentification4) SetIdentificationType(value string) {
	g.IdentificationType = (*Max35Text)(&value)
}
