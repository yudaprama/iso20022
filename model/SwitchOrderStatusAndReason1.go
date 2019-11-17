package model

// Status report of the individual orders of a bulk or multiple order that was previously received.
type SwitchOrderStatusAndReason1 struct {

	// Reference assigned to a set of orders or trades in order to link them together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Unique and unambiguous identifier for an order, as assigned by the instructing party.
	OrderReference *Max35Text `xml:"OrdrRef"`

	// Unique and unambiguous investor's identification of an order. This reference can typically be used in a hub scenario to give the reference of the order as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unique and unambiguous identifier for an order execution, as assigned by a confirming party.
	DealReference *Max35Text `xml:"DealRef,omitempty"`

	// Unique and unambiguous identifier for an order cancellation, as assigned by the instructing party.
	CancellationReference *Max35Text `xml:"CxlRef,omitempty"`

	// Status of the switch order is accepted or already executed or sent to next party or received. There is no reason attached.
	Status *OrderStatus4Code `xml:"Sts"`

	// Status of the switch order is cancelled. This status is used for an order that has been accepted or that has been entered in an order book but that can not be executed.
	Cancelled *CancelledStatus2 `xml:"Canc"`

	// Status of the switch order is conditionally accepted.
	ConditionallyAccepted *ConditionallyAcceptedStatus2 `xml:"CondlyAccptd"`

	// Status of the switch order is rejected.
	Rejected []*RejectedStatus6 `xml:"Rjctd"`

	// Status of the switch order is suspended.
	Suspended *SuspendedStatus2 `xml:"Sspd"`

	// Status of the switch order is in repair.
	InRepair *InRepairStatus2 `xml:"InRpr"`

	// Status of the switch order is partially settled.
	PartiallySettled *PartiallySettledStatus1 `xml:"PrtlySttld"`

	// Information about a switch leg that is rejected or repaired.
	LegInformation []*SwitchLegReferences1 `xml:"LegInf,omitempty"`

	// Party that initiates the status of the order.
	StatusInitiator *PartyIdentification2Choice `xml:"StsInitr,omitempty"`

	// Order data.
	OrderData *FundOrderData2 `xml:"OrdrData,omitempty"`

	// Information that has been added to the original order.
	NewDetails *ExpectedExecutionDetails2 `xml:"NewDtls,omitempty"`
}

func (s *SwitchOrderStatusAndReason1) SetMasterReference(value string) {
	s.MasterReference = (*Max35Text)(&value)
}

func (s *SwitchOrderStatusAndReason1) SetOrderReference(value string) {
	s.OrderReference = (*Max35Text)(&value)
}

func (s *SwitchOrderStatusAndReason1) SetClientReference(value string) {
	s.ClientReference = (*Max35Text)(&value)
}

func (s *SwitchOrderStatusAndReason1) SetDealReference(value string) {
	s.DealReference = (*Max35Text)(&value)
}

func (s *SwitchOrderStatusAndReason1) SetCancellationReference(value string) {
	s.CancellationReference = (*Max35Text)(&value)
}

func (s *SwitchOrderStatusAndReason1) SetStatus(value string) {
	s.Status = (*OrderStatus4Code)(&value)
}

func (s *SwitchOrderStatusAndReason1) AddCancelled() *CancelledStatus2 {
	s.Cancelled = new(CancelledStatus2)
	return s.Cancelled
}

func (s *SwitchOrderStatusAndReason1) AddConditionallyAccepted() *ConditionallyAcceptedStatus2 {
	s.ConditionallyAccepted = new(ConditionallyAcceptedStatus2)
	return s.ConditionallyAccepted
}

func (s *SwitchOrderStatusAndReason1) AddRejected() *RejectedStatus6 {
	newValue := new(RejectedStatus6)
	s.Rejected = append(s.Rejected, newValue)
	return newValue
}

func (s *SwitchOrderStatusAndReason1) AddSuspended() *SuspendedStatus2 {
	s.Suspended = new(SuspendedStatus2)
	return s.Suspended
}

func (s *SwitchOrderStatusAndReason1) AddInRepair() *InRepairStatus2 {
	s.InRepair = new(InRepairStatus2)
	return s.InRepair
}

func (s *SwitchOrderStatusAndReason1) AddPartiallySettled() *PartiallySettledStatus1 {
	s.PartiallySettled = new(PartiallySettledStatus1)
	return s.PartiallySettled
}

func (s *SwitchOrderStatusAndReason1) AddLegInformation() *SwitchLegReferences1 {
	newValue := new(SwitchLegReferences1)
	s.LegInformation = append(s.LegInformation, newValue)
	return newValue
}

func (s *SwitchOrderStatusAndReason1) AddStatusInitiator() *PartyIdentification2Choice {
	s.StatusInitiator = new(PartyIdentification2Choice)
	return s.StatusInitiator
}

func (s *SwitchOrderStatusAndReason1) AddOrderData() *FundOrderData2 {
	s.OrderData = new(FundOrderData2)
	return s.OrderData
}

func (s *SwitchOrderStatusAndReason1) AddNewDetails() *ExpectedExecutionDetails2 {
	s.NewDetails = new(ExpectedExecutionDetails2)
	return s.NewDetails
}
