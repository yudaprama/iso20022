package model

// Specifies the security option of a corporate event.
type SecuritiesOption54 struct {

	// Minimum quantity of securities to be accepted (used in the framework of conditional privilege on election). In case of proration, if this minimum quantity is not reached then the instruction is void.
	ConditionalQuantity *FinancialInstrumentQuantity15Choice `xml:"CondlQty,omitempty"`

	// Quantity of securities to which this instruction applies.
	InstructedQuantity *Quantity40Choice `xml:"InstdQty"`
}

func (s *SecuritiesOption54) AddConditionalQuantity() *FinancialInstrumentQuantity15Choice {
	s.ConditionalQuantity = new(FinancialInstrumentQuantity15Choice)
	return s.ConditionalQuantity
}

func (s *SecuritiesOption54) AddInstructedQuantity() *Quantity40Choice {
	s.InstructedQuantity = new(Quantity40Choice)
	return s.InstructedQuantity
}
