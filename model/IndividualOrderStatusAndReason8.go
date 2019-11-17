package model

// Status report of an individual order of a bulk or multiple or switch order cancellation instruction that was previously received.
type IndividualOrderStatusAndReason8 struct {

	// Reference assigned to a set of orders or trades in order to link them together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Unique and unambiguous identifier for the order, as assigned by the instructing party.
	OrderReference *Max35Text `xml:"OrdrRef"`

	// Unique and unambiguous investor's identification of the order. This reference can typically be used in a hub scenario to give the reference of the order as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unique and unambiguous identifier for the order cancellation, as assigned by the instructing party.
	CancellationReference *Max35Text `xml:"CxlRef,omitempty"`

	// Cancellation status of the order cancellation.
	CancellationStatus *CancellationStatus22Choice `xml:"CxlSts"`

	// Party that initiates the status of the individual order cancellation.
	StatusInitiator *PartyIdentification113 `xml:"StsInitr,omitempty"`

	// Account information of the individual order cancellation for which the status is given.
	InvestmentAccountDetails *InvestmentAccount58 `xml:"InvstmtAcctDtls,omitempty"`

	// Financial instrument information of the individual order cancellation for which the status is given.
	FinancialInstrumentDetails *FinancialInstrument57 `xml:"FinInstrmDtls,omitempty"`
}

func (i *IndividualOrderStatusAndReason8) SetMasterReference(value string) {
	i.MasterReference = (*Max35Text)(&value)
}

func (i *IndividualOrderStatusAndReason8) SetOrderReference(value string) {
	i.OrderReference = (*Max35Text)(&value)
}

func (i *IndividualOrderStatusAndReason8) SetClientReference(value string) {
	i.ClientReference = (*Max35Text)(&value)
}

func (i *IndividualOrderStatusAndReason8) SetCancellationReference(value string) {
	i.CancellationReference = (*Max35Text)(&value)
}

func (i *IndividualOrderStatusAndReason8) AddCancellationStatus() *CancellationStatus22Choice {
	i.CancellationStatus = new(CancellationStatus22Choice)
	return i.CancellationStatus
}

func (i *IndividualOrderStatusAndReason8) AddStatusInitiator() *PartyIdentification113 {
	i.StatusInitiator = new(PartyIdentification113)
	return i.StatusInitiator
}

func (i *IndividualOrderStatusAndReason8) AddInvestmentAccountDetails() *InvestmentAccount58 {
	i.InvestmentAccountDetails = new(InvestmentAccount58)
	return i.InvestmentAccountDetails
}

func (i *IndividualOrderStatusAndReason8) AddFinancialInstrumentDetails() *FinancialInstrument57 {
	i.FinancialInstrumentDetails = new(FinancialInstrument57)
	return i.FinancialInstrumentDetails
}
