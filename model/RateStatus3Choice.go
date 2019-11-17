package model

// Choice between a standard code or proprietary code to specify a rate status.
type RateStatus3Choice struct {

	// Standard code to specify the status of the rate.
	Code *RateStatus1Code `xml:"Cd"`

	// Proprietary identification of the status of the rate.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (r *RateStatus3Choice) SetCode(value string) {
	r.Code = (*RateStatus1Code)(&value)
}

func (r *RateStatus3Choice) AddProprietary() *GenericIdentification30 {
	r.Proprietary = new(GenericIdentification30)
	return r.Proprietary
}
