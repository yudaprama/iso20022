package model

// Information about how an order is to be processed.
type OrderWaiver1 struct {

	// Reason why the order has to be handled differently, probably in a manual fashion, because, for example, the investment manager has agreed a waiver to the terms.
	OrderWaiverReason []*OrderWaiverReason3Choice `xml:"OrdrWvrRsn,omitempty"`

	// Additional information about the order waiver.
	InformationValue *Max350Text `xml:"InfVal,omitempty"`
}

func (o *OrderWaiver1) AddOrderWaiverReason() *OrderWaiverReason3Choice {
	newValue := new(OrderWaiverReason3Choice)
	o.OrderWaiverReason = append(o.OrderWaiverReason, newValue)
	return newValue
}

func (o *OrderWaiver1) SetInformationValue(value string) {
	o.InformationValue = (*Max350Text)(&value)
}
