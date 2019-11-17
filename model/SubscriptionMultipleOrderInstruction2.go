package model

// Information about a subscription multiple order.
type SubscriptionMultipleOrderInstruction2 struct {

	// Common information related to all the orders to be cancelled.
	MultipleOrderDetails *SubscriptionMultipleOrder3 `xml:"MltplOrdrDtls"`

	// Information about parties related to the transaction.
	RelatedPartyDetails []*Intermediary8 `xml:"RltdPtyDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (s *SubscriptionMultipleOrderInstruction2) AddMultipleOrderDetails() *SubscriptionMultipleOrder3 {
	s.MultipleOrderDetails = new(SubscriptionMultipleOrder3)
	return s.MultipleOrderDetails
}

func (s *SubscriptionMultipleOrderInstruction2) AddRelatedPartyDetails() *Intermediary8 {
	newValue := new(Intermediary8)
	s.RelatedPartyDetails = append(s.RelatedPartyDetails, newValue)
	return newValue
}

func (s *SubscriptionMultipleOrderInstruction2) AddExtension() *Extension1 {
	newValue := new(Extension1)
	s.Extension = append(s.Extension, newValue)
	return newValue
}
