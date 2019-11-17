package model

// Net position of a segregated holding of a single security within the overall position held in the securities account, eg, sub-balance per status.
type SubBalanceInformation9 struct {

	// Reason for the sub-balance.
	SubBalanceType *SubBalanceType5Choice `xml:"SubBalTp"`

	// Quantity of securities in the sub-balance.
	Quantity *SubBalanceQuantity3Choice `xml:"Qty"`

	// Provides additional subbalance information.
	SubBalanceAdditionalDetails *Max140Text `xml:"SubBalAddtlDtls,omitempty"`

	// Breakdown of the aggregate quantity reported into significant lots, for example, tax lots.
	QuantityBreakdown []*QuantityBreakdown13 `xml:"QtyBrkdwn,omitempty"`

	// Provides additional instrument sub-balance information on all or parts of the reported financial instrument (unregistered, tax exempt, etc.).
	AdditionalBalanceBreakdownDetails []*AdditionalBalanceInformation9 `xml:"AddtlBalBrkdwnDtls,omitempty"`
}

func (s *SubBalanceInformation9) AddSubBalanceType() *SubBalanceType5Choice {
	s.SubBalanceType = new(SubBalanceType5Choice)
	return s.SubBalanceType
}

func (s *SubBalanceInformation9) AddQuantity() *SubBalanceQuantity3Choice {
	s.Quantity = new(SubBalanceQuantity3Choice)
	return s.Quantity
}

func (s *SubBalanceInformation9) SetSubBalanceAdditionalDetails(value string) {
	s.SubBalanceAdditionalDetails = (*Max140Text)(&value)
}

func (s *SubBalanceInformation9) AddQuantityBreakdown() *QuantityBreakdown13 {
	newValue := new(QuantityBreakdown13)
	s.QuantityBreakdown = append(s.QuantityBreakdown, newValue)
	return newValue
}

func (s *SubBalanceInformation9) AddAdditionalBalanceBreakdownDetails() *AdditionalBalanceInformation9 {
	newValue := new(AdditionalBalanceInformation9)
	s.AdditionalBalanceBreakdownDetails = append(s.AdditionalBalanceBreakdownDetails, newValue)
	return newValue
}
