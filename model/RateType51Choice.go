package model

// Choice between a standard code or proprietary code to specify a rate type.
type RateType51Choice struct {

	// Standard code to specify the type of gross dividend rate.
	Code *GrossDividendRateType3Code `xml:"Cd"`

	// Proprietary identification of the type of gross dividend rate.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (r *RateType51Choice) SetCode(value string) {
	r.Code = (*GrossDividendRateType3Code)(&value)
}

func (r *RateType51Choice) AddProprietary() *GenericIdentification47 {
	r.Proprietary = new(GenericIdentification47)
	return r.Proprietary
}
