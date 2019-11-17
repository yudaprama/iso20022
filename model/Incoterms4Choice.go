package model

// Specifies the applicable Incoterm and associated location.
type Incoterms4Choice struct {

	// Specifies the applicable Incoterm by means of a code.
	Code *ExternalIncoterms1Code `xml:"Cd"`

	// Specifies the applicable Incoterm by means of a proprietary scheme.
	Proprietary *GenericIdentification13 `xml:"Prtry"`
}

func (i *Incoterms4Choice) SetCode(value string) {
	i.Code = (*ExternalIncoterms1Code)(&value)
}

func (i *Incoterms4Choice) AddProprietary() *GenericIdentification13 {
	i.Proprietary = new(GenericIdentification13)
	return i.Proprietary
}
