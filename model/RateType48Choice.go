package model

// Choice between a standard code or proprietary code to specify a rate type.
type RateType48Choice struct {

	// Standard code to specify the type of tax rate.
	Code *RateType3Code `xml:"Cd"`

	// Proprietary identification of the type of tax rate.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (r *RateType48Choice) SetCode(value string) {
	r.Code = (*RateType3Code)(&value)
}

func (r *RateType48Choice) AddProprietary() *GenericIdentification47 {
	r.Proprietary = new(GenericIdentification47)
	return r.Proprietary
}
