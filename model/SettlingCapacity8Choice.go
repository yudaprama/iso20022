package model

// Choice of format for the settlement capacity information.
type SettlingCapacity8Choice struct {

	// Settlement capacity expressed as an ISO 20022 code.
	Code *SettlingCapacity2Code `xml:"Cd"`

	// Settlement capacity expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (s *SettlingCapacity8Choice) SetCode(value string) {
	s.Code = (*SettlingCapacity2Code)(&value)
}

func (s *SettlingCapacity8Choice) AddProprietary() *GenericIdentification47 {
	s.Proprietary = new(GenericIdentification47)
	return s.Proprietary
}
