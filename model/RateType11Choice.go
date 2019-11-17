package model

// Choice between a standard code or proprietary code to specify a rate type.
type RateType11Choice struct {

	// Standard code to specify the type of tax rate.
	Code *TaxType4Code `xml:"Cd"`

	// Proprietary identification of the type of tax rate.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (r *RateType11Choice) SetCode(value string) {
	r.Code = (*TaxType4Code)(&value)
}

func (r *RateType11Choice) AddProprietary() *GenericIdentification20 {
	r.Proprietary = new(GenericIdentification20)
	return r.Proprietary
}
