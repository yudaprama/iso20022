package model

// Specifies the type of garnishment.
type GarnishmentType1 struct {

	// Provides the type details of the garnishment.
	CodeOrProprietary *GarnishmentType1Choice `xml:"CdOrPrtry"`

	// Identification of the issuer of the garnishment type.
	Issuer *Max35Text `xml:"Issr,omitempty"`
}

func (g *GarnishmentType1) AddCodeOrProprietary() *GarnishmentType1Choice {
	g.CodeOrProprietary = new(GarnishmentType1Choice)
	return g.CodeOrProprietary
}

func (g *GarnishmentType1) SetIssuer(value string) {
	g.Issuer = (*Max35Text)(&value)
}
