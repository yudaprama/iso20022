package model

// Choice between a standard code or proprietary code to specify a rate type.
type RateType33Choice struct {

	// Standard code to specify the type of rate.
	Code *RateType7Code `xml:"Cd"`

	// Proprietary identification of the type of rate.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (r *RateType33Choice) SetCode(value string) {
	r.Code = (*RateType7Code)(&value)
}

func (r *RateType33Choice) AddProprietary() *GenericIdentification30 {
	r.Proprietary = new(GenericIdentification30)
	return r.Proprietary
}
