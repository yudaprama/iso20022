package model

// Choice of format for the securities restriction information.
type Restriction1Choice struct {

	// Restrictions expressed as an ISO 20022 code.
	Code *OwnershipLegalRestrictions1Code `xml:"Cd"`

	// Restrictions expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (r *Restriction1Choice) SetCode(value string) {
	r.Code = (*OwnershipLegalRestrictions1Code)(&value)
}

func (r *Restriction1Choice) AddProprietary() *GenericIdentification20 {
	r.Proprietary = new(GenericIdentification20)
	return r.Proprietary
}
