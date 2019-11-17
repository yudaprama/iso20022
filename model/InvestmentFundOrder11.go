package model

// References of an order confirmation and an order confirmation cancellation.
type InvestmentFundOrder11 struct {

	// Unique and unambiguous identifier for the order, as assigned by the instructing party.
	OrderReference *Max35Text `xml:"OrdrRef"`

	// Unique and unambiguous investor's identification of the order. This reference can typically be used in a hub scenario to give the reference of the order as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unique and unambiguous identifier for the order execution, as assigned by the confirming party.
	DealReference *Max35Text `xml:"DealRef,omitempty"`

	// Unique and unambiguous identifier for the order confirmation cancellation, as assigned by the confirming party.
	CancellationReference *Max35Text `xml:"CxlRef,omitempty"`

	// Reason for the cancellation of the confirmation.
	CancellationReason *CancellationReason31Choice `xml:"CxlRsn,omitempty"`
}

func (i *InvestmentFundOrder11) SetOrderReference(value string) {
	i.OrderReference = (*Max35Text)(&value)
}

func (i *InvestmentFundOrder11) SetClientReference(value string) {
	i.ClientReference = (*Max35Text)(&value)
}

func (i *InvestmentFundOrder11) SetDealReference(value string) {
	i.DealReference = (*Max35Text)(&value)
}

func (i *InvestmentFundOrder11) SetCancellationReference(value string) {
	i.CancellationReference = (*Max35Text)(&value)
}

func (i *InvestmentFundOrder11) AddCancellationReason() *CancellationReason31Choice {
	i.CancellationReason = new(CancellationReason31Choice)
	return i.CancellationReason
}
