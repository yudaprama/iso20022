package model

// Choice of formats to  express the type of net dividend rate.
type NetDividendRateType1FormatChoice struct {

	// Standard code to  specify the type of net dividend rate.
	Code *NetDividendRateType1Code `xml:"Cd"`

	// Proprietary code to  express the type of net dividend rate.
	Proprietary *GenericIdentification13 `xml:"Prtry"`
}

func (n *NetDividendRateType1FormatChoice) SetCode(value string) {
	n.Code = (*NetDividendRateType1Code)(&value)
}

func (n *NetDividendRateType1FormatChoice) AddProprietary() *GenericIdentification13 {
	n.Proprietary = new(GenericIdentification13)
	return n.Proprietary
}
