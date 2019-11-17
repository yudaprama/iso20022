package model

// Information about a redemption bulk order.
type RedemptionBulkOrderInstruction2 struct {

	// Common information related to all the orders.
	BulkOrderDetails *RedemptionBulkOrder3 `xml:"BlkOrdrDtls"`

	// Information about parties related to the transaction.
	RelatedPartyDetails []*Intermediary8 `xml:"RltdPtyDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (r *RedemptionBulkOrderInstruction2) AddBulkOrderDetails() *RedemptionBulkOrder3 {
	r.BulkOrderDetails = new(RedemptionBulkOrder3)
	return r.BulkOrderDetails
}

func (r *RedemptionBulkOrderInstruction2) AddRelatedPartyDetails() *Intermediary8 {
	newValue := new(Intermediary8)
	r.RelatedPartyDetails = append(r.RelatedPartyDetails, newValue)
	return newValue
}

func (r *RedemptionBulkOrderInstruction2) AddExtension() *Extension1 {
	newValue := new(Extension1)
	r.Extension = append(r.Extension, newValue)
	return newValue
}
