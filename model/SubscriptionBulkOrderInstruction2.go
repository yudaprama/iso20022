package model

// Information about a subscription bulk order.
type SubscriptionBulkOrderInstruction2 struct {

	// Common information related to all the orders to be cancelled.
	BulkOrderDetails *SubscriptionBulkOrder3 `xml:"BlkOrdrDtls"`

	// Information about parties related to the transaction.
	RelatedPartyDetails []*Intermediary8 `xml:"RltdPtyDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (s *SubscriptionBulkOrderInstruction2) AddBulkOrderDetails() *SubscriptionBulkOrder3 {
	s.BulkOrderDetails = new(SubscriptionBulkOrder3)
	return s.BulkOrderDetails
}

func (s *SubscriptionBulkOrderInstruction2) AddRelatedPartyDetails() *Intermediary8 {
	newValue := new(Intermediary8)
	s.RelatedPartyDetails = append(s.RelatedPartyDetails, newValue)
	return newValue
}

func (s *SubscriptionBulkOrderInstruction2) AddExtension() *Extension1 {
	newValue := new(Extension1)
	s.Extension = append(s.Extension, newValue)
	return newValue
}
