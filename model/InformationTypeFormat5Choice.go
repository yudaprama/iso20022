package model

// Choice between a standard code or proprietary code to specify the information type format required.
type InformationTypeFormat5Choice struct {

	// Standard code to specify the information type required.
	Code *CorporateActionInformationType1Code `xml:"Cd"`

	// Proprietary identification of the information type.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (i *InformationTypeFormat5Choice) SetCode(value string) {
	i.Code = (*CorporateActionInformationType1Code)(&value)
}

func (i *InformationTypeFormat5Choice) AddProprietary() *GenericIdentification47 {
	i.Proprietary = new(GenericIdentification47)
	return i.Proprietary
}
