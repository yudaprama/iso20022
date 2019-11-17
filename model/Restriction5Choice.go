package model

// Choice of format for the securities restriction information.
type Restriction5Choice struct {

	// Restrictions expressed as an ISO 20022 code.
	Code *OwnershipLegalRestrictions1Code `xml:"Cd"`

	// Restrictions expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (r *Restriction5Choice) SetCode(value string) {
	r.Code = (*OwnershipLegalRestrictions1Code)(&value)
}

func (r *Restriction5Choice) AddProprietary() *GenericIdentification30 {
	r.Proprietary = new(GenericIdentification30)
	return r.Proprietary
}
