package model

// Choice between a standard code or proprietary code to specify a rate type.
type RateType47Choice struct {

	// Standard code to specify the type of tax rate.
	Code *DividendRateType1Code `xml:"Cd"`

	// Proprietary identification of the type of tax rate.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (r *RateType47Choice) SetCode(value string) {
	r.Code = (*DividendRateType1Code)(&value)
}

func (r *RateType47Choice) AddProprietary() *GenericIdentification47 {
	r.Proprietary = new(GenericIdentification47)
	return r.Proprietary
}
