package model

// Choice between an instructed quantity or a quantity to receive.
type InstructedOrQuantityToReceive1Choice struct {

	// Quantity of underlying securities to which this instruction applies.
	InstructedQuantity *Quantity5Choice `xml:"InstdQty"`

	// Quantity of the benefits that the account owner wants to receive, for example, as a result of dividend reinvestment.
	QuantityToReceive *Quantity5Choice `xml:"QtyToRcv"`
}

func (i *InstructedOrQuantityToReceive1Choice) AddInstructedQuantity() *Quantity5Choice {
	i.InstructedQuantity = new(Quantity5Choice)
	return i.InstructedQuantity
}

func (i *InstructedOrQuantityToReceive1Choice) AddQuantityToReceive() *Quantity5Choice {
	i.QuantityToReceive = new(Quantity5Choice)
	return i.QuantityToReceive
}
