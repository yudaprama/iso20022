package model

// Sets of elements to identify a name of the identification scheme.
type PersonIdentificationSchemeName1Choice struct {

	// Name of the identification scheme, in a coded form as published in an external list.
	Code *ExternalPersonIdentification1Code `xml:"Cd"`

	// Name of the identification scheme, in a free text form.
	Proprietary *Max35Text `xml:"Prtry"`
}

func (p *PersonIdentificationSchemeName1Choice) SetCode(value string) {
	p.Code = (*ExternalPersonIdentification1Code)(&value)
}

func (p *PersonIdentificationSchemeName1Choice) SetProprietary(value string) {
	p.Proprietary = (*Max35Text)(&value)
}
