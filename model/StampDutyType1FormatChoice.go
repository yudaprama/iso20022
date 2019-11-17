package model

// Choice of formats to express the stamp duty information.
type StampDutyType1FormatChoice struct {

	// Standard code to specify the stamp duty information.
	Code *StampDutyType1Code `xml:"Cd"`

	// Proprietary code to express the stamp duty information.
	Proprietary *GenericIdentification13 `xml:"Prtry"`
}

func (s *StampDutyType1FormatChoice) SetCode(value string) {
	s.Code = (*StampDutyType1Code)(&value)
}

func (s *StampDutyType1FormatChoice) AddProprietary() *GenericIdentification13 {
	s.Proprietary = new(GenericIdentification13)
	return s.Proprietary
}
