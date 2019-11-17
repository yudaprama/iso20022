package model

// Choice of format for the settlement status.
type SettlementStatus1Choice struct {

	// Provides the status of settlement of an instruction/financial instrument movement.
	Code *SecuritiesSettlementStatus1Code `xml:"Cd"`

	// Provides the status of settlement of an instruction/financial instrument movement.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (s *SettlementStatus1Choice) SetCode(value string) {
	s.Code = (*SecuritiesSettlementStatus1Code)(&value)
}

func (s *SettlementStatus1Choice) AddProprietary() *GenericIdentification20 {
	s.Proprietary = new(GenericIdentification20)
	return s.Proprietary
}
