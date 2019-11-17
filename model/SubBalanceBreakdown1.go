package model

// Net position of a segregated holding of a single financial instrument within the overall position held in the securities account, for example, sub-balance per status.
type SubBalanceBreakdown1 struct {

	// Reason for the sub-balance.
	SubBalanceType *SubBalanceType9Choice `xml:"SubBalTp"`

	// Quantity of financial instrument in the sub-balance.
	Quantity *SubBalanceQuantity5Choice `xml:"Qty"`
}

func (s *SubBalanceBreakdown1) AddSubBalanceType() *SubBalanceType9Choice {
	s.SubBalanceType = new(SubBalanceType9Choice)
	return s.SubBalanceType
}

func (s *SubBalanceBreakdown1) AddQuantity() *SubBalanceQuantity5Choice {
	s.Quantity = new(SubBalanceQuantity5Choice)
	return s.Quantity
}
