package model

// Identifies an order linked to an account opening.
type InvestmentFundOrder4 struct {

	// Unique and unambiguous identifier for an order, as assigned by the instructing party.
	OrderReference *Max35Text `xml:"OrdrRef,omitempty"`

	// Unique and unambiguous identifier for a group of individual orders, as assigned by the instructing party. This identifier links the individual orders together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`
}

func (i *InvestmentFundOrder4) SetOrderReference(value string) {
	i.OrderReference = (*Max35Text)(&value)
}

func (i *InvestmentFundOrder4) SetMasterReference(value string) {
	i.MasterReference = (*Max35Text)(&value)
}
