package model

// Choice between a standard code or proprietary code to specify a rate type.
type RateType22Choice struct {

	// Standard code to specify the type of gross dividend rate.
	Code *GrossDividendRateType3Code `xml:"Cd"`

	// Proprietary identification of the type of gross dividend rate.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (r *RateType22Choice) SetCode(value string) {
	r.Code = (*GrossDividendRateType3Code)(&value)
}

func (r *RateType22Choice) AddProprietary() *GenericIdentification20 {
	r.Proprietary = new(GenericIdentification20)
	return r.Proprietary
}
