package model

// Choice of status for an order.
type OrderStatus5Choice struct {

	// Status of the individual order is accepted or already executed or sent to next party or received. There is no reason attached.
	Status *OrderStatus4Code `xml:"Sts"`

	// Status of the individual order is cancelled. This status is used for an order that has been accepted or that has been entered in an order book but that can not be executed.
	Cancelled *CancelledStatusReason16 `xml:"Canc"`

	// Status of the individual order is conditionally accepted.
	ConditionallyAccepted *ConditionallyAcceptedStatus3Choice `xml:"CondlyAccptd"`

	// Status of the individual order is rejected. This status is used for an order that has not been accepted or entered in an order book.
	Rejected []*RejectedStatus9 `xml:"Rjctd"`

	// Status of the individual order is suspended.
	Suspended *SuspendedStatusReason4Choice `xml:"Sspd"`

	// Status of the individual order is in repair.
	InRepair *InRepairStatusReason4Choice `xml:"InRpr"`

	// Status of the individual order is partially settled.
	PartiallySettled *PartiallySettledStatus10 `xml:"PrtlySttld"`
}

func (o *OrderStatus5Choice) SetStatus(value string) {
	o.Status = (*OrderStatus4Code)(&value)
}

func (o *OrderStatus5Choice) AddCancelled() *CancelledStatusReason16 {
	o.Cancelled = new(CancelledStatusReason16)
	return o.Cancelled
}

func (o *OrderStatus5Choice) AddConditionallyAccepted() *ConditionallyAcceptedStatus3Choice {
	o.ConditionallyAccepted = new(ConditionallyAcceptedStatus3Choice)
	return o.ConditionallyAccepted
}

func (o *OrderStatus5Choice) AddRejected() *RejectedStatus9 {
	newValue := new(RejectedStatus9)
	o.Rejected = append(o.Rejected, newValue)
	return newValue
}

func (o *OrderStatus5Choice) AddSuspended() *SuspendedStatusReason4Choice {
	o.Suspended = new(SuspendedStatusReason4Choice)
	return o.Suspended
}

func (o *OrderStatus5Choice) AddInRepair() *InRepairStatusReason4Choice {
	o.InRepair = new(InRepairStatusReason4Choice)
	return o.InRepair
}

func (o *OrderStatus5Choice) AddPartiallySettled() *PartiallySettledStatus10 {
	o.PartiallySettled = new(PartiallySettledStatus10)
	return o.PartiallySettled
}
