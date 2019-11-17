package model

// Sub-balances providing additional information on a specific position but that is not to be accounted for in the building of the aggregate balance, for example, registered.
type AdditionalBalanceInformation16 struct {

	// Reason for the sub-balance.
	SubBalanceType *SubBalanceType14Choice `xml:"SubBalTp"`

	// Quantity of securities in the sub-balance.
	Quantity *SubBalanceQuantity7Choice `xml:"Qty"`

	// Provides additional subbalance information.
	SubBalanceAdditionalDetails *RestrictedFINXMax140Text `xml:"SubBalAddtlDtls,omitempty"`
}

func (a *AdditionalBalanceInformation16) AddSubBalanceType() *SubBalanceType14Choice {
	a.SubBalanceType = new(SubBalanceType14Choice)
	return a.SubBalanceType
}

func (a *AdditionalBalanceInformation16) AddQuantity() *SubBalanceQuantity7Choice {
	a.Quantity = new(SubBalanceQuantity7Choice)
	return a.Quantity
}

func (a *AdditionalBalanceInformation16) SetSubBalanceAdditionalDetails(value string) {
	a.SubBalanceAdditionalDetails = (*RestrictedFINXMax140Text)(&value)
}
