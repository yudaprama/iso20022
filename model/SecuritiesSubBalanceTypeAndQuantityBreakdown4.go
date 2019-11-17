package model

// Quantity breakdown information for a specific securities balance.
type SecuritiesSubBalanceTypeAndQuantityBreakdown4 struct {

	// Specifies the securities sub balance type indicator (example restriction type for a market infrastructure).
	Type *SecuritiesBalanceType8Choice `xml:"Tp"`

	// Breakdown of a quantity into lots such as tax lots, instrument series.
	QuantityBreakdown []*QuantityBreakdown34 `xml:"QtyBrkdwn,omitempty"`
}

func (s *SecuritiesSubBalanceTypeAndQuantityBreakdown4) AddType() *SecuritiesBalanceType8Choice {
	s.Type = new(SecuritiesBalanceType8Choice)
	return s.Type
}

func (s *SecuritiesSubBalanceTypeAndQuantityBreakdown4) AddQuantityBreakdown() *QuantityBreakdown34 {
	newValue := new(QuantityBreakdown34)
	s.QuantityBreakdown = append(s.QuantityBreakdown, newValue)
	return newValue
}
