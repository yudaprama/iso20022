package model

// Choice between a standard code or proprietary code to specify the type of intermediate securities distribution.
type IntermediateSecuritiesDistributionTypeFormat18Choice struct {

	// Standard code to specify the type of intermediate security distribution.
	Code *IntermediateSecurityDistributionType5Code `xml:"Cd"`

	// Proprietary identification of the type of intermediate security distribution.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (i *IntermediateSecuritiesDistributionTypeFormat18Choice) SetCode(value string) {
	i.Code = (*IntermediateSecurityDistributionType5Code)(&value)
}

func (i *IntermediateSecuritiesDistributionTypeFormat18Choice) AddProprietary() *GenericIdentification47 {
	i.Proprietary = new(GenericIdentification47)
	return i.Proprietary
}
