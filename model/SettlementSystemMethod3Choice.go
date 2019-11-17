package model

// Choice of format for the settlement system/method.
type SettlementSystemMethod3Choice struct {

	// Settlement system expressed as an ISO 20022 code.
	Code *SettlementSystemMethod1Code `xml:"Cd"`

	// Settlement system expressed as a proprietary code.
	Proprietary *GenericIdentification38 `xml:"Prtry"`
}

func (s *SettlementSystemMethod3Choice) SetCode(value string) {
	s.Code = (*SettlementSystemMethod1Code)(&value)
}

func (s *SettlementSystemMethod3Choice) AddProprietary() *GenericIdentification38 {
	s.Proprietary = new(GenericIdentification38)
	return s.Proprietary
}
