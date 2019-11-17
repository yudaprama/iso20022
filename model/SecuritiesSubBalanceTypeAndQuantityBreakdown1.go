package model

// Quantity breakdown information for a specific securities balance.
type SecuritiesSubBalanceTypeAndQuantityBreakdown1 struct {

	// Specifies the securities sub balance type indicator (example restriction type for a market infrastructure).
	Type *SecuritiesBalanceType3Choice `xml:"Tp"`

	// Breakdown of a quantity into lots such as tax lots, instrument series.
	QuantityBreakdown []*QuantityBreakdown12 `xml:"QtyBrkdwn,omitempty"`
}

func (s *SecuritiesSubBalanceTypeAndQuantityBreakdown1) AddType() *SecuritiesBalanceType3Choice {
	s.Type = new(SecuritiesBalanceType3Choice)
	return s.Type
}

func (s *SecuritiesSubBalanceTypeAndQuantityBreakdown1) AddQuantityBreakdown() *QuantityBreakdown12 {
	newValue := new(QuantityBreakdown12)
	s.QuantityBreakdown = append(s.QuantityBreakdown, newValue)
	return newValue
}
