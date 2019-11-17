package model

// Choice of formats for the specification of the  order type.
type OrderType2Choice struct {

	// Order type expressed as a code.
	Type *FundOrderType7Code `xml:"Tp"`

	// Order type expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (o *OrderType2Choice) SetType(value string) {
	o.Type = (*FundOrderType7Code)(&value)
}

func (o *OrderType2Choice) AddProprietary() *GenericIdentification47 {
	o.Proprietary = new(GenericIdentification47)
	return o.Proprietary
}
