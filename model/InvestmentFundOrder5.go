package model

// Reference of an order and order cancellation.
type InvestmentFundOrder5 struct {

	// Unique and unambiguous identifier for an order, as assigned by the instructing party.
	OrderReference *Max35Text `xml:"OrdrRef"`

	// Unique and unambiguous investor's identification of an order. This reference can typically be used in a hub scenario to give the reference of the order as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unique and unambiguous identifier for an order cancellation, as assigned by the instructing party.
	CancellationReference *Max35Text `xml:"CxlRef,omitempty"`
}

func (i *InvestmentFundOrder5) SetOrderReference(value string) {
	i.OrderReference = (*Max35Text)(&value)
}

func (i *InvestmentFundOrder5) SetClientReference(value string) {
	i.ClientReference = (*Max35Text)(&value)
}

func (i *InvestmentFundOrder5) SetCancellationReference(value string) {
	i.CancellationReference = (*Max35Text)(&value)
}
