package model

// Choice of format for the type of party.
type PartyType1Choice struct {

	// Type of party.
	//
	Code *ExternalTypeOfParty1Code `xml:"Cd"`

	// Type of party expressed as a proprietary code.
	Proprietary *GenericIdentification1 `xml:"Prtry"`
}

func (p *PartyType1Choice) SetCode(value string) {
	p.Code = (*ExternalTypeOfParty1Code)(&value)
}

func (p *PartyType1Choice) AddProprietary() *GenericIdentification1 {
	p.Proprietary = new(GenericIdentification1)
	return p.Proprietary
}
