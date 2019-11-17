package model

// Choice between a standard code or proprietary code to specify a rate type.
type RateType10Choice struct {

	// Standard code to specify the type of tax rate.
	Code *RateType3Code `xml:"Cd"`

	// Proprietary identification of the type of tax rate.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (r *RateType10Choice) SetCode(value string) {
	r.Code = (*RateType3Code)(&value)
}

func (r *RateType10Choice) AddProprietary() *GenericIdentification20 {
	r.Proprietary = new(GenericIdentification20)
	return r.Proprietary
}
