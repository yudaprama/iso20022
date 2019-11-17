package model

// Reference of an order, order cancellation and master reference.
type InvestmentFundOrder1 struct {

	// Reference assigned to a set of orders or trades in order to link them together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Reference of an order and order cancellation.
	OrderReferences []*InvestmentFundOrder5 `xml:"OrdrRefs"`
}

func (i *InvestmentFundOrder1) SetMasterReference(value string) {
	i.MasterReference = (*Max35Text)(&value)
}

func (i *InvestmentFundOrder1) AddOrderReferences() *InvestmentFundOrder5 {
	newValue := new(InvestmentFundOrder5)
	i.OrderReferences = append(i.OrderReferences, newValue)
	return newValue
}
