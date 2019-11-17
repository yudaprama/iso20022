package model

// Status report of an individual order of a bulk or multiple or switch order cancellation instruction that was previously received.
type IndividualOrderStatusAndReason4 struct {

	// Reference assigned to a set of orders or trades in order to link them together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Unique and unambiguous identifier for an order, as assigned by the instructing party.
	OrderReference *Max35Text `xml:"OrdrRef"`

	// Unique and unambiguous investor's identification of an order. This reference can typically be used in a hub scenario to give the reference of the order as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unique and unambiguous identifier for an order cancellation, as assigned by the instructing party.
	CancellationReference *Max35Text `xml:"CxlRef,omitempty"`

	// Cancellation status of the order.
	Status *OrderCancellationStatus1Code `xml:"Sts"`

	// Status of the individual order cancellation request is rejected.
	Rejected *RejectedStatus7 `xml:"Rjctd"`

	// Party that initiates the status of the individual order cancellation.
	StatusInitiator *PartyIdentification2Choice `xml:"StsInitr,omitempty"`

	// Account information of the individual order cancellation for which the status is given.
	InvestmentAccountDetails *InvestmentAccount13 `xml:"InvstmtAcctDtls,omitempty"`

	// Financial instrument information of the individual order cancellation for which the status is given.
	FinancialInstrumentDetails *FinancialInstrument10 `xml:"FinInstrmDtls,omitempty"`
}

func (i *IndividualOrderStatusAndReason4) SetMasterReference(value string) {
	i.MasterReference = (*Max35Text)(&value)
}

func (i *IndividualOrderStatusAndReason4) SetOrderReference(value string) {
	i.OrderReference = (*Max35Text)(&value)
}

func (i *IndividualOrderStatusAndReason4) SetClientReference(value string) {
	i.ClientReference = (*Max35Text)(&value)
}

func (i *IndividualOrderStatusAndReason4) SetCancellationReference(value string) {
	i.CancellationReference = (*Max35Text)(&value)
}

func (i *IndividualOrderStatusAndReason4) SetStatus(value string) {
	i.Status = (*OrderCancellationStatus1Code)(&value)
}

func (i *IndividualOrderStatusAndReason4) AddRejected() *RejectedStatus7 {
	i.Rejected = new(RejectedStatus7)
	return i.Rejected
}

func (i *IndividualOrderStatusAndReason4) AddStatusInitiator() *PartyIdentification2Choice {
	i.StatusInitiator = new(PartyIdentification2Choice)
	return i.StatusInitiator
}

func (i *IndividualOrderStatusAndReason4) AddInvestmentAccountDetails() *InvestmentAccount13 {
	i.InvestmentAccountDetails = new(InvestmentAccount13)
	return i.InvestmentAccountDetails
}

func (i *IndividualOrderStatusAndReason4) AddFinancialInstrumentDetails() *FinancialInstrument10 {
	i.FinancialInstrumentDetails = new(FinancialInstrument10)
	return i.FinancialInstrumentDetails
}
