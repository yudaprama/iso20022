package model

// Choice of format for the settlement system/method.
type SettlementSystemMethod1Choice struct {

	// Settlement system expressed as an ISO 20022 code.
	Code *SettlementSystemMethod1Code `xml:"Cd"`

	// Settlement system expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (s *SettlementSystemMethod1Choice) SetCode(value string) {
	s.Code = (*SettlementSystemMethod1Code)(&value)
}

func (s *SettlementSystemMethod1Choice) AddProprietary() *GenericIdentification20 {
	s.Proprietary = new(GenericIdentification20)
	return s.Proprietary
}
