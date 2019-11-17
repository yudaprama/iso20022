package model

// Choice of format for the securities restriction information.
type Restriction6Choice struct {

	// Restrictions expressed as an ISO 20022 code.
	Code *OwnershipLegalRestrictions1Code `xml:"Cd"`

	// Restrictions expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (r *Restriction6Choice) SetCode(value string) {
	r.Code = (*OwnershipLegalRestrictions1Code)(&value)
}

func (r *Restriction6Choice) AddProprietary() *GenericIdentification47 {
	r.Proprietary = new(GenericIdentification47)
	return r.Proprietary
}
