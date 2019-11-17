package model

// Choice of format for the settlement capacity information.
type SettlingCapacity3Choice struct {

	// Settlement capacity expressed as an ISO 20022 code.
	Code *SettlingCapacity1Code `xml:"Cd"`

	// Settlement capacity expressed as a proprietary code.
	Proprietary *GenericIdentification38 `xml:"Prtry"`
}

func (s *SettlingCapacity3Choice) SetCode(value string) {
	s.Code = (*SettlingCapacity1Code)(&value)
}

func (s *SettlingCapacity3Choice) AddProprietary() *GenericIdentification38 {
	s.Proprietary = new(GenericIdentification38)
	return s.Proprietary
}
