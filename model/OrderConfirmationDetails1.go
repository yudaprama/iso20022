package model

// Order confirmation details.
type OrderConfirmationDetails1 struct {

	// Indicates whether a confirmation amendment message will follow the confirmation cancellation instruction or not.
	AmendmentIndicator *YesNoIndicator `xml:"AmdmntInd"`

	// General information related to the execution of investment fund orders.
	BulkExecutionDetails *RedemptionBulkExecution3 `xml:"BlkExctnDtls"`

	// Information about parties related to the transaction.
	RelatedPartyDetails []*Intermediary9 `xml:"RltdPtyDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (o *OrderConfirmationDetails1) SetAmendmentIndicator(value string) {
	o.AmendmentIndicator = (*YesNoIndicator)(&value)
}

func (o *OrderConfirmationDetails1) AddBulkExecutionDetails() *RedemptionBulkExecution3 {
	o.BulkExecutionDetails = new(RedemptionBulkExecution3)
	return o.BulkExecutionDetails
}

func (o *OrderConfirmationDetails1) AddRelatedPartyDetails() *Intermediary9 {
	newValue := new(Intermediary9)
	o.RelatedPartyDetails = append(o.RelatedPartyDetails, newValue)
	return newValue
}

func (o *OrderConfirmationDetails1) AddExtension() *Extension1 {
	newValue := new(Extension1)
	o.Extension = append(o.Extension, newValue)
	return newValue
}
