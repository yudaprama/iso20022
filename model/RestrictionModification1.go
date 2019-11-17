package model

// Specifies the type of change to a restriction.
type RestrictionModification1 struct {

	// Specifies the type of change.
	ModificationCode *Modification1Code `xml:"ModCd,omitempty"`

	// Restriction.
	Restriction *Restriction1 `xml:"Rstrctn"`
}

func (r *RestrictionModification1) SetModificationCode(value string) {
	r.ModificationCode = (*Modification1Code)(&value)
}

func (r *RestrictionModification1) AddRestriction() *Restriction1 {
	r.Restriction = new(Restriction1)
	return r.Restriction
}
