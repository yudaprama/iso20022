package model

// Choice of format for the settlement system/method.
type SettlementSystemMethod4Choice struct {

	// Settlement system expressed as an ISO 20022 code.
	Code *SettlementSystemMethod1Code `xml:"Cd"`

	// Settlement system expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (s *SettlementSystemMethod4Choice) SetCode(value string) {
	s.Code = (*SettlementSystemMethod1Code)(&value)
}

func (s *SettlementSystemMethod4Choice) AddProprietary() *GenericIdentification30 {
	s.Proprietary = new(GenericIdentification30)
	return s.Proprietary
}
