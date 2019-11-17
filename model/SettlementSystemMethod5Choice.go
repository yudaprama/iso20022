package model

// Choice of format for the settlement system/method.
type SettlementSystemMethod5Choice struct {

	// Settlement system expressed as an ISO 20022 code.
	Code *SettlementSystemMethod1Code `xml:"Cd"`

	// Settlement system expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (s *SettlementSystemMethod5Choice) SetCode(value string) {
	s.Code = (*SettlementSystemMethod1Code)(&value)
}

func (s *SettlementSystemMethod5Choice) AddProprietary() *GenericIdentification47 {
	s.Proprietary = new(GenericIdentification47)
	return s.Proprietary
}
