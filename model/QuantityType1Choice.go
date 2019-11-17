package model

// Choice of formats for the quantity.
type QuantityType1Choice struct {

	// Quantity type expressed as a code.
	Code *OrderQuantityType2Code `xml:"Cd"`

	// Quantity type expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (q *QuantityType1Choice) SetCode(value string) {
	q.Code = (*OrderQuantityType2Code)(&value)
}

func (q *QuantityType1Choice) AddProprietary() *GenericIdentification47 {
	q.Proprietary = new(GenericIdentification47)
	return q.Proprietary
}
