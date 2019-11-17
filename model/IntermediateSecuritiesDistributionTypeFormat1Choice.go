package model

// Choice between a standard code or proprietary code to specify the type of intermediate securities distribution.
type IntermediateSecuritiesDistributionTypeFormat1Choice struct {

	// Standard code to specify the type of intermediate security distribution.
	Code *IntermediateSecurityDistributionType3Code `xml:"Cd"`

	// Proprietary identification of the type of intermediate security distribution.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (i *IntermediateSecuritiesDistributionTypeFormat1Choice) SetCode(value string) {
	i.Code = (*IntermediateSecurityDistributionType3Code)(&value)
}

func (i *IntermediateSecuritiesDistributionTypeFormat1Choice) AddProprietary() *GenericIdentification20 {
	i.Proprietary = new(GenericIdentification20)
	return i.Proprietary
}
