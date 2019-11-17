package model

// Choice between formats for the balance information.
type SubBalanceQuantity1Choice struct {

	// Quantity of securities in the sub-balance.
	Quantity *FinancialInstrumentQuantityChoice `xml:"Qty"`

	// Quantity of securities in the sub-balance.
	QuantityAsDSS *GenericIdentification6 `xml:"QtyAsDSS"`

	// Quantity of securities in the sub-balance and whether the balance is available.
	QuantityAndAvailability *QuantityAndAvailability `xml:"QtyAndAvlbty"`
}

func (s *SubBalanceQuantity1Choice) AddQuantity() *FinancialInstrumentQuantityChoice {
	s.Quantity = new(FinancialInstrumentQuantityChoice)
	return s.Quantity
}

func (s *SubBalanceQuantity1Choice) AddQuantityAsDSS() *GenericIdentification6 {
	s.QuantityAsDSS = new(GenericIdentification6)
	return s.QuantityAsDSS
}

func (s *SubBalanceQuantity1Choice) AddQuantityAndAvailability() *QuantityAndAvailability {
	s.QuantityAndAvailability = new(QuantityAndAvailability)
	return s.QuantityAndAvailability
}
