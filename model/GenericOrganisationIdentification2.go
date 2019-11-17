package model

// Information related to an identification, eg, party identification or account identification.
type GenericOrganisationIdentification2 struct {

	// Identification assigned by an institution.
	Identification *RestrictedFINXMax35Text `xml:"Id"`

	// Name of the identification scheme.
	SchemeName *OrganisationIdentificationSchemeName2Choice `xml:"SchmeNm,omitempty"`

	// Entity that assigns the identification.
	Issuer *RestrictedFINXMax35Text `xml:"Issr,omitempty"`
}

func (g *GenericOrganisationIdentification2) SetIdentification(value string) {
	g.Identification = (*RestrictedFINXMax35Text)(&value)
}

func (g *GenericOrganisationIdentification2) AddSchemeName() *OrganisationIdentificationSchemeName2Choice {
	g.SchemeName = new(OrganisationIdentificationSchemeName2Choice)
	return g.SchemeName
}

func (g *GenericOrganisationIdentification2) SetIssuer(value string) {
	g.Issuer = (*RestrictedFINXMax35Text)(&value)
}
