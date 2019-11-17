package model

// Choice between a standard code or proprietary code to specify a rate type.
type RateType28Choice struct {

	// Standard code to specify the type of withholding tax rate.
	Code *WithholdingTaxRateType1Code `xml:"Cd"`

	// Proprietary identification of the type of withholding tax rate.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (r *RateType28Choice) SetCode(value string) {
	r.Code = (*WithholdingTaxRateType1Code)(&value)
}

func (r *RateType28Choice) AddProprietary() *GenericIdentification20 {
	r.Proprietary = new(GenericIdentification20)
	return r.Proprietary
}
