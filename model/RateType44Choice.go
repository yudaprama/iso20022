package model

// Choice between a standard code or proprietary code to specify a rate type.
type RateType44Choice struct {

	// Standard code to specify the type of net dividend rate.
	Code *NetDividendRateType3Code `xml:"Cd"`

	// Proprietary identification of the type of net dividend rate.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (r *RateType44Choice) SetCode(value string) {
	r.Code = (*NetDividendRateType3Code)(&value)
}

func (r *RateType44Choice) AddProprietary() *GenericIdentification47 {
	r.Proprietary = new(GenericIdentification47)
	return r.Proprietary
}
