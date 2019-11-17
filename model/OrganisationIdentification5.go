package model

// Official identification of an organisation (legal entity) in a specific register.
type OrganisationIdentification5 struct {

	// Name of the register of legal entities.
	RegistrationNumber *Max35Text `xml:"RegnNb"`

	// Name of the register managed by a registration authority.
	RegisterName *Max35Text `xml:"RegrNm,omitempty"`
}

func (o *OrganisationIdentification5) SetRegistrationNumber(value string) {
	o.RegistrationNumber = (*Max35Text)(&value)
}

func (o *OrganisationIdentification5) SetRegisterName(value string) {
	o.RegisterName = (*Max35Text)(&value)
}
