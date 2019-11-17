package model

// Specifies the security option of a corporate event.
type SecuritiesOption52 struct {

	// Minimum quantity of securities to be accepted (used in the framework of conditional privilege on election). In case of proration, if this minimum quantity is not reached then the instruction is void.
	ConditionalQuantity *FinancialInstrumentQuantity1Choice `xml:"CondlQty,omitempty"`

	// Quantity of securities to which this instruction applies.
	InstructedQuantity *Quantity20Choice `xml:"InstdQty"`
}

func (s *SecuritiesOption52) AddConditionalQuantity() *FinancialInstrumentQuantity1Choice {
	s.ConditionalQuantity = new(FinancialInstrumentQuantity1Choice)
	return s.ConditionalQuantity
}

func (s *SecuritiesOption52) AddInstructedQuantity() *Quantity20Choice {
	s.InstructedQuantity = new(Quantity20Choice)
	return s.InstructedQuantity
}
