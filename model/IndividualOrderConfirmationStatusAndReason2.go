package model

// Status report of the individual orders confirmation that was previously received.
type IndividualOrderConfirmationStatusAndReason2 struct {

	// Reference assigned to a set of orders or trades in order to link them together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Unique and unambiguous identifier for the order, as assigned by the instructing party.
	OrderReference *Max35Text `xml:"OrdrRef"`

	// Status of the order confirmation cancellation.
	Confirmation *ConfirmationStatus1Choice `xml:"Conf"`

	// Unique and unambiguous investor's identification of an order. This reference can typically be used in a hub scenario to give the reference of the order as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unique and unambiguous identifier for the order execution, as assigned by the confirming party.
	DealReference *Max35Text `xml:"DealRef,omitempty"`

	// Party that initiates the status of the order confirmation.
	StatusInitiator *PartyIdentification113 `xml:"StsInitr,omitempty"`

	// Account information of the individual order confirmation for which the status is given.
	InvestmentAccountDetails *InvestmentAccount58 `xml:"InvstmtAcctDtls,omitempty"`

	// Financial instrument information of the individual order confirmation for which the status is given.
	FinancialInstrumentDetails *FinancialInstrument57 `xml:"FinInstrmDtls,omitempty"`
}

func (i *IndividualOrderConfirmationStatusAndReason2) SetMasterReference(value string) {
	i.MasterReference = (*Max35Text)(&value)
}

func (i *IndividualOrderConfirmationStatusAndReason2) SetOrderReference(value string) {
	i.OrderReference = (*Max35Text)(&value)
}

func (i *IndividualOrderConfirmationStatusAndReason2) AddConfirmation() *ConfirmationStatus1Choice {
	i.Confirmation = new(ConfirmationStatus1Choice)
	return i.Confirmation
}

func (i *IndividualOrderConfirmationStatusAndReason2) SetClientReference(value string) {
	i.ClientReference = (*Max35Text)(&value)
}

func (i *IndividualOrderConfirmationStatusAndReason2) SetDealReference(value string) {
	i.DealReference = (*Max35Text)(&value)
}

func (i *IndividualOrderConfirmationStatusAndReason2) AddStatusInitiator() *PartyIdentification113 {
	i.StatusInitiator = new(PartyIdentification113)
	return i.StatusInitiator
}

func (i *IndividualOrderConfirmationStatusAndReason2) AddInvestmentAccountDetails() *InvestmentAccount58 {
	i.InvestmentAccountDetails = new(InvestmentAccount58)
	return i.InvestmentAccountDetails
}

func (i *IndividualOrderConfirmationStatusAndReason2) AddFinancialInstrumentDetails() *FinancialInstrument57 {
	i.FinancialInstrumentDetails = new(FinancialInstrument57)
	return i.FinancialInstrumentDetails
}
