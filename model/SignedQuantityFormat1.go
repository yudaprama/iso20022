package model

// Signed quantity of security formats.
type SignedQuantityFormat1 struct {

	// Sign of the quantity of security.
	ShortLongPosition *ShortLong1Code `xml:"ShrtLngPos"`

	// Choice between different quantity of security formats.
	QuantityChoice *Quantity2Choice `xml:"QtyChc"`
}

func (s *SignedQuantityFormat1) SetShortLongPosition(value string) {
	s.ShortLongPosition = (*ShortLong1Code)(&value)
}

func (s *SignedQuantityFormat1) AddQuantityChoice() *Quantity2Choice {
	s.QuantityChoice = new(Quantity2Choice)
	return s.QuantityChoice
}
