package model

// References of an order and order cancellation.
type InvestmentFundOrder9 struct {

	// Unique and unambiguous identifier for the order, as assigned by the instructing party.
	OrderReference *Max35Text `xml:"OrdrRef"`

	// Unique and unambiguous investor's identification of the order. This reference can typically be used in a hub scenario to give the reference of the order as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unique and unambiguous identifier for the order cancellation, as assigned by the instructing party.
	CancellationReference *Max35Text `xml:"CxlRef,omitempty"`

	// Reason for the cancellation.
	CancellationReason *CancellationReason32Choice `xml:"CxlRsn,omitempty"`
}

func (i *InvestmentFundOrder9) SetOrderReference(value string) {
	i.OrderReference = (*Max35Text)(&value)
}

func (i *InvestmentFundOrder9) SetClientReference(value string) {
	i.ClientReference = (*Max35Text)(&value)
}

func (i *InvestmentFundOrder9) SetCancellationReference(value string) {
	i.CancellationReference = (*Max35Text)(&value)
}

func (i *InvestmentFundOrder9) AddCancellationReason() *CancellationReason32Choice {
	i.CancellationReason = new(CancellationReason32Choice)
	return i.CancellationReason
}
