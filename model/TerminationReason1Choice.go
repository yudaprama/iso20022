package model

// Choice of format for the undertaking termination reason.
type TerminationReason1Choice struct {

	// Termination reason.
	Code *TerminationReason1Code `xml:"Cd"`

	// Termination reason expressed as a proprietary code.
	Proprietary *GenericIdentification1 `xml:"Prtry"`
}

func (t *TerminationReason1Choice) SetCode(value string) {
	t.Code = (*TerminationReason1Code)(&value)
}

func (t *TerminationReason1Choice) AddProprietary() *GenericIdentification1 {
	t.Proprietary = new(GenericIdentification1)
	return t.Proprietary
}
