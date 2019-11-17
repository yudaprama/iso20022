package model

// Choice of formats for order breakdown type.
type OrderBreakdownType1Choice struct {

	// Order breakdown type expressed as a code.
	Code *FundOrderType5Code `xml:"Cd"`

	// Order breakdown type expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (o *OrderBreakdownType1Choice) SetCode(value string) {
	o.Code = (*FundOrderType5Code)(&value)
}

func (o *OrderBreakdownType1Choice) AddProprietary() *GenericIdentification47 {
	o.Proprietary = new(GenericIdentification47)
	return o.Proprietary
}
