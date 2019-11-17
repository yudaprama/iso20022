package model

// Choice of formats for an accepted status reason code.
type AcceptedStatusReason1Choice struct {

	// Reason for the status expressed as a code.
	Code *AcceptedStatusReason1Code `xml:"Cd"`

	// Reason for the status expressed as a proprietary code.
	Proprietary *GenericIdentification36 `xml:"Prtry"`
}

func (a *AcceptedStatusReason1Choice) SetCode(value string) {
	a.Code = (*AcceptedStatusReason1Code)(&value)
}

func (a *AcceptedStatusReason1Choice) AddProprietary() *GenericIdentification36 {
	a.Proprietary = new(GenericIdentification36)
	return a.Proprietary
}
