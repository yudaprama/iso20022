package model

// Specifies the security option of a corporate event.
type SecuritiesOption2 struct {

	// Minimum quantity of securities to be accepted (used in the framework of conditional privilege on election). In case of proration, if this minimum quantity is not reached then the instruction is void.
	ConditionalQuantity *FinancialInstrumentQuantity1Choice `xml:"CondlQty,omitempty"`

	// Quantity instructed to be received over and above normal ensured entitlement.
	OverAndAboveNormalEnsuredEntitlementQuantity *FinancialInstrumentQuantity1Choice `xml:"OverAndAbovNrmlNsrdEntitlmntQty,omitempty"`

	// Specifies whether the quantity of financial instrument is a quantity of securities instructed or a quantity to receive.
	InstructedOrQuantityToReceive *InstructedOrQuantityToReceive1Choice `xml:"InstdOrQtyToRcv"`
}

func (s *SecuritiesOption2) AddConditionalQuantity() *FinancialInstrumentQuantity1Choice {
	s.ConditionalQuantity = new(FinancialInstrumentQuantity1Choice)
	return s.ConditionalQuantity
}

func (s *SecuritiesOption2) AddOverAndAboveNormalEnsuredEntitlementQuantity() *FinancialInstrumentQuantity1Choice {
	s.OverAndAboveNormalEnsuredEntitlementQuantity = new(FinancialInstrumentQuantity1Choice)
	return s.OverAndAboveNormalEnsuredEntitlementQuantity
}

func (s *SecuritiesOption2) AddInstructedOrQuantityToReceive() *InstructedOrQuantityToReceive1Choice {
	s.InstructedOrQuantityToReceive = new(InstructedOrQuantityToReceive1Choice)
	return s.InstructedOrQuantityToReceive
}
