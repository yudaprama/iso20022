package model

// Choice of format for the form of securities.
type FormOfSecurity6Choice struct {

	// Form of the security expressed as an ISO 20022 code.
	Code *FormOfSecurity1Code `xml:"Cd"`

	// Form of the security expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (f *FormOfSecurity6Choice) SetCode(value string) {
	f.Code = (*FormOfSecurity1Code)(&value)
}

func (f *FormOfSecurity6Choice) AddProprietary() *GenericIdentification30 {
	f.Proprietary = new(GenericIdentification30)
	return f.Proprietary
}
