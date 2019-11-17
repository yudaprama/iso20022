package model

// Choice between formats for the balance information.
type SubBalanceQuantity7Choice struct {

	// Quantity of securities in the sub-balance.
	Quantity *FinancialInstrumentQuantity15Choice `xml:"Qty"`

	// Quantity of securities in the sub-balance.
	Proprietary *GenericIdentification144 `xml:"Prtry"`

	// Quantity of securities in the sub-balance and whether the balance is available.
	QuantityAndAvailability *QuantityAndAvailability2 `xml:"QtyAndAvlbty"`
}

func (s *SubBalanceQuantity7Choice) AddQuantity() *FinancialInstrumentQuantity15Choice {
	s.Quantity = new(FinancialInstrumentQuantity15Choice)
	return s.Quantity
}

func (s *SubBalanceQuantity7Choice) AddProprietary() *GenericIdentification144 {
	s.Proprietary = new(GenericIdentification144)
	return s.Proprietary
}

func (s *SubBalanceQuantity7Choice) AddQuantityAndAvailability() *QuantityAndAvailability2 {
	s.QuantityAndAvailability = new(QuantityAndAvailability2)
	return s.QuantityAndAvailability
}
