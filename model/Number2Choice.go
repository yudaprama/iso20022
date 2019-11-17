package model

// Choice number format.
type Number2Choice struct {

	// Number of maximum 3 numeric text.
	Short *Exact3NumericText `xml:"Shrt"`

	// Number of maximum 35 text, with the possibility to provide an issuer for the number identification.
	Long *GenericIdentification1 `xml:"Lng"`
}

func (n *Number2Choice) SetShort(value string) {
	n.Short = (*Exact3NumericText)(&value)
}

func (n *Number2Choice) AddLong() *GenericIdentification1 {
	n.Long = new(GenericIdentification1)
	return n.Long
}
