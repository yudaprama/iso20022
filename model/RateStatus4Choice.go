package model

// Choice between a standard code or proprietary code to specify a rate status.
type RateStatus4Choice struct {

	// Standard code to specify the status of the rate.
	Code *RateStatus1Code `xml:"Cd"`

	// Proprietary identification of the status of the rate.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (r *RateStatus4Choice) SetCode(value string) {
	r.Code = (*RateStatus1Code)(&value)
}

func (r *RateStatus4Choice) AddProprietary() *GenericIdentification47 {
	r.Proprietary = new(GenericIdentification47)
	return r.Proprietary
}
