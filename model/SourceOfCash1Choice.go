package model

// Choice of formats for the source of the cash.
type SourceOfCash1Choice struct {

	// Source of cash expressed as a code.
	Code *SourceOfCash1Code `xml:"Cd"`

	// Source of cash expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (s *SourceOfCash1Choice) SetCode(value string) {
	s.Code = (*SourceOfCash1Code)(&value)
}

func (s *SourceOfCash1Choice) AddProprietary() *GenericIdentification47 {
	s.Proprietary = new(GenericIdentification47)
	return s.Proprietary
}
