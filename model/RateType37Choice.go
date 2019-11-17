package model

// Choice between a standard code or proprietary code to specify a rate type.
type RateType37Choice struct {

	// Standard code to specify the type of tax rate.
	Code *RateType3Code `xml:"Cd"`

	// Proprietary identification of the type of tax rate.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (r *RateType37Choice) SetCode(value string) {
	r.Code = (*RateType3Code)(&value)
}

func (r *RateType37Choice) AddProprietary() *GenericIdentification30 {
	r.Proprietary = new(GenericIdentification30)
	return r.Proprietary
}
