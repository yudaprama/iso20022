package model

// Status report of the individual orders of a bulk or multiple order that was previously received.
type IndividualOrderStatusAndReason2 struct {

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

	// Status of the individual order is accepted or already executed or sent to next party or received. There is no reason attached.
	Status *OrderStatus4Code `xml:"Sts"`

	// Status of the individual order is cancelled. This status is used for an order that has been accepted or that has been entered in an order book but that can not be executed.
	Cancelled *CancelledStatus2 `xml:"Canc"`

	// Status of the individual order is conditionally accepted.
	ConditionallyAccepted *ConditionallyAcceptedStatus2 `xml:"CondlyAccptd"`

	// Status of the individual order is rejected. This status is used for an order that has not been accepted or entered in an order book.
	Rejected []*RejectedStatus6 `xml:"Rjctd"`

	// Status of the individual order is suspended.
	Suspended *SuspendedStatus2 `xml:"Sspd"`

	// Status of the individual order is in repair.
	InRepair *InRepairStatus2 `xml:"InRpr"`

	// Status of the individual order is partially settled.
	PartiallySettled *PartiallySettledStatus1 `xml:"PrtlySttld"`

	// Elements from the original individual order that have been repaired so that the order can be accepted.
	RepairedConditions *RepairedConditions3 `xml:"RprdConds,omitempty"`

	// Party that initiates the status of the order cancellation.
	StatusInitiator *PartyIdentification2Choice `xml:"StsInitr,omitempty"`

	// Order data.
	OrderData *FundOrderData1 `xml:"OrdrData,omitempty"`

	// Information that has been added to the original order.
	NewDetails *ExpectedExecutionDetails2 `xml:"NewDtls,omitempty"`
}

func (i *IndividualOrderStatusAndReason2) SetMasterReference(value string) {
	i.MasterReference = (*Max35Text)(&value)
}

func (i *IndividualOrderStatusAndReason2) SetOrderReference(value string) {
	i.OrderReference = (*Max35Text)(&value)
}

func (i *IndividualOrderStatusAndReason2) SetClientReference(value string) {
	i.ClientReference = (*Max35Text)(&value)
}

func (i *IndividualOrderStatusAndReason2) SetDealReference(value string) {
	i.DealReference = (*Max35Text)(&value)
}

func (i *IndividualOrderStatusAndReason2) SetCancellationReference(value string) {
	i.CancellationReference = (*Max35Text)(&value)
}

func (i *IndividualOrderStatusAndReason2) SetStatus(value string) {
	i.Status = (*OrderStatus4Code)(&value)
}

func (i *IndividualOrderStatusAndReason2) AddCancelled() *CancelledStatus2 {
	i.Cancelled = new(CancelledStatus2)
	return i.Cancelled
}

func (i *IndividualOrderStatusAndReason2) AddConditionallyAccepted() *ConditionallyAcceptedStatus2 {
	i.ConditionallyAccepted = new(ConditionallyAcceptedStatus2)
	return i.ConditionallyAccepted
}

func (i *IndividualOrderStatusAndReason2) AddRejected() *RejectedStatus6 {
	newValue := new(RejectedStatus6)
	i.Rejected = append(i.Rejected, newValue)
	return newValue
}

func (i *IndividualOrderStatusAndReason2) AddSuspended() *SuspendedStatus2 {
	i.Suspended = new(SuspendedStatus2)
	return i.Suspended
}

func (i *IndividualOrderStatusAndReason2) AddInRepair() *InRepairStatus2 {
	i.InRepair = new(InRepairStatus2)
	return i.InRepair
}

func (i *IndividualOrderStatusAndReason2) AddPartiallySettled() *PartiallySettledStatus1 {
	i.PartiallySettled = new(PartiallySettledStatus1)
	return i.PartiallySettled
}

func (i *IndividualOrderStatusAndReason2) AddRepairedConditions() *RepairedConditions3 {
	i.RepairedConditions = new(RepairedConditions3)
	return i.RepairedConditions
}

func (i *IndividualOrderStatusAndReason2) AddStatusInitiator() *PartyIdentification2Choice {
	i.StatusInitiator = new(PartyIdentification2Choice)
	return i.StatusInitiator
}

func (i *IndividualOrderStatusAndReason2) AddOrderData() *FundOrderData1 {
	i.OrderData = new(FundOrderData1)
	return i.OrderData
}

func (i *IndividualOrderStatusAndReason2) AddNewDetails() *ExpectedExecutionDetails2 {
	i.NewDetails = new(ExpectedExecutionDetails2)
	return i.NewDetails
}
