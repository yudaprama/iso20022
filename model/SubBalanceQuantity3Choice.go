package model

// Choice between formats for the balance information.
type SubBalanceQuantity3Choice struct {

	// Quantity of securities in the sub-balance.
	Quantity *FinancialInstrumentQuantity1Choice `xml:"Qty"`

	// Quantity of securities in the sub-balance.
	Proprietary *GenericIdentification22 `xml:"Prtry"`

	// Quantity of securities in the sub-balance and whether the balance is available.
	QuantityAndAvailability *QuantityAndAvailability1 `xml:"QtyAndAvlbty"`
}

func (s *SubBalanceQuantity3Choice) AddQuantity() *FinancialInstrumentQuantity1Choice {
	s.Quantity = new(FinancialInstrumentQuantity1Choice)
	return s.Quantity
}

func (s *SubBalanceQuantity3Choice) AddProprietary() *GenericIdentification22 {
	s.Proprietary = new(GenericIdentification22)
	return s.Proprietary
}

func (s *SubBalanceQuantity3Choice) AddQuantityAndAvailability() *QuantityAndAvailability1 {
	s.QuantityAndAvailability = new(QuantityAndAvailability1)
	return s.QuantityAndAvailability
}
