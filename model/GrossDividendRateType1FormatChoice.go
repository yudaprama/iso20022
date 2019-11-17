package model

// Choice of formats to  express the type of gross dividend rate.
type GrossDividendRateType1FormatChoice struct {

	// Standard code to  specify the type of gross dividend rate.
	Code *GrossDividendRateType1Code `xml:"Cd"`

	// Proprietary code to  express the type of gross dividend rate.
	Proprietary *GenericIdentification13 `xml:"Prtry"`
}

func (g *GrossDividendRateType1FormatChoice) SetCode(value string) {
	g.Code = (*GrossDividendRateType1Code)(&value)
}

func (g *GrossDividendRateType1FormatChoice) AddProprietary() *GenericIdentification13 {
	g.Proprietary = new(GenericIdentification13)
	return g.Proprietary
}
