package model

// Net position of a segregated holding of a single security within the overall position held in the securities account, for example sub-balance per status.
type SubBalanceInformation16 struct {

	// Reason for the sub-balance.
	SubBalanceType *SubBalanceType13Choice `xml:"SubBalTp"`

	// Quantity of securities in the sub-balance.
	Quantity *SubBalanceQuantity7Choice `xml:"Qty"`

	// Provides additional subbalance information.
	SubBalanceAdditionalDetails *RestrictedFINXMax140Text `xml:"SubBalAddtlDtls,omitempty"`

	// Provides additional instrument sub-balance information on all or parts of the reported financial instrument (unregistered, tax exempt, etc.).
	AdditionalBalanceBreakdownDetails []*AdditionalBalanceInformation16 `xml:"AddtlBalBrkdwnDtls,omitempty"`
}

func (s *SubBalanceInformation16) AddSubBalanceType() *SubBalanceType13Choice {
	s.SubBalanceType = new(SubBalanceType13Choice)
	return s.SubBalanceType
}

func (s *SubBalanceInformation16) AddQuantity() *SubBalanceQuantity7Choice {
	s.Quantity = new(SubBalanceQuantity7Choice)
	return s.Quantity
}

func (s *SubBalanceInformation16) SetSubBalanceAdditionalDetails(value string) {
	s.SubBalanceAdditionalDetails = (*RestrictedFINXMax140Text)(&value)
}

func (s *SubBalanceInformation16) AddAdditionalBalanceBreakdownDetails() *AdditionalBalanceInformation16 {
	newValue := new(AdditionalBalanceInformation16)
	s.AdditionalBalanceBreakdownDetails = append(s.AdditionalBalanceBreakdownDetails, newValue)
	return newValue
}
