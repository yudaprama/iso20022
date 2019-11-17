package model

// Signed quantity of security formats.
type SignedQuantityFormat9 struct {

	// Sign of the quantity of security.
	ShortLongPosition *ShortLong1Code `xml:"ShrtLngPos"`

	// Quantity of security.
	Quantity *FinancialInstrumentQuantity15Choice `xml:"Qty"`
}

func (s *SignedQuantityFormat9) SetShortLongPosition(value string) {
	s.ShortLongPosition = (*ShortLong1Code)(&value)
}

func (s *SignedQuantityFormat9) AddQuantity() *FinancialInstrumentQuantity15Choice {
	s.Quantity = new(FinancialInstrumentQuantity15Choice)
	return s.Quantity
}
