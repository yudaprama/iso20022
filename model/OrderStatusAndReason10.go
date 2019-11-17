package model

// Status report of a bulk or multiple or switch order that was previously received.
type OrderStatusAndReason10 struct {

	// Reference assigned to a set of orders or trades in order to link them together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Status of a 'bulk' of orders. Can be used if all the individual orders conveyed in a bulk or multiple order message have the same status.
	OrderStatus *OrderStatus3Choice `xml:"OrdrSts"`

	// Party that initiates the status of the order.
	StatusInitiator *PartyIdentification113 `xml:"StsInitr,omitempty"`
}

func (o *OrderStatusAndReason10) SetMasterReference(value string) {
	o.MasterReference = (*Max35Text)(&value)
}

func (o *OrderStatusAndReason10) AddOrderStatus() *OrderStatus3Choice {
	o.OrderStatus = new(OrderStatus3Choice)
	return o.OrderStatus
}

func (o *OrderStatusAndReason10) AddStatusInitiator() *PartyIdentification113 {
	o.StatusInitiator = new(PartyIdentification113)
	return o.StatusInitiator
}
