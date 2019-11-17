package model

// Restriction on capability or operations allowed.
type Restriction1 struct {

	// Type of the restriction.
	RestrictionType *CodeOrProprietary1Choice `xml:"RstrctnTp"`

	// Date from when the restriction is valid.
	ValidFrom *ISODateTime `xml:"VldFr"`

	// Date until when the restriction is valid.
	ValidUntil *ISODateTime `xml:"VldUntil,omitempty"`
}

func (r *Restriction1) AddRestrictionType() *CodeOrProprietary1Choice {
	r.RestrictionType = new(CodeOrProprietary1Choice)
	return r.RestrictionType
}

func (r *Restriction1) SetValidFrom(value string) {
	r.ValidFrom = (*ISODateTime)(&value)
}

func (r *Restriction1) SetValidUntil(value string) {
	r.ValidUntil = (*ISODateTime)(&value)
}
