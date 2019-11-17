package model

// Choice between a standard code or proprietary code to specify a rate type.
type RateType46Choice struct {

	// Standard code to specify the type of withholding tax rate.
	Code *WithholdingTaxRateType1Code `xml:"Cd"`

	// Proprietary identification of the type of withholding tax rate.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (r *RateType46Choice) SetCode(value string) {
	r.Code = (*WithholdingTaxRateType1Code)(&value)
}

func (r *RateType46Choice) AddProprietary() *GenericIdentification47 {
	r.Proprietary = new(GenericIdentification47)
	return r.Proprietary
}
