package model

// Choice between a standard code or proprietary code to specify a rate type.
type RateType45Choice struct {

	// Standard code to specify the type of rate.
	Code *RateType7Code `xml:"Cd"`

	// Proprietary identification of the type of rate.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (r *RateType45Choice) SetCode(value string) {
	r.Code = (*RateType7Code)(&value)
}

func (r *RateType45Choice) AddProprietary() *GenericIdentification47 {
	r.Proprietary = new(GenericIdentification47)
	return r.Proprietary
}
