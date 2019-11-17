package model

// Order confirmation details.
type SubscriptionOrderConfirmation1 struct {

	// Indicates whether a confirmation amendment message will follow the confirmation cancellation instruction or not.
	AmendmentIndicator *YesNoIndicator `xml:"AmdmntInd"`

	// General information related to the execution of investment fund order.
	MultipleExecutionDetails *SubscriptionMultipleExecution3 `xml:"MltplExctnDtls"`

	// Information about parties related to the transaction.
	RelatedPartyDetails []*Intermediary9 `xml:"RltdPtyDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (s *SubscriptionOrderConfirmation1) SetAmendmentIndicator(value string) {
	s.AmendmentIndicator = (*YesNoIndicator)(&value)
}

func (s *SubscriptionOrderConfirmation1) AddMultipleExecutionDetails() *SubscriptionMultipleExecution3 {
	s.MultipleExecutionDetails = new(SubscriptionMultipleExecution3)
	return s.MultipleExecutionDetails
}

func (s *SubscriptionOrderConfirmation1) AddRelatedPartyDetails() *Intermediary9 {
	newValue := new(Intermediary9)
	s.RelatedPartyDetails = append(s.RelatedPartyDetails, newValue)
	return newValue
}

func (s *SubscriptionOrderConfirmation1) AddExtension() *Extension1 {
	newValue := new(Extension1)
	s.Extension = append(s.Extension, newValue)
	return newValue
}
