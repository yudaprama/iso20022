package model

// Choice of format for the form of securities.
type FormOfSecurity7Choice struct {

	// Form of the security expressed as an ISO 20022 code.
	Code *FormOfSecurity1Code `xml:"Cd"`

	// Form of the security expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (f *FormOfSecurity7Choice) SetCode(value string) {
	f.Code = (*FormOfSecurity1Code)(&value)
}

func (f *FormOfSecurity7Choice) AddProprietary() *GenericIdentification47 {
	f.Proprietary = new(GenericIdentification47)
	return f.Proprietary
}
