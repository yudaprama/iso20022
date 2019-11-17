package model

// Choice between a standard code or proprietary code to specify the information type format required.
type InformationTypeFormat4Choice struct {

	// Standard code to specify the information type required.
	Code *CorporateActionInformationType1Code `xml:"Cd"`

	// Proprietary identification of the information type.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (i *InformationTypeFormat4Choice) SetCode(value string) {
	i.Code = (*CorporateActionInformationType1Code)(&value)
}

func (i *InformationTypeFormat4Choice) AddProprietary() *GenericIdentification30 {
	i.Proprietary = new(GenericIdentification30)
	return i.Proprietary
}
