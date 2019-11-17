package model

// Choice of formats for the specification of how information is to be distributed.
type InformationDistribution1Choice struct {

	// Information distribution expressed as a code.
	Code *InformationDistribution2Code `xml:"Cd"`

	// Information distribution expressed as a code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (i *InformationDistribution1Choice) SetCode(value string) {
	i.Code = (*InformationDistribution2Code)(&value)
}

func (i *InformationDistribution1Choice) AddProprietary() *GenericIdentification47 {
	i.Proprietary = new(GenericIdentification47)
	return i.Proprietary
}
