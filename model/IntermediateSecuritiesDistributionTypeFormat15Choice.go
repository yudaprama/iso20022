package model

// Choice between a standard code or proprietary code to specify the type of intermediate securities distribution.
type IntermediateSecuritiesDistributionTypeFormat15Choice struct {

	// Standard code to specify the type of intermediate security distribution.
	Code *IntermediateSecurityDistributionType5Code `xml:"Cd"`

	// Proprietary identification of the type of intermediate security distribution.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (i *IntermediateSecuritiesDistributionTypeFormat15Choice) SetCode(value string) {
	i.Code = (*IntermediateSecurityDistributionType5Code)(&value)
}

func (i *IntermediateSecuritiesDistributionTypeFormat15Choice) AddProprietary() *GenericIdentification30 {
	i.Proprietary = new(GenericIdentification30)
	return i.Proprietary
}
