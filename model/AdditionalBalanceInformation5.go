package model

// Subbalances providing additional information on a specific position but that is not to be accounted for in the building of the aggregate balance, for example, registered.
type AdditionalBalanceInformation5 struct {

	// Reason for the sub-balance.
	SubBalanceType *SubBalanceType6Choice `xml:"SubBalTp"`

	// Quantity of securities in the sub-balance.
	Quantity *SubBalanceQuantity3Choice `xml:"Qty"`

	// Provides additional subbalance information.
	SubBalanceAdditionalDetails *Max140Text `xml:"SubBalAddtlDtls,omitempty"`

	// Breakdown of the aggregate quantity reported into significant lots, for example, tax lots.
	QuantityBreakdown []*QuantityBreakdown3 `xml:"QtyBrkdwn,omitempty"`
}

func (a *AdditionalBalanceInformation5) AddSubBalanceType() *SubBalanceType6Choice {
	a.SubBalanceType = new(SubBalanceType6Choice)
	return a.SubBalanceType
}

func (a *AdditionalBalanceInformation5) AddQuantity() *SubBalanceQuantity3Choice {
	a.Quantity = new(SubBalanceQuantity3Choice)
	return a.Quantity
}

func (a *AdditionalBalanceInformation5) SetSubBalanceAdditionalDetails(value string) {
	a.SubBalanceAdditionalDetails = (*Max140Text)(&value)
}

func (a *AdditionalBalanceInformation5) AddQuantityBreakdown() *QuantityBreakdown3 {
	newValue := new(QuantityBreakdown3)
	a.QuantityBreakdown = append(a.QuantityBreakdown, newValue)
	return newValue
}
