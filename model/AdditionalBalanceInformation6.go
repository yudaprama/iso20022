package model

// Subbalances providing additional information on a specific position but that is not to be accounted for in the building of the aggregate balance, for example, registered.
type AdditionalBalanceInformation6 struct {

	// Reason for the sub-balance.
	SubBalanceType *SubBalanceType6Choice `xml:"SubBalTp"`

	// Quantity of securities in the sub-balance.
	Quantity *SubBalanceQuantity3Choice `xml:"Qty"`

	// Provides additional subbalance information.
	SubBalanceAdditionalDetails *Max140Text `xml:"SubBalAddtlDtls,omitempty"`
}

func (a *AdditionalBalanceInformation6) AddSubBalanceType() *SubBalanceType6Choice {
	a.SubBalanceType = new(SubBalanceType6Choice)
	return a.SubBalanceType
}

func (a *AdditionalBalanceInformation6) AddQuantity() *SubBalanceQuantity3Choice {
	a.Quantity = new(SubBalanceQuantity3Choice)
	return a.Quantity
}

func (a *AdditionalBalanceInformation6) SetSubBalanceAdditionalDetails(value string) {
	a.SubBalanceAdditionalDetails = (*Max140Text)(&value)
}
