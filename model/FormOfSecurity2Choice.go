package model

// Choice of format for the form of securities.
type FormOfSecurity2Choice struct {

	// Form of the security expressed as an ISO 20022 code.
	Code *FormOfSecurity1Code `xml:"Cd"`

	// Form of the security expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (f *FormOfSecurity2Choice) SetCode(value string) {
	f.Code = (*FormOfSecurity1Code)(&value)
}

func (f *FormOfSecurity2Choice) AddProprietary() *GenericIdentification20 {
	f.Proprietary = new(GenericIdentification20)
	return f.Proprietary
}
