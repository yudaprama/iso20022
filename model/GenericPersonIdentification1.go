package model

// Information related to an identification of a person.
type GenericPersonIdentification1 struct {

	// Unique and unambiguous identification of a person.
	Identification *Max35Text `xml:"Id"`

	// Name of the identification scheme.
	SchemeName *PersonIdentificationSchemeName1Choice `xml:"SchmeNm,omitempty"`

	// Entity that assigns the identification.
	Issuer *Max35Text `xml:"Issr,omitempty"`
}

func (g *GenericPersonIdentification1) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GenericPersonIdentification1) AddSchemeName() *PersonIdentificationSchemeName1Choice {
	g.SchemeName = new(PersonIdentificationSchemeName1Choice)
	return g.SchemeName
}

func (g *GenericPersonIdentification1) SetIssuer(value string) {
	g.Issuer = (*Max35Text)(&value)
}
