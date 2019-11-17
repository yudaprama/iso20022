package model

// Signed quantity of security formats.
type SignedQuantityFormat7 struct {

	// Sign of the quantity of security.
	ShortLongPosition *ShortLong1Code `xml:"ShrtLngPos"`

	// Choice between different quantity of security formats.
	QuantityChoice *Quantity19Choice `xml:"QtyChc"`
}

func (s *SignedQuantityFormat7) SetShortLongPosition(value string) {
	s.ShortLongPosition = (*ShortLong1Code)(&value)
}

func (s *SignedQuantityFormat7) AddQuantityChoice() *Quantity19Choice {
	s.QuantityChoice = new(Quantity19Choice)
	return s.QuantityChoice
}
