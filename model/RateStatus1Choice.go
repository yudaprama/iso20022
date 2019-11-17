package model

// Choice between a standard code or proprietary code to specify a rate status.
type RateStatus1Choice struct {

	// Standard code to specify the status of the rate.
	Code *RateStatus1Code `xml:"Cd"`

	// Proprietary identification of the status of the rate.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (r *RateStatus1Choice) SetCode(value string) {
	r.Code = (*RateStatus1Code)(&value)
}

func (r *RateStatus1Choice) AddProprietary() *GenericIdentification20 {
	r.Proprietary = new(GenericIdentification20)
	return r.Proprietary
}
