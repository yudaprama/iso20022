package model

// Status report of a bulk or multiple or switch order cancellation instruction that was previously received.
type OrderStatusAndReason9 struct {

	// Reference assigned to a set of orders or trades in order to link them together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Cancellation status of the order cancellation.
	CancellationStatus *CancellationStatus22Choice `xml:"CxlSts"`

	// Party that initiates the status of the order cancellation.
	StatusInitiator *PartyIdentification113 `xml:"StsInitr,omitempty"`
}

func (o *OrderStatusAndReason9) SetMasterReference(value string) {
	o.MasterReference = (*Max35Text)(&value)
}

func (o *OrderStatusAndReason9) AddCancellationStatus() *CancellationStatus22Choice {
	o.CancellationStatus = new(CancellationStatus22Choice)
	return o.CancellationStatus
}

func (o *OrderStatusAndReason9) AddStatusInitiator() *PartyIdentification113 {
	o.StatusInitiator = new(PartyIdentification113)
	return o.StatusInitiator
}
