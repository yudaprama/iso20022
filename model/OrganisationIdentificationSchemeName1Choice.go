package model

// Sets of elements to identify a name of the organisation identification scheme.
type OrganisationIdentificationSchemeName1Choice struct {

	// Name of the identification scheme, in a coded form as published in an external list.
	Code *ExternalOrganisationIdentification1Code `xml:"Cd"`

	// Name of the identification scheme, in a free text form.
	Proprietary *Max35Text `xml:"Prtry"`
}

func (o *OrganisationIdentificationSchemeName1Choice) SetCode(value string) {
	o.Code = (*ExternalOrganisationIdentification1Code)(&value)
}

func (o *OrganisationIdentificationSchemeName1Choice) SetProprietary(value string) {
	o.Proprietary = (*Max35Text)(&value)
}
