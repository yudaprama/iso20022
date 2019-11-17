package model

// Choice of format for the form of securities.
type FormOfSecurity4Choice struct {

	// Form of the security expressed as an ISO 20022 code.
	Code *FormOfSecurity1Code `xml:"Cd"`

	// Form of the security expressed as a proprietary code.
	Proprietary *GenericIdentification38 `xml:"Prtry"`
}

func (f *FormOfSecurity4Choice) SetCode(value string) {
	f.Code = (*FormOfSecurity1Code)(&value)
}

func (f *FormOfSecurity4Choice) AddProprietary() *GenericIdentification38 {
	f.Proprietary = new(GenericIdentification38)
	return f.Proprietary
}
