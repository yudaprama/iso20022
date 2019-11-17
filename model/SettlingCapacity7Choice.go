package model

// Choice of format for the settlement capacity information.
type SettlingCapacity7Choice struct {

	// Settlement capacity expressed as an ISO 20022 code.
	Code *SettlingCapacity2Code `xml:"Cd"`

	// Settlement capacity expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (s *SettlingCapacity7Choice) SetCode(value string) {
	s.Code = (*SettlingCapacity2Code)(&value)
}

func (s *SettlingCapacity7Choice) AddProprietary() *GenericIdentification30 {
	s.Proprietary = new(GenericIdentification30)
	return s.Proprietary
}
