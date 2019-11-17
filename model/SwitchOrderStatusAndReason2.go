package model

// Status report of the individual orders of a bulk or multiple order that was previously received.
type SwitchOrderStatusAndReason2 struct {

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

	// Status of the switch order.
	OrderStatus *OrderStatus4Choice `xml:"OrdrSts"`

	// Information about a switch leg that is rejected or repaired.
	LegInformation []*SwitchLegReferences2 `xml:"LegInf,omitempty"`

	// Party that initiates the status of the order.
	StatusInitiator *PartyIdentification113 `xml:"StsInitr,omitempty"`

	// Order data.
	OrderData *FundOrderData6 `xml:"OrdrData,omitempty"`

	// Expected execution information.
	NewDetails *ExpectedExecutionDetails2 `xml:"NewDtls,omitempty"`
}

func (s *SwitchOrderStatusAndReason2) SetMasterReference(value string) {
	s.MasterReference = (*Max35Text)(&value)
}

func (s *SwitchOrderStatusAndReason2) SetOrderReference(value string) {
	s.OrderReference = (*Max35Text)(&value)
}

func (s *SwitchOrderStatusAndReason2) SetClientReference(value string) {
	s.ClientReference = (*Max35Text)(&value)
}

func (s *SwitchOrderStatusAndReason2) SetDealReference(value string) {
	s.DealReference = (*Max35Text)(&value)
}

func (s *SwitchOrderStatusAndReason2) SetCancellationReference(value string) {
	s.CancellationReference = (*Max35Text)(&value)
}

func (s *SwitchOrderStatusAndReason2) AddOrderStatus() *OrderStatus4Choice {
	s.OrderStatus = new(OrderStatus4Choice)
	return s.OrderStatus
}

func (s *SwitchOrderStatusAndReason2) AddLegInformation() *SwitchLegReferences2 {
	newValue := new(SwitchLegReferences2)
	s.LegInformation = append(s.LegInformation, newValue)
	return newValue
}

func (s *SwitchOrderStatusAndReason2) AddStatusInitiator() *PartyIdentification113 {
	s.StatusInitiator = new(PartyIdentification113)
	return s.StatusInitiator
}

func (s *SwitchOrderStatusAndReason2) AddOrderData() *FundOrderData6 {
	s.OrderData = new(FundOrderData6)
	return s.OrderData
}

func (s *SwitchOrderStatusAndReason2) AddNewDetails() *ExpectedExecutionDetails2 {
	s.NewDetails = new(ExpectedExecutionDetails2)
	return s.NewDetails
}
