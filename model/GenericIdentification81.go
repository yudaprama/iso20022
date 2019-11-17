package model

// Information related to the identification of a party.
type GenericIdentification81 struct {

	// Identification of a party, such as a tax or social security identifier.
	Identification *Max35Text `xml:"Id"`

	// Type of identification.
	IdentificationType *OtherIdentification3Choice `xml:"IdTp"`
}

func (g *GenericIdentification81) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GenericIdentification81) AddIdentificationType() *OtherIdentification3Choice {
	g.IdentificationType = new(OtherIdentification3Choice)
	return g.IdentificationType
}
