package model

// Choice of formats for the type of organisation.
type OrganisationType1Choice struct {

	// Type of organisation expressed as a code.
	Code *OrganisationType1Code `xml:"Cd"`

	// Type of organisation expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (o *OrganisationType1Choice) SetCode(value string) {
	o.Code = (*OrganisationType1Code)(&value)
}

func (o *OrganisationType1Choice) AddProprietary() *GenericIdentification47 {
	o.Proprietary = new(GenericIdentification47)
	return o.Proprietary
}
