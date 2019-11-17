package model

// Choice of formats for the specification of the type  of insurance.
type InsuranceType2Choice struct {

	// Type of insurance contract expressed as a code.
	Code *Insurance1Code `xml:"Cd"`

	// Type of insurance contract expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (i *InsuranceType2Choice) SetCode(value string) {
	i.Code = (*Insurance1Code)(&value)
}

func (i *InsuranceType2Choice) AddProprietary() *GenericIdentification47 {
	i.Proprietary = new(GenericIdentification47)
	return i.Proprietary
}
