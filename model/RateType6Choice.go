package model

// Choice between a standard code or proprietary code to specify a rate type.
type RateType6Choice struct {

	// Standard code to specify the type of rate.
	Code *RateType7Code `xml:"Cd"`

	// Proprietary identification of the type of rate.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (r *RateType6Choice) SetCode(value string) {
	r.Code = (*RateType7Code)(&value)
}

func (r *RateType6Choice) AddProprietary() *GenericIdentification20 {
	r.Proprietary = new(GenericIdentification20)
	return r.Proprietary
}
