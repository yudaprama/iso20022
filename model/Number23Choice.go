package model

// Choice number format.
type Number23Choice struct {

	// Number of maximum 3 numeric text.
	Short *Exact3NumericText `xml:"Shrt"`

	// Number of maximum 35 text, with the possibility to provide an issuer for the number identification.
	Long *GenericIdentification18 `xml:"Lng"`
}

func (n *Number23Choice) SetShort(value string) {
	n.Short = (*Exact3NumericText)(&value)
}

func (n *Number23Choice) AddLong() *GenericIdentification18 {
	n.Long = new(GenericIdentification18)
	return n.Long
}
