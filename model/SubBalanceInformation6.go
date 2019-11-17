package model

// Net position of a segregated holding of a single security within the overall position held in the securities account, eg, sub-balance per status.
type SubBalanceInformation6 struct {

	// Reason for the sub-balance.
	SubBalanceType *SubBalanceType5Choice `xml:"SubBalTp"`

	// Quantity of securities in the sub-balance.
	Quantity *SubBalanceQuantity3Choice `xml:"Qty"`

	// Provides additional subbalance information.
	SubBalanceAdditionalDetails *Max140Text `xml:"SubBalAddtlDtls,omitempty"`

	// Provides additional instrument sub-balance information on all or parts of the reported financial instrument (unregistered, tax exempt, etc.).
	AdditionalBalanceBreakdownDetails []*AdditionalBalanceInformation6 `xml:"AddtlBalBrkdwnDtls,omitempty"`
}

func (s *SubBalanceInformation6) AddSubBalanceType() *SubBalanceType5Choice {
	s.SubBalanceType = new(SubBalanceType5Choice)
	return s.SubBalanceType
}

func (s *SubBalanceInformation6) AddQuantity() *SubBalanceQuantity3Choice {
	s.Quantity = new(SubBalanceQuantity3Choice)
	return s.Quantity
}

func (s *SubBalanceInformation6) SetSubBalanceAdditionalDetails(value string) {
	s.SubBalanceAdditionalDetails = (*Max140Text)(&value)
}

func (s *SubBalanceInformation6) AddAdditionalBalanceBreakdownDetails() *AdditionalBalanceInformation6 {
	newValue := new(AdditionalBalanceInformation6)
	s.AdditionalBalanceBreakdownDetails = append(s.AdditionalBalanceBreakdownDetails, newValue)
	return newValue
}
