package model

// Reference of an order and of an order cancellation.
type InvestmentFundOrder2 struct {

	// Reference assigned to a set of orders or trades in order to link them together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Unique and unambiguous identifier for an order, as assigned by the instructing party.
	OrderReference *Max35Text `xml:"OrdrRef"`

	// Unique and unambiguous investor's identification of an order. This reference can typically be used in a hub scenario to give the reference of the order as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unique and unambiguous identifier for an order cancellation, as assigned by the instructing party.
	CancellationReference *Max35Text `xml:"CxlRef,omitempty"`

	// Unique and unambiguous identifier for an order execution, as assigned by a confirming party.
	DealReference *Max35Text `xml:"DealRef,omitempty"`

	// Account information of the individual order instruction or individual order cancellation request for which the status is requested.
	InvestmentAccountDetails *InvestmentAccount13 `xml:"InvstmtAcctDtls,omitempty"`

	// Financial instrument information of the individual order or individual order cancellation request for which the status is requested.
	FinancialInstrumentDetails *FinancialInstrument10 `xml:"FinInstrmDtls,omitempty"`
}

func (i *InvestmentFundOrder2) SetMasterReference(value string) {
	i.MasterReference = (*Max35Text)(&value)
}

func (i *InvestmentFundOrder2) SetOrderReference(value string) {
	i.OrderReference = (*Max35Text)(&value)
}

func (i *InvestmentFundOrder2) SetClientReference(value string) {
	i.ClientReference = (*Max35Text)(&value)
}

func (i *InvestmentFundOrder2) SetCancellationReference(value string) {
	i.CancellationReference = (*Max35Text)(&value)
}

func (i *InvestmentFundOrder2) SetDealReference(value string) {
	i.DealReference = (*Max35Text)(&value)
}

func (i *InvestmentFundOrder2) AddInvestmentAccountDetails() *InvestmentAccount13 {
	i.InvestmentAccountDetails = new(InvestmentAccount13)
	return i.InvestmentAccountDetails
}

func (i *InvestmentFundOrder2) AddFinancialInstrumentDetails() *FinancialInstrument10 {
	i.FinancialInstrumentDetails = new(FinancialInstrument10)
	return i.FinancialInstrumentDetails
}
