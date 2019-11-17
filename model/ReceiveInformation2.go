package model

// Completion of a securities settlement instruction, wherein securities are delivered/debited from a securities account and received/credited to the designated securities account.
type ReceiveInformation2 struct {

	// Charge related to the transfer of a financial instrument.
	ChargeDetails []*Charge4 `xml:"ChrgDtls,omitempty"`

	// Tax related to the transfer of a financial instrument.
	TaxDetails []*Tax3 `xml:"TaxDtls,omitempty"`

	// Chain of parties involved in the settlement of a transaction.
	SettlementPartiesDetails *ReceivingPartiesAndAccount1 `xml:"SttlmPtiesDtls"`

	// Indicates whether the financial instrument is to be physically  delivered.
	PhysicalTransferIndicator *YesNoIndicator `xml:"PhysTrfInd"`

	// Parameters of a physical delivery.
	PhysicalTransferDetails *DeliveryParameters2 `xml:"PhysTrfDtls,omitempty"`
}

func (r *ReceiveInformation2) AddChargeDetails() *Charge4 {
	newValue := new(Charge4)
	r.ChargeDetails = append(r.ChargeDetails, newValue)
	return newValue
}

func (r *ReceiveInformation2) AddTaxDetails() *Tax3 {
	newValue := new(Tax3)
	r.TaxDetails = append(r.TaxDetails, newValue)
	return newValue
}

func (r *ReceiveInformation2) AddSettlementPartiesDetails() *ReceivingPartiesAndAccount1 {
	r.SettlementPartiesDetails = new(ReceivingPartiesAndAccount1)
	return r.SettlementPartiesDetails
}

func (r *ReceiveInformation2) SetPhysicalTransferIndicator(value string) {
	r.PhysicalTransferIndicator = (*YesNoIndicator)(&value)
}

func (r *ReceiveInformation2) AddPhysicalTransferDetails() *DeliveryParameters2 {
	r.PhysicalTransferDetails = new(DeliveryParameters2)
	return r.PhysicalTransferDetails
}
