package model

// Choice of format for the settlement capacity information.
type SettlingCapacity4Choice struct {

	// Settlement capacity expressed as an ISO 20022 code.
	Code *SettlingCapacity2Code `xml:"Cd"`

	// Settlement capacity expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (s *SettlingCapacity4Choice) SetCode(value string) {
	s.Code = (*SettlingCapacity2Code)(&value)
}

func (s *SettlingCapacity4Choice) AddProprietary() *GenericIdentification20 {
	s.Proprietary = new(GenericIdentification20)
	return s.Proprietary
}
