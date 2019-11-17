package model

// Choice between a standard code or proprietary code to specify a rate type.
type RateType39Choice struct {

	// Standard code to specify the type of net dividend rate.
	Code *NetDividendRateType2Code `xml:"Cd"`

	// Proprietary identification of the type of net dividend rate.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (r *RateType39Choice) SetCode(value string) {
	r.Code = (*NetDividendRateType2Code)(&value)
}

func (r *RateType39Choice) AddProprietary() *GenericIdentification30 {
	r.Proprietary = new(GenericIdentification30)
	return r.Proprietary
}
