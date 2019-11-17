package model

// Subbalances providing additional information on a specific position but that is not to be accounted for in the building of the aggregate balance, for example, registered.
type AdditionalBalanceInformation15 struct {

	// Reason for the sub-balance.
	SubBalanceType *SubBalanceType12Choice `xml:"SubBalTp"`

	// Quantity of securities in the sub-balance.
	Quantity *Balance9 `xml:"Qty"`

	// Provides additional sub-balance information.
	SubBalanceAdditionalDetails *Max140Text `xml:"SubBalAddtlDtls,omitempty"`

	// Breakdown of the aggregate quantity reported into significant lots, for example, tax lots.
	QuantityBreakdown []*QuantityBreakdown27 `xml:"QtyBrkdwn,omitempty"`
}

func (a *AdditionalBalanceInformation15) AddSubBalanceType() *SubBalanceType12Choice {
	a.SubBalanceType = new(SubBalanceType12Choice)
	return a.SubBalanceType
}

func (a *AdditionalBalanceInformation15) AddQuantity() *Balance9 {
	a.Quantity = new(Balance9)
	return a.Quantity
}

func (a *AdditionalBalanceInformation15) SetSubBalanceAdditionalDetails(value string) {
	a.SubBalanceAdditionalDetails = (*Max140Text)(&value)
}

func (a *AdditionalBalanceInformation15) AddQuantityBreakdown() *QuantityBreakdown27 {
	newValue := new(QuantityBreakdown27)
	a.QuantityBreakdown = append(a.QuantityBreakdown, newValue)
	return newValue
}
