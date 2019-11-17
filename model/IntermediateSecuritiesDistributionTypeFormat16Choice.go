package model

// Choice between a standard code or proprietary code to specify the type of intermediate securities distribution.
type IntermediateSecuritiesDistributionTypeFormat16Choice struct {

	// Standard code to specify the type of intermediate security distribution.
	Code *IntermediateSecurityDistributionType4Code `xml:"Cd"`

	// Proprietary identification of the type of intermediate security distribution.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (i *IntermediateSecuritiesDistributionTypeFormat16Choice) SetCode(value string) {
	i.Code = (*IntermediateSecurityDistributionType4Code)(&value)
}

func (i *IntermediateSecuritiesDistributionTypeFormat16Choice) AddProprietary() *GenericIdentification30 {
	i.Proprietary = new(GenericIdentification30)
	return i.Proprietary
}
