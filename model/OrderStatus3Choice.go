package model

// Choice of status for an order.
type OrderStatus3Choice struct {

	// Status of all the orders in the order message. There is no reason attached.
	Status *OrderStatus4Code `xml:"Sts"`

	// Status of all the orders in the order message is cancelled. This status is used for orders that have been accepted or that have been entered in an order book but that can not be executed.
	Cancelled *CancelledStatusReason16 `xml:"Canc"`

	// Status of all the orders in the order message is conditionally accepted.
	ConditionallyAccepted *ConditionallyAcceptedStatus3Choice `xml:"CondlyAccptd"`

	// Status of all the orders in the order message is rejected. This status is used for orders that have not been accepted or entered in an order book.
	Rejected []*RejectedStatus9 `xml:"Rjctd"`

	// Status of all the orders in the order message is suspended.
	Suspended *SuspendedStatusReason4Choice `xml:"Sspd"`

	// Status of all the orders in the order message is partially settled.
	PartiallySettled *PartiallySettledStatus10 `xml:"PrtlySttld"`
}

func (o *OrderStatus3Choice) SetStatus(value string) {
	o.Status = (*OrderStatus4Code)(&value)
}

func (o *OrderStatus3Choice) AddCancelled() *CancelledStatusReason16 {
	o.Cancelled = new(CancelledStatusReason16)
	return o.Cancelled
}

func (o *OrderStatus3Choice) AddConditionallyAccepted() *ConditionallyAcceptedStatus3Choice {
	o.ConditionallyAccepted = new(ConditionallyAcceptedStatus3Choice)
	return o.ConditionallyAccepted
}

func (o *OrderStatus3Choice) AddRejected() *RejectedStatus9 {
	newValue := new(RejectedStatus9)
	o.Rejected = append(o.Rejected, newValue)
	return newValue
}

func (o *OrderStatus3Choice) AddSuspended() *SuspendedStatusReason4Choice {
	o.Suspended = new(SuspendedStatusReason4Choice)
	return o.Suspended
}

func (o *OrderStatus3Choice) AddPartiallySettled() *PartiallySettledStatus10 {
	o.PartiallySettled = new(PartiallySettledStatus10)
	return o.PartiallySettled
}
