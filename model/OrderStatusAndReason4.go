package model

// Status report of a bulk or multiple or switch order cancellation instruction that was previously received.
type OrderStatusAndReason4 struct {

	// Status of the order.
	Status *OrderStatus3Code `xml:"Sts"`

	// Status of the order is rejected.
	Rejected *RejectedStatus4 `xml:"Rjctd"`

	// Party that initiates the status of the order.
	StatusInitiator *PartyIdentification2Choice `xml:"StsInitr,omitempty"`
}

func (o *OrderStatusAndReason4) SetStatus(value string) {
	o.Status = (*OrderStatus3Code)(&value)
}

func (o *OrderStatusAndReason4) AddRejected() *RejectedStatus4 {
	o.Rejected = new(RejectedStatus4)
	return o.Rejected
}

func (o *OrderStatusAndReason4) AddStatusInitiator() *PartyIdentification2Choice {
	o.StatusInitiator = new(PartyIdentification2Choice)
	return o.StatusInitiator
}
