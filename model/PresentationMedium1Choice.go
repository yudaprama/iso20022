package model

// Choice of format for the presentation medium.
type PresentationMedium1Choice struct {

	// Presentation medium.
	//
	Code *PresentationMedium1Code `xml:"Cd"`

	// Presentation medium expressed as a proprietary code.
	Proprietary *GenericIdentification1 `xml:"Prtry"`
}

func (p *PresentationMedium1Choice) SetCode(value string) {
	p.Code = (*PresentationMedium1Code)(&value)
}

func (p *PresentationMedium1Choice) AddProprietary() *GenericIdentification1 {
	p.Proprietary = new(GenericIdentification1)
	return p.Proprietary
}
