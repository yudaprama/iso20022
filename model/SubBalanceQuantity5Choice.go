package model

// Choice between formats for the balance information.
type SubBalanceQuantity5Choice struct {

	// Quantity of financial instrument in the sub-balance.
	Quantity *FinancialInstrumentQuantity1Choice `xml:"Qty"`

	// Quantity of financial instrument in the sub-balance expressed in a proprietary format.
	Proprietary *GenericIdentification56 `xml:"Prtry"`
}

func (s *SubBalanceQuantity5Choice) AddQuantity() *FinancialInstrumentQuantity1Choice {
	s.Quantity = new(FinancialInstrumentQuantity1Choice)
	return s.Quantity
}

func (s *SubBalanceQuantity5Choice) AddProprietary() *GenericIdentification56 {
	s.Proprietary = new(GenericIdentification56)
	return s.Proprietary
}
