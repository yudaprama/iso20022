package model

// Information related to the identification of an individual person.
type GenericIdentification10 struct {

	// Name or number assigned by an entity to enable recognition of that entity, eg, account identifier.
	Identification *Max35Text `xml:"Id"`

	// Specifies the nature of the identification.
	IdentificationType *PersonIdentificationType1Code `xml:"IdTp"`

	// Specifies the nature of the identification.
	ExtendedIdentificationType *Extended350Code `xml:"XtndedIdTp"`
}

func (g *GenericIdentification10) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GenericIdentification10) SetIdentificationType(value string) {
	g.IdentificationType = (*PersonIdentificationType1Code)(&value)
}

func (g *GenericIdentification10) SetExtendedIdentificationType(value string) {
	g.ExtendedIdentificationType = (*Extended350Code)(&value)
}
