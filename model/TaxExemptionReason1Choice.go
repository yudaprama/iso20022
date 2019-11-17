package model

// Choice of formats for the specification of the tax exemption reason.
type TaxExemptionReason1Choice struct {

	// Tax exemption reason expressed as a code.
	Code *TaxExemptReason1Code `xml:"Cd"`

	// Tax exemption reason expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (t *TaxExemptionReason1Choice) SetCode(value string) {
	t.Code = (*TaxExemptReason1Code)(&value)
}

func (t *TaxExemptionReason1Choice) AddProprietary() *GenericIdentification47 {
	t.Proprietary = new(GenericIdentification47)
	return t.Proprietary
}
