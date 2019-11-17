package model

// Status report of a bulk or multiple or switch order cancellation instruction that was previously received.
type OrderStatusAndReason8 struct {

	// Reference assigned to a set of orders or trades in order to link them together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Cancellation status of the order.
	Status *OrderCancellationStatus1Code `xml:"Sts"`

	// Status of the order cancellation request is rejected.
	Rejected *RejectedStatus7 `xml:"Rjctd"`

	// Party that initiates the status of the order cancellation.
	StatusInitiator *PartyIdentification2Choice `xml:"StsInitr,omitempty"`
}

func (o *OrderStatusAndReason8) SetMasterReference(value string) {
	o.MasterReference = (*Max35Text)(&value)
}

func (o *OrderStatusAndReason8) SetStatus(value string) {
	o.Status = (*OrderCancellationStatus1Code)(&value)
}

func (o *OrderStatusAndReason8) AddRejected() *RejectedStatus7 {
	o.Rejected = new(RejectedStatus7)
	return o.Rejected
}

func (o *OrderStatusAndReason8) AddStatusInitiator() *PartyIdentification2Choice {
	o.StatusInitiator = new(PartyIdentification2Choice)
	return o.StatusInitiator
}
