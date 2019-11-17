package model

// Choice of formats for the order type.
type FundOrderType4Choice struct {

	// Type of the investment fund order expressed as a code.
	Code *FundOrderType8Code `xml:"Cd"`

	// Type of the investment fund order expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (f *FundOrderType4Choice) SetCode(value string) {
	f.Code = (*FundOrderType8Code)(&value)
}

func (f *FundOrderType4Choice) AddProprietary() *GenericIdentification47 {
	f.Proprietary = new(GenericIdentification47)
	return f.Proprietary
}
