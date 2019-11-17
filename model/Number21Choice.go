package model

// Choice number format.
type Number21Choice struct {

	// Number of maximum 4 numeric text.
	NumberIdentification *Max4NumericText `xml:"NbId"`

	// Number of maximum 35 text, with the possibility to provide an issuer for the number identification.
	Proprietary *GenericIdentification29 `xml:"Prtry"`
}

func (n *Number21Choice) SetNumberIdentification(value string) {
	n.NumberIdentification = (*Max4NumericText)(&value)
}

func (n *Number21Choice) AddProprietary() *GenericIdentification29 {
	n.Proprietary = new(GenericIdentification29)
	return n.Proprietary
}
