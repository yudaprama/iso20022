package model

// Status report of the individual orders of a bulk or multiple order that was previously received.
type IndividualOrderStatusAndReason1 struct {

	// Status of the order is accepted or already executed or sent to next party or received. There is no reason attached.
	Status *OrderStatus2Code `xml:"Sts"`

	// Status of the individual order details is cancelled.
	Cancelled *CancelledStatus1 `xml:"Canc"`

	// Status of the individual order details is conditionally accepted.
	ConditionallyAccepted *ConditionallyAcceptedStatus1 `xml:"CondlyAccptd"`

	// Status of the individual order details is in repair.
	InRepair *InRepairStatus1 `xml:"InRpr"`

	// Status of the individual order details is rejected.
	Rejected *RejectedStatus3 `xml:"Rjctd"`

	// Status of the individual order details is suspended.
	Suspended *SuspendedStatus1 `xml:"Sspd"`

	// Elements from the original individual order details that have been repaired so that the order can be accepted.
	RepairedConditions *RepairedConditions2 `xml:"RprdConds"`

	// Unique and unambiguous identifier for an order, as assigned by the instructing party.
	OrderReference *Max35Text `xml:"OrdrRef"`

	// Party that initiates the status of the order.
	StatusInitiator *PartyIdentification2Choice `xml:"StsInitr,omitempty"`

	// Choice between the investment account and the financial instrument.
	InvestmentAccountOrFinancialInstrument *InvestmentAccountOrFinancialInstrument1Choice `xml:"InvstmtAcctOrFinInstrm,omitempty"`

	// Information that has been added to the original order.
	NewDetails *ExpectedExecutionDetails1 `xml:"NewDtls,omitempty"`
}

func (i *IndividualOrderStatusAndReason1) SetStatus(value string) {
	i.Status = (*OrderStatus2Code)(&value)
}

func (i *IndividualOrderStatusAndReason1) AddCancelled() *CancelledStatus1 {
	i.Cancelled = new(CancelledStatus1)
	return i.Cancelled
}

func (i *IndividualOrderStatusAndReason1) AddConditionallyAccepted() *ConditionallyAcceptedStatus1 {
	i.ConditionallyAccepted = new(ConditionallyAcceptedStatus1)
	return i.ConditionallyAccepted
}

func (i *IndividualOrderStatusAndReason1) AddInRepair() *InRepairStatus1 {
	i.InRepair = new(InRepairStatus1)
	return i.InRepair
}

func (i *IndividualOrderStatusAndReason1) AddRejected() *RejectedStatus3 {
	i.Rejected = new(RejectedStatus3)
	return i.Rejected
}

func (i *IndividualOrderStatusAndReason1) AddSuspended() *SuspendedStatus1 {
	i.Suspended = new(SuspendedStatus1)
	return i.Suspended
}

func (i *IndividualOrderStatusAndReason1) AddRepairedConditions() *RepairedConditions2 {
	i.RepairedConditions = new(RepairedConditions2)
	return i.RepairedConditions
}

func (i *IndividualOrderStatusAndReason1) SetOrderReference(value string) {
	i.OrderReference = (*Max35Text)(&value)
}

func (i *IndividualOrderStatusAndReason1) AddStatusInitiator() *PartyIdentification2Choice {
	i.StatusInitiator = new(PartyIdentification2Choice)
	return i.StatusInitiator
}

func (i *IndividualOrderStatusAndReason1) AddInvestmentAccountOrFinancialInstrument() *InvestmentAccountOrFinancialInstrument1Choice {
	i.InvestmentAccountOrFinancialInstrument = new(InvestmentAccountOrFinancialInstrument1Choice)
	return i.InvestmentAccountOrFinancialInstrument
}

func (i *IndividualOrderStatusAndReason1) AddNewDetails() *ExpectedExecutionDetails1 {
	i.NewDetails = new(ExpectedExecutionDetails1)
	return i.NewDetails
}
