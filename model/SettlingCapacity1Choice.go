package model

// Choice of format for the settlement capacity information.
type SettlingCapacity1Choice struct {

	// Settlement capacity expressed as an ISO 20022 code.
	Code *SettlingCapacity1Code `xml:"Cd"`

	// Settlement capacity expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (s *SettlingCapacity1Choice) SetCode(value string) {
	s.Code = (*SettlingCapacity1Code)(&value)
}

func (s *SettlingCapacity1Choice) AddProprietary() *GenericIdentification20 {
	s.Proprietary = new(GenericIdentification20)
	return s.Proprietary
}
