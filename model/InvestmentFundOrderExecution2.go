package model

// Reference of an order, client or deal reference.
type InvestmentFundOrderExecution2 struct {

	// Unique and unambiguous identifier for an order, as assigned by the instructing party.
	OrderReference *Max35Text `xml:"OrdrRef,omitempty"`

	// Unique and unambiguous investor's identification of an order. This reference can typically be used in a hub scenario to give the reference of the order as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unique and unambiguous identifier for an order cancellation, as assigned by the instructing party.
	DealReference *Max35Text `xml:"DealRef,omitempty"`
}

func (i *InvestmentFundOrderExecution2) SetOrderReference(value string) {
	i.OrderReference = (*Max35Text)(&value)
}

func (i *InvestmentFundOrderExecution2) SetClientReference(value string) {
	i.ClientReference = (*Max35Text)(&value)
}

func (i *InvestmentFundOrderExecution2) SetDealReference(value string) {
	i.DealReference = (*Max35Text)(&value)
}
