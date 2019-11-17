package model

// Parameters applied to the settlement of a security transfer.
type ReceiveInformation1 struct {

	// Chain of parties involved in the settlement of a transaction.
	SettlementPartiesDetails *ReceivingPartiesAndAccount1 `xml:"SttlmPtiesDtls"`

	// Indicates whether the financial instrument is to be physically delivered.
	PhysicalTransferIndicator *YesNoIndicator `xml:"PhysTrfInd"`

	// Parameters of a physical delivery.
	PhysicalTransferDetails *DeliveryParameters2 `xml:"PhysTrfDtls,omitempty"`
}

func (r *ReceiveInformation1) AddSettlementPartiesDetails() *ReceivingPartiesAndAccount1 {
	r.SettlementPartiesDetails = new(ReceivingPartiesAndAccount1)
	return r.SettlementPartiesDetails
}

func (r *ReceiveInformation1) SetPhysicalTransferIndicator(value string) {
	r.PhysicalTransferIndicator = (*YesNoIndicator)(&value)
}

func (r *ReceiveInformation1) AddPhysicalTransferDetails() *DeliveryParameters2 {
	r.PhysicalTransferDetails = new(DeliveryParameters2)
	return r.PhysicalTransferDetails
}
