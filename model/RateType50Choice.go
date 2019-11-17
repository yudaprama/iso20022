package model

// Choice between a standard code or proprietary code to specify a rate type.
type RateType50Choice struct {

	// Standard code to specify the type of net dividend rate.
	Code *NetDividendRateType2Code `xml:"Cd"`

	// Proprietary identification of the type of net dividend rate.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (r *RateType50Choice) SetCode(value string) {
	r.Code = (*NetDividendRateType2Code)(&value)
}

func (r *RateType50Choice) AddProprietary() *GenericIdentification47 {
	r.Proprietary = new(GenericIdentification47)
	return r.Proprietary
}
