package model

// Choice of formats for the profile type.
type ProfileType1Choice struct {

	// Type of profile expressed as a code.
	Code *ProfileType1Code `xml:"Cd"`

	// Type of profile expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (p *ProfileType1Choice) SetCode(value string) {
	p.Code = (*ProfileType1Code)(&value)
}

func (p *ProfileType1Choice) AddProprietary() *GenericIdentification47 {
	p.Proprietary = new(GenericIdentification47)
	return p.Proprietary
}
