package model

// Choice of processing change command status.
type SettlementConditionModificationStatus3Choice struct {

	// Provides the status of the securities settlement condition modification request.
	Code *SettlementConditionModificationStatus1Code `xml:"Cd"`

	// Provides the status of the securities settlement condition modification request.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (s *SettlementConditionModificationStatus3Choice) SetCode(value string) {
	s.Code = (*SettlementConditionModificationStatus1Code)(&value)
}

func (s *SettlementConditionModificationStatus3Choice) AddProprietary() *GenericIdentification30 {
	s.Proprietary = new(GenericIdentification30)
	return s.Proprietary
}
