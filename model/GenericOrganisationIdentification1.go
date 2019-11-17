package model

// Information related to an identification of an organisation.
type GenericOrganisationIdentification1 struct {

	// Identification assigned by an institution.
	Identification *Max35Text `xml:"Id"`

	// Name of the identification scheme.
	SchemeName *OrganisationIdentificationSchemeName1Choice `xml:"SchmeNm,omitempty"`

	// Entity that assigns the identification.
	Issuer *Max35Text `xml:"Issr,omitempty"`
}

func (g *GenericOrganisationIdentification1) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GenericOrganisationIdentification1) AddSchemeName() *OrganisationIdentificationSchemeName1Choice {
	g.SchemeName = new(OrganisationIdentificationSchemeName1Choice)
	return g.SchemeName
}

func (g *GenericOrganisationIdentification1) SetIssuer(value string) {
	g.Issuer = (*Max35Text)(&value)
}
