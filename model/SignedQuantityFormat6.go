package model

// Signed quantity of security formats.
type SignedQuantityFormat6 struct {

	// Sign of the quantity of security.
	ShortLongPosition *ShortLong1Code `xml:"ShrtLngPos"`

	// Quantity of security.
	Quantity *FinancialInstrumentQuantity1Choice `xml:"Qty"`
}

func (s *SignedQuantityFormat6) SetShortLongPosition(value string) {
	s.ShortLongPosition = (*ShortLong1Code)(&value)
}

func (s *SignedQuantityFormat6) AddQuantity() *FinancialInstrumentQuantity1Choice {
	s.Quantity = new(FinancialInstrumentQuantity1Choice)
	return s.Quantity
}
