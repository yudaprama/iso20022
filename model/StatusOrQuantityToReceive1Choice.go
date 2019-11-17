package model

// Choice between a status quantity or a quantity to receive for the quantity of financial instrument.
type StatusOrQuantityToReceive1Choice struct {

	// Quantity of securities that has been assigned the status indicated.
	StatusQuantity *Quantity6Choice `xml:"StsQty"`

	// Quantity of the benefits that the account owner wants to receive, for example, as a result of dividend reinvestment.
	QuantityToReceive *Quantity6Choice `xml:"QtyToRcv"`
}

func (s *StatusOrQuantityToReceive1Choice) AddStatusQuantity() *Quantity6Choice {
	s.StatusQuantity = new(Quantity6Choice)
	return s.StatusQuantity
}

func (s *StatusOrQuantityToReceive1Choice) AddQuantityToReceive() *Quantity6Choice {
	s.QuantityToReceive = new(Quantity6Choice)
	return s.QuantityToReceive
}
