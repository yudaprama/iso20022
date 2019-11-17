package model

// Choice of processing change command status.
type SettlementConditionModificationStatus4Choice struct {

	// Provides the status of the securities settlement condition modification request.
	Code *SettlementConditionModificationStatus1Code `xml:"Cd"`

	// Provides the status of the securities settlement condition modification request.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (s *SettlementConditionModificationStatus4Choice) SetCode(value string) {
	s.Code = (*SettlementConditionModificationStatus1Code)(&value)
}

func (s *SettlementConditionModificationStatus4Choice) AddProprietary() *GenericIdentification47 {
	s.Proprietary = new(GenericIdentification47)
	return s.Proprietary
}
