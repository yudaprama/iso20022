package model

// Status report of a bulk or multiple or switch order that was previously received.
type OrderStatusAndReason3 struct {

	// Status of the order is accepted or already executed or sent to next party or received. There is no reason attached.
	Status *OrderStatus2Code `xml:"Sts"`

	// Status of the order details is cancelled. This status is used for orders that have been accepted or that have been entered in an order book but that can not be executed.
	Cancelled *CancelledStatus1 `xml:"Canc"`

	// Status of the order details is conditionally accepted.
	ConditionallyAccepted *ConditionallyAcceptedStatus1 `xml:"CondlyAccptd"`

	// Status of the order details is rejected. This status is used for orders that have not been accepted or entered in an order book.
	Rejected *RejectedStatus3 `xml:"Rjctd"`

	// Status of the order details is suspended.
	Suspended *SuspendedStatus1 `xml:"Sspd"`

	// Party that initiates the status of the order.
	StatusInitiator *PartyIdentification2Choice `xml:"StsInitr,omitempty"`

	// Unique and unambiguous technical identification of an instance of a leg within a switch.
	SwitchOrderLegIdentification *Max35Text `xml:"SwtchOrdrLegId,omitempty"`
}

func (o *OrderStatusAndReason3) SetStatus(value string) {
	o.Status = (*OrderStatus2Code)(&value)
}

func (o *OrderStatusAndReason3) AddCancelled() *CancelledStatus1 {
	o.Cancelled = new(CancelledStatus1)
	return o.Cancelled
}

func (o *OrderStatusAndReason3) AddConditionallyAccepted() *ConditionallyAcceptedStatus1 {
	o.ConditionallyAccepted = new(ConditionallyAcceptedStatus1)
	return o.ConditionallyAccepted
}

func (o *OrderStatusAndReason3) AddRejected() *RejectedStatus3 {
	o.Rejected = new(RejectedStatus3)
	return o.Rejected
}

func (o *OrderStatusAndReason3) AddSuspended() *SuspendedStatus1 {
	o.Suspended = new(SuspendedStatus1)
	return o.Suspended
}

func (o *OrderStatusAndReason3) AddStatusInitiator() *PartyIdentification2Choice {
	o.StatusInitiator = new(PartyIdentification2Choice)
	return o.StatusInitiator
}

func (o *OrderStatusAndReason3) SetSwitchOrderLegIdentification(value string) {
	o.SwitchOrderLegIdentification = (*Max35Text)(&value)
}
