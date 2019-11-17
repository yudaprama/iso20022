package model

// Quantity breakdown information for a specific securities balance.
type SecuritiesSubBalanceTypeAndQuantityBreakdown3 struct {

	// Specifies the securities sub balance type indicator (example restriction type for a market infrastructure).
	Type *SecuritiesBalanceType6Choice `xml:"Tp"`

	// Breakdown of a quantity into lots such as tax lots, instrument series.
	QuantityBreakdown []*QuantityBreakdown32 `xml:"QtyBrkdwn,omitempty"`
}

func (s *SecuritiesSubBalanceTypeAndQuantityBreakdown3) AddType() *SecuritiesBalanceType6Choice {
	s.Type = new(SecuritiesBalanceType6Choice)
	return s.Type
}

func (s *SecuritiesSubBalanceTypeAndQuantityBreakdown3) AddQuantityBreakdown() *QuantityBreakdown32 {
	newValue := new(QuantityBreakdown32)
	s.QuantityBreakdown = append(s.QuantityBreakdown, newValue)
	return newValue
}
