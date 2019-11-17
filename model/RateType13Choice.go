package model

// Choice between a standard code or proprietary code to specify a rate type.
type RateType13Choice struct {

	// Standard code to specify the type of gross dividend rate.
	Code *GrossDividendRateType1Code `xml:"Cd"`

	// Proprietary identification of the type of gross dividend rate.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (r *RateType13Choice) SetCode(value string) {
	r.Code = (*GrossDividendRateType1Code)(&value)
}

func (r *RateType13Choice) AddProprietary() *GenericIdentification20 {
	r.Proprietary = new(GenericIdentification20)
	return r.Proprietary
}
