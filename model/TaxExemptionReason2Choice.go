package model

// Choice of formats for the specification of the tax exemption reason.
type TaxExemptionReason2Choice struct {

	// Tax exemption reason expressed as a code.
	Code *TaxExemptReason3Code `xml:"Cd"`

	// Tax exemption reason expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (t *TaxExemptionReason2Choice) SetCode(value string) {
	t.Code = (*TaxExemptReason3Code)(&value)
}

func (t *TaxExemptionReason2Choice) AddProprietary() *GenericIdentification47 {
	t.Proprietary = new(GenericIdentification47)
	return t.Proprietary
}
