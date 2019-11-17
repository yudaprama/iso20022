package model

// Choice between a standard code or proprietary code to specify a rate type.
type RateType17Choice struct {

	// Standard code to specify the type of tax rate.
	Code *DividendRateType1Code `xml:"Cd"`

	// Proprietary identification of the type of tax rate.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (r *RateType17Choice) SetCode(value string) {
	r.Code = (*DividendRateType1Code)(&value)
}

func (r *RateType17Choice) AddProprietary() *GenericIdentification20 {
	r.Proprietary = new(GenericIdentification20)
	return r.Proprietary
}
