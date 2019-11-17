package model

// Status report of the individual orders of a bulk or multiple order that was previously received.
type IndividualOrderStatusAndReason7 struct {

	// Reference assigned to a set of orders or trades in order to link them together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Unique and unambiguous identifier for the order, as assigned by the instructing party.
	OrderReference *Max35Text `xml:"OrdrRef"`

	// Unique and unambiguous investor's identification of the order. This reference can typically be used in a hub scenario to give the reference of the order as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unique and unambiguous identifier for the order execution, as assigned by the confirming party.
	DealReference *Max35Text `xml:"DealRef,omitempty"`

	// Unique and unambiguous identifier for the order cancellation, as assigned by the instructing party.
	CancellationReference *Max35Text `xml:"CxlRef,omitempty"`

	// Status of the individual order.
	OrderStatus *OrderStatus5Choice `xml:"OrdrSts"`

	// Elements from the original individual order that have been repaired so that the order can be accepted.
	RepairedFee []*Fee3 `xml:"RprdFee,omitempty"`

	// Party that initiates the status of the order.
	StatusInitiator *PartyIdentification113 `xml:"StsInitr,omitempty"`

	// Order data.
	OrderData *FundOrderData5 `xml:"OrdrData,omitempty"`

	// Expected execution information.
	NewDetails *ExpectedExecutionDetails4 `xml:"NewDtls,omitempty"`

	// Information about gating and hold back of redemption proceeds.
	GatingOrHoldBackDetails *HoldBackInformation3 `xml:"GtgOrHldBckDtls,omitempty"`
}

func (i *IndividualOrderStatusAndReason7) SetMasterReference(value string) {
	i.MasterReference = (*Max35Text)(&value)
}

func (i *IndividualOrderStatusAndReason7) SetOrderReference(value string) {
	i.OrderReference = (*Max35Text)(&value)
}

func (i *IndividualOrderStatusAndReason7) SetClientReference(value string) {
	i.ClientReference = (*Max35Text)(&value)
}

func (i *IndividualOrderStatusAndReason7) SetDealReference(value string) {
	i.DealReference = (*Max35Text)(&value)
}

func (i *IndividualOrderStatusAndReason7) SetCancellationReference(value string) {
	i.CancellationReference = (*Max35Text)(&value)
}

func (i *IndividualOrderStatusAndReason7) AddOrderStatus() *OrderStatus5Choice {
	i.OrderStatus = new(OrderStatus5Choice)
	return i.OrderStatus
}

func (i *IndividualOrderStatusAndReason7) AddRepairedFee() *Fee3 {
	newValue := new(Fee3)
	i.RepairedFee = append(i.RepairedFee, newValue)
	return newValue
}

func (i *IndividualOrderStatusAndReason7) AddStatusInitiator() *PartyIdentification113 {
	i.StatusInitiator = new(PartyIdentification113)
	return i.StatusInitiator
}

func (i *IndividualOrderStatusAndReason7) AddOrderData() *FundOrderData5 {
	i.OrderData = new(FundOrderData5)
	return i.OrderData
}

func (i *IndividualOrderStatusAndReason7) AddNewDetails() *ExpectedExecutionDetails4 {
	i.NewDetails = new(ExpectedExecutionDetails4)
	return i.NewDetails
}

func (i *IndividualOrderStatusAndReason7) AddGatingOrHoldBackDetails() *HoldBackInformation3 {
	i.GatingOrHoldBackDetails = new(HoldBackInformation3)
	return i.GatingOrHoldBackDetails
}
