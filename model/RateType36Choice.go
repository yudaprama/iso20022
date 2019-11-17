package model

// Choice between a standard code or proprietary code to specify a rate type.
type RateType36Choice struct {

	// Standard code to specify the type of tax rate.
	Code *DividendRateType1Code `xml:"Cd"`

	// Proprietary identification of the type of tax rate.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (r *RateType36Choice) SetCode(value string) {
	r.Code = (*DividendRateType1Code)(&value)
}

func (r *RateType36Choice) AddProprietary() *GenericIdentification30 {
	r.Proprietary = new(GenericIdentification30)
	return r.Proprietary
}
