package model

// Choice between a standard code or proprietary code to specify the type of intermediate securities distribution.
type IntermediateSecuritiesDistributionTypeFormat2Choice struct {

	// Standard code to specify the type of intermediate security distribution.
	Code *IntermediateSecurityDistributionType2Code `xml:"Cd"`

	// Proprietary identification of the type of intermediate security distribution.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (i *IntermediateSecuritiesDistributionTypeFormat2Choice) SetCode(value string) {
	i.Code = (*IntermediateSecurityDistributionType2Code)(&value)
}

func (i *IntermediateSecuritiesDistributionTypeFormat2Choice) AddProprietary() *GenericIdentification20 {
	i.Proprietary = new(GenericIdentification20)
	return i.Proprietary
}
