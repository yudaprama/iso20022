package model

// Choice between a standard code or proprietary code to specify a rate type.
type RateType38Choice struct {

	// Standard code to specify the type of gross dividend rate.
	Code *GrossDividendRateType2Code `xml:"Cd"`

	// Proprietary identification of the type of gross dividend rate.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (r *RateType38Choice) SetCode(value string) {
	r.Code = (*GrossDividendRateType2Code)(&value)
}

func (r *RateType38Choice) AddProprietary() *GenericIdentification30 {
	r.Proprietary = new(GenericIdentification30)
	return r.Proprietary
}
