package model

// Choice of formats for the restriction or limitation status.
type RestrictionStatus1Choice struct {

	// Status of the restriction expressed as a code.
	Code *RestrictionStatus1Code `xml:"Cd"`

	// Status of the restriction expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (r *RestrictionStatus1Choice) SetCode(value string) {
	r.Code = (*RestrictionStatus1Code)(&value)
}

func (r *RestrictionStatus1Choice) AddProprietary() *GenericIdentification47 {
	r.Proprietary = new(GenericIdentification47)
	return r.Proprietary
}
