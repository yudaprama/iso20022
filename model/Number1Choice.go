package model

// Choice of format for the number.
type Number1Choice struct {

	// Number of maximum 3 numeric text.
	NumberIdentification *Max3NumericText `xml:"NbId"`

	// Proprietary number format.
	Proprietary *GenericIdentification7 `xml:"Prtry"`
}

func (n *Number1Choice) SetNumberIdentification(value string) {
	n.NumberIdentification = (*Max3NumericText)(&value)
}

func (n *Number1Choice) AddProprietary() *GenericIdentification7 {
	n.Proprietary = new(GenericIdentification7)
	return n.Proprietary
}
