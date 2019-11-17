package model

// Subbalances providing additional information on a specific position but that is not to be accounted for in the building of the aggregate balance, for example, registered.
type AdditionalBalanceInformation17 struct {

	// Reason for the sub-balance.
	SubBalanceType *SubBalanceType14Choice `xml:"SubBalTp"`

	// Quantity of securities in the sub-balance.
	Quantity *Balance13 `xml:"Qty"`

	// Provides additional sub-balance information.
	SubBalanceAdditionalDetails *RestrictedFINXMax140Text `xml:"SubBalAddtlDtls,omitempty"`

	// Breakdown of the aggregate quantity reported into significant lots, for example, tax lots.
	QuantityBreakdown []*QuantityBreakdown40 `xml:"QtyBrkdwn,omitempty"`
}

func (a *AdditionalBalanceInformation17) AddSubBalanceType() *SubBalanceType14Choice {
	a.SubBalanceType = new(SubBalanceType14Choice)
	return a.SubBalanceType
}

func (a *AdditionalBalanceInformation17) AddQuantity() *Balance13 {
	a.Quantity = new(Balance13)
	return a.Quantity
}

func (a *AdditionalBalanceInformation17) SetSubBalanceAdditionalDetails(value string) {
	a.SubBalanceAdditionalDetails = (*RestrictedFINXMax140Text)(&value)
}

func (a *AdditionalBalanceInformation17) AddQuantityBreakdown() *QuantityBreakdown40 {
	newValue := new(QuantityBreakdown40)
	a.QuantityBreakdown = append(a.QuantityBreakdown, newValue)
	return newValue
}
