package model

// Choice between a standard code or proprietary code to specify the type of intermediate securities distribution.
type IntermediateSecuritiesDistributionTypeFormat17Choice struct {

	// Standard code to specify the type of intermediate security distribution.
	Code *IntermediateSecurityDistributionType4Code `xml:"Cd"`

	// Proprietary identification of the type of intermediate security distribution.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (i *IntermediateSecuritiesDistributionTypeFormat17Choice) SetCode(value string) {
	i.Code = (*IntermediateSecurityDistributionType4Code)(&value)
}

func (i *IntermediateSecuritiesDistributionTypeFormat17Choice) AddProprietary() *GenericIdentification47 {
	i.Proprietary = new(GenericIdentification47)
	return i.Proprietary
}
