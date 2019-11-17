package model

// Choice of formats for the specification of the type of political exposure.
type PoliticalExposureType1Choice struct {

	// Type of political exposure expressed as a code.
	Code *PoliticalExposureType1Code `xml:"Cd"`

	// Type of political exposure expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (p *PoliticalExposureType1Choice) SetCode(value string) {
	p.Code = (*PoliticalExposureType1Code)(&value)
}

func (p *PoliticalExposureType1Choice) AddProprietary() *GenericIdentification47 {
	p.Proprietary = new(GenericIdentification47)
	return p.Proprietary
}
