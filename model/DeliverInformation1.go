package model

// Parameters applied to the settlement of a security transfer.
type DeliverInformation1 struct {

	// Chain of parties involved in the settlement of a transaction.
	SettlementPartiesDetails *DeliveringPartiesAndAccount1 `xml:"SttlmPtiesDtls"`

	// Indicates whether the financial instrument is to be physically delivered.
	PhysicalTransferIndicator *YesNoIndicator `xml:"PhysTrfInd"`

	// Parameters of a physical delivery.
	PhysicalTransferDetails *DeliveryParameters2 `xml:"PhysTrfDtls,omitempty"`
}

func (d *DeliverInformation1) AddSettlementPartiesDetails() *DeliveringPartiesAndAccount1 {
	d.SettlementPartiesDetails = new(DeliveringPartiesAndAccount1)
	return d.SettlementPartiesDetails
}

func (d *DeliverInformation1) SetPhysicalTransferIndicator(value string) {
	d.PhysicalTransferIndicator = (*YesNoIndicator)(&value)
}

func (d *DeliverInformation1) AddPhysicalTransferDetails() *DeliveryParameters2 {
	d.PhysicalTransferDetails = new(DeliveryParameters2)
	return d.PhysicalTransferDetails
}
