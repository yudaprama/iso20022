package model

// Choice between a standard code or proprietary code to specify the information type format required.
type InformationTypeFormat1Choice struct {

	// Standard code to specify the information type required.
	Code *CorporateActionInformationType1Code `xml:"Cd"`

	// Proprietary identification of the information type.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (i *InformationTypeFormat1Choice) SetCode(value string) {
	i.Code = (*CorporateActionInformationType1Code)(&value)
}

func (i *InformationTypeFormat1Choice) AddProprietary() *GenericIdentification20 {
	i.Proprietary = new(GenericIdentification20)
	return i.Proprietary
}
