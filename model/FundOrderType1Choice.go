package model

// Choice of formats for the specification of the fund order type.
type FundOrderType1Choice struct {

	// Order type expressed as a code.
	Type *FundOrderType6Code `xml:"Tp"`

	// Order type expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (f *FundOrderType1Choice) SetType(value string) {
	f.Type = (*FundOrderType6Code)(&value)
}

func (f *FundOrderType1Choice) AddProprietary() *GenericIdentification47 {
	f.Proprietary = new(GenericIdentification47)
	return f.Proprietary
}
