package model

// Choice of formats to  express the type of intermediate security distribution.
type IntermediateSecurityDistributionType1FormatChoice struct {

	// Standard code to  specify the type of intermediate security distribution.
	Code *IntermediateSecurityDistributionType1Code `xml:"Cd"`

	// Proprietary code to  express the type of intermediate security distribution.
	Proprietary *GenericIdentification13 `xml:"Prtry"`
}

func (i *IntermediateSecurityDistributionType1FormatChoice) SetCode(value string) {
	i.Code = (*IntermediateSecurityDistributionType1Code)(&value)
}

func (i *IntermediateSecurityDistributionType1FormatChoice) AddProprietary() *GenericIdentification13 {
	i.Proprietary = new(GenericIdentification13)
	return i.Proprietary
}
