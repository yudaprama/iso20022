package model

// Choice of formats for the specification of the type  of insurance.
type InsuranceType1Choice struct {

	// Type of insurance contract expressed as a code.
	Code *Insurance1Code `xml:"Cd"`

	// Type of insurance contract expressed as a proprietary code.
	Proprietary *GenericIdentification41 `xml:"Prtry"`
}

func (i *InsuranceType1Choice) SetCode(value string) {
	i.Code = (*Insurance1Code)(&value)
}

func (i *InsuranceType1Choice) AddProprietary() *GenericIdentification41 {
	i.Proprietary = new(GenericIdentification41)
	return i.Proprietary
}
