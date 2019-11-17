package model

// Choice of format for the securities lending type.
type SecuritiesLendingType1Choice struct {

	// Securities lending type expressed as an ISO 20022 code.
	Code *SecuritiesLendingType1Code `xml:"Cd"`

	// Securities lending type expressed as a proprietary code.
	Proprietary *GenericIdentification38 `xml:"Prtry"`
}

func (s *SecuritiesLendingType1Choice) SetCode(value string) {
	s.Code = (*SecuritiesLendingType1Code)(&value)
}

func (s *SecuritiesLendingType1Choice) AddProprietary() *GenericIdentification38 {
	s.Proprietary = new(GenericIdentification38)
	return s.Proprietary
}
