package model

// Choice of formats for the exemption reason.
type ExemptionReason1Choice struct {

	// Exemption reason expressed as a code.
	Code *TaxExemptReason1Code `xml:"Cd"`

	// Exemption reason expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (e *ExemptionReason1Choice) SetCode(value string) {
	e.Code = (*TaxExemptReason1Code)(&value)
}

func (e *ExemptionReason1Choice) AddProprietary() *GenericIdentification47 {
	e.Proprietary = new(GenericIdentification47)
	return e.Proprietary
}
