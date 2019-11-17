package model

// Choice between a standard code or proprietary code to specify a rate type.
type RateType7Choice struct {

	// Standard code to specify the type of net dividend rate.
	Code *NetDividendRateType1Code `xml:"Cd"`

	// Proprietary identification of the type of net dividend rate.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (r *RateType7Choice) SetCode(value string) {
	r.Code = (*NetDividendRateType1Code)(&value)
}

func (r *RateType7Choice) AddProprietary() *GenericIdentification20 {
	r.Proprietary = new(GenericIdentification20)
	return r.Proprietary
}
