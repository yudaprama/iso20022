package model

// Choice between a standard code or proprietary code to specify a rate type.
type RateType40Choice struct {

	// Standard code to specify the type of gross dividend rate.
	Code *GrossDividendRateType3Code `xml:"Cd"`

	// Proprietary identification of the type of gross dividend rate.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (r *RateType40Choice) SetCode(value string) {
	r.Code = (*GrossDividendRateType3Code)(&value)
}

func (r *RateType40Choice) AddProprietary() *GenericIdentification30 {
	r.Proprietary = new(GenericIdentification30)
	return r.Proprietary
}
