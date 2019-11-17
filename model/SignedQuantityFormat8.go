package model

// Signed quantity of security formats.
type SignedQuantityFormat8 struct {

	// Sign of the quantity of security.
	ShortLongPosition *ShortLong1Code `xml:"ShrtLngPos"`

	// Choice between different quantity of security formats.
	QuantityChoice *Quantity21Choice `xml:"QtyChc"`
}

func (s *SignedQuantityFormat8) SetShortLongPosition(value string) {
	s.ShortLongPosition = (*ShortLong1Code)(&value)
}

func (s *SignedQuantityFormat8) AddQuantityChoice() *Quantity21Choice {
	s.QuantityChoice = new(Quantity21Choice)
	return s.QuantityChoice
}
