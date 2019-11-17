package model

// Choice of processing change command status.
type SettlementConditionModificationStatus1Choice struct {

	// Provides the status of the securities settlement condition modification request.
	Code *SettlementConditionModificationStatus1Code `xml:"Cd"`

	// Provides the status of the securities settlement condition modification request.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (s *SettlementConditionModificationStatus1Choice) SetCode(value string) {
	s.Code = (*SettlementConditionModificationStatus1Code)(&value)
}

func (s *SettlementConditionModificationStatus1Choice) AddProprietary() *GenericIdentification20 {
	s.Proprietary = new(GenericIdentification20)
	return s.Proprietary
}
