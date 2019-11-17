package model

// Status report of a bulk or multiple or switch order that was previously received.
type OrderStatusAndReason7 struct {

	// Reference assigned to a set of orders or trades in order to link them together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Status of all the orders in the order message. There is no reason attached.
	Status *OrderStatus4Code `xml:"Sts"`

	// Status of all the orders in the order message is cancelled. This status is used for orders that have been accepted or that have been entered in an order book but that can not be executed.
	Cancelled *CancelledStatus2 `xml:"Canc"`

	// Status of all the orders in the order message is conditionally accepted.
	ConditionallyAccepted *ConditionallyAcceptedStatus2 `xml:"CondlyAccptd"`

	// Status of all the orders in the order message is rejected. This status is used for orders that have not been accepted or entered in an order book.
	Rejected []*RejectedStatus6 `xml:"Rjctd"`

	// Status of all the orders in the order message is suspended.
	Suspended *SuspendedStatus2 `xml:"Sspd"`

	// Status of all the orders in the order message is partially settled.
	PartiallySettled *PartiallySettledStatus1 `xml:"PrtlySttld"`

	// Party that initiates the status of the order.
	StatusInitiator *PartyIdentification2Choice `xml:"StsInitr,omitempty"`
}

func (o *OrderStatusAndReason7) SetMasterReference(value string) {
	o.MasterReference = (*Max35Text)(&value)
}

func (o *OrderStatusAndReason7) SetStatus(value string) {
	o.Status = (*OrderStatus4Code)(&value)
}

func (o *OrderStatusAndReason7) AddCancelled() *CancelledStatus2 {
	o.Cancelled = new(CancelledStatus2)
	return o.Cancelled
}

func (o *OrderStatusAndReason7) AddConditionallyAccepted() *ConditionallyAcceptedStatus2 {
	o.ConditionallyAccepted = new(ConditionallyAcceptedStatus2)
	return o.ConditionallyAccepted
}

func (o *OrderStatusAndReason7) AddRejected() *RejectedStatus6 {
	newValue := new(RejectedStatus6)
	o.Rejected = append(o.Rejected, newValue)
	return newValue
}

func (o *OrderStatusAndReason7) AddSuspended() *SuspendedStatus2 {
	o.Suspended = new(SuspendedStatus2)
	return o.Suspended
}

func (o *OrderStatusAndReason7) AddPartiallySettled() *PartiallySettledStatus1 {
	o.PartiallySettled = new(PartiallySettledStatus1)
	return o.PartiallySettled
}

func (o *OrderStatusAndReason7) AddStatusInitiator() *PartyIdentification2Choice {
	o.StatusInitiator = new(PartyIdentification2Choice)
	return o.StatusInitiator
}
