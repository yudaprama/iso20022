package model

// Choice of formats for a suspended status.
type SuspendedStatusReason5Choice struct {

	// Suspended reason expressed as a code.
	Code *SuspendedStatusReason3Code `xml:"Cd"`

	// Suspended reason expressed as a proprietary code.
	Proprietary *GenericIdentification1 `xml:"Prtry"`
}

func (s *SuspendedStatusReason5Choice) SetCode(value string) {
	s.Code = (*SuspendedStatusReason3Code)(&value)
}

func (s *SuspendedStatusReason5Choice) AddProprietary() *GenericIdentification1 {
	s.Proprietary = new(GenericIdentification1)
	return s.Proprietary
}
