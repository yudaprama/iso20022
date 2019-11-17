package model

// Order confirmation details.
type RedemptionOrderConfirmation1 struct {

	// Indicates whether a confirmation amendment message will follow the confirmation cancellation instruction or not.
	AmendmentIndicator *YesNoIndicator `xml:"AmdmntInd"`

	// General information related to the execution of investment fund orders.
	MultipleExecutionDetails *RedemptionMultipleExecution3 `xml:"MltplExctnDtls"`

	// Information about parties related to the transaction.
	RelatedPartyDetails []*Intermediary9 `xml:"RltdPtyDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (r *RedemptionOrderConfirmation1) SetAmendmentIndicator(value string) {
	r.AmendmentIndicator = (*YesNoIndicator)(&value)
}

func (r *RedemptionOrderConfirmation1) AddMultipleExecutionDetails() *RedemptionMultipleExecution3 {
	r.MultipleExecutionDetails = new(RedemptionMultipleExecution3)
	return r.MultipleExecutionDetails
}

func (r *RedemptionOrderConfirmation1) AddRelatedPartyDetails() *Intermediary9 {
	newValue := new(Intermediary9)
	r.RelatedPartyDetails = append(r.RelatedPartyDetails, newValue)
	return newValue
}

func (r *RedemptionOrderConfirmation1) AddExtension() *Extension1 {
	newValue := new(Extension1)
	r.Extension = append(r.Extension, newValue)
	return newValue
}
