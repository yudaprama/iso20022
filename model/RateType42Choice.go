package model

// Choice between a standard code or proprietary code to specify a rate type.
type RateType42Choice struct {

	// Standard code to specify the type of withholding tax rate.
	Code *WithholdingTaxRateType1Code `xml:"Cd"`

	// Proprietary identification of the type of withholding tax rate.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (r *RateType42Choice) SetCode(value string) {
	r.Code = (*WithholdingTaxRateType1Code)(&value)
}

func (r *RateType42Choice) AddProprietary() *GenericIdentification30 {
	r.Proprietary = new(GenericIdentification30)
	return r.Proprietary
}
