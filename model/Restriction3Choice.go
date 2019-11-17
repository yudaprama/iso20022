package model

// Choice of format for the securities restriction information.
type Restriction3Choice struct {

	// Restrictions expressed as an ISO 20022 code.
	Code *OwnershipLegalRestrictions1Code `xml:"Cd"`

	// Restrictions expressed as a proprietary code.
	Proprietary *GenericIdentification38 `xml:"Prtry"`
}

func (r *Restriction3Choice) SetCode(value string) {
	r.Code = (*OwnershipLegalRestrictions1Code)(&value)
}

func (r *Restriction3Choice) AddProprietary() *GenericIdentification38 {
	r.Proprietary = new(GenericIdentification38)
	return r.Proprietary
}
