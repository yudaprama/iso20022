package model

// Choice between a standard code or proprietary code to specify a rate type.
type RateType23Choice struct {

	// Standard code to specify the type of net dividend rate.
	Code *NetDividendRateType3Code `xml:"Cd"`

	// Proprietary identification of the type of net dividend rate.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (r *RateType23Choice) SetCode(value string) {
	r.Code = (*NetDividendRateType3Code)(&value)
}

func (r *RateType23Choice) AddProprietary() *GenericIdentification20 {
	r.Proprietary = new(GenericIdentification20)
	return r.Proprietary
}
