package model

// Completion of a securities settlement instruction, wherein securities are delivered/debited from a securities account and received/credited to the designated securities account.
type DeliverInformation2 struct {

	// Charge related to the transfer of a financial instrument.
	ChargeDetails []*Charge4 `xml:"ChrgDtls,omitempty"`

	// Tax related to the transfer of a financial instrument.
	TaxDetails []*Tax3 `xml:"TaxDtls,omitempty"`

	// Chain of parties involved in the settlement of a transaction.
	SettlementPartiesDetails *DeliveringPartiesAndAccount1 `xml:"SttlmPtiesDtls"`

	// Indicates whether the financial instrument is to be physically  delivered.
	PhysicalTransferIndicator *YesNoIndicator `xml:"PhysTrfInd"`

	// Parameters of a physical delivery.
	PhysicalTransferDetails *DeliveryParameters2 `xml:"PhysTrfDtls,omitempty"`
}

func (d *DeliverInformation2) AddChargeDetails() *Charge4 {
	newValue := new(Charge4)
	d.ChargeDetails = append(d.ChargeDetails, newValue)
	return newValue
}

func (d *DeliverInformation2) AddTaxDetails() *Tax3 {
	newValue := new(Tax3)
	d.TaxDetails = append(d.TaxDetails, newValue)
	return newValue
}

func (d *DeliverInformation2) AddSettlementPartiesDetails() *DeliveringPartiesAndAccount1 {
	d.SettlementPartiesDetails = new(DeliveringPartiesAndAccount1)
	return d.SettlementPartiesDetails
}

func (d *DeliverInformation2) SetPhysicalTransferIndicator(value string) {
	d.PhysicalTransferIndicator = (*YesNoIndicator)(&value)
}

func (d *DeliverInformation2) AddPhysicalTransferDetails() *DeliveryParameters2 {
	d.PhysicalTransferDetails = new(DeliveryParameters2)
	return d.PhysicalTransferDetails
}
