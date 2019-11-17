package model

// Reference of an order, deal reference, client reference and master reference.
type InvestmentFundOrderExecution1 struct {

	// Unique and unambiguous identifier for a group of individual orders, as assigned by the instructing party. This identifier links the individual orders together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Indicates whether a confirmation amendment message will follow the confirmation cancellation instruction or not.
	AmendmentIndicator *YesNoIndicator `xml:"AmdmntInd"`

	// Reference of an order, client or deal reference.
	OrderReferences []*InvestmentFundOrderExecution2 `xml:"OrdrRefs"`
}

func (i *InvestmentFundOrderExecution1) SetMasterReference(value string) {
	i.MasterReference = (*Max35Text)(&value)
}

func (i *InvestmentFundOrderExecution1) SetAmendmentIndicator(value string) {
	i.AmendmentIndicator = (*YesNoIndicator)(&value)
}

func (i *InvestmentFundOrderExecution1) AddOrderReferences() *InvestmentFundOrderExecution2 {
	newValue := new(InvestmentFundOrderExecution2)
	i.OrderReferences = append(i.OrderReferences, newValue)
	return newValue
}
