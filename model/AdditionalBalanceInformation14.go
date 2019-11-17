package model

// Sub-balances providing additional information on a specific position but that is not to be accounted for in the building of the aggregate balance, for example, registered.
type AdditionalBalanceInformation14 struct {

	// Reason for the sub-balance.
	SubBalanceType *SubBalanceType12Choice `xml:"SubBalTp"`

	// Quantity of securities in the sub-balance.
	Quantity *SubBalanceQuantity6Choice `xml:"Qty"`

	// Provides additional subbalance information.
	SubBalanceAdditionalDetails *Max140Text `xml:"SubBalAddtlDtls,omitempty"`
}

func (a *AdditionalBalanceInformation14) AddSubBalanceType() *SubBalanceType12Choice {
	a.SubBalanceType = new(SubBalanceType12Choice)
	return a.SubBalanceType
}

func (a *AdditionalBalanceInformation14) AddQuantity() *SubBalanceQuantity6Choice {
	a.Quantity = new(SubBalanceQuantity6Choice)
	return a.Quantity
}

func (a *AdditionalBalanceInformation14) SetSubBalanceAdditionalDetails(value string) {
	a.SubBalanceAdditionalDetails = (*Max140Text)(&value)
}
