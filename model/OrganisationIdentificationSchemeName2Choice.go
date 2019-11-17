package model

// Sets of elements to identify a name of the organisation identification scheme.
type OrganisationIdentificationSchemeName2Choice struct {

	// Name of the identification scheme, in a coded form as published in an external list.
	Code *ExternalOrganisationIdentification1Code `xml:"Cd"`

	// Name of the identification scheme, in a free text form.
	Proprietary *RestrictedFINXMax35Text `xml:"Prtry"`
}

func (o *OrganisationIdentificationSchemeName2Choice) SetCode(value string) {
	o.Code = (*ExternalOrganisationIdentification1Code)(&value)
}

func (o *OrganisationIdentificationSchemeName2Choice) SetProprietary(value string) {
	o.Proprietary = (*RestrictedFINXMax35Text)(&value)
}
