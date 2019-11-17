package model

// Net position of a segregated holding of a single security within the overall position held in the securities account, eg, sub-balance per status.
type SubBalanceInformation1 struct {

	// Quantity of securities in the sub-balance.
	Quantity *SubBalanceQuantity1Choice `xml:"Qty"`

	// Reason for the sub-balance.
	SubBalanceType *SecuritiesBalanceType1Choice `xml:"SubBalTp"`

	// Net position of a segregated holding of a single security within the overall position held in a securities account, eg, sub-balance per status.
	AdditionalBalanceBreakdownDetails []*AdditionalBalanceInformation `xml:"AddtlBalBrkdwnDtls,omitempty"`
}

func (s *SubBalanceInformation1) AddQuantity() *SubBalanceQuantity1Choice {
	s.Quantity = new(SubBalanceQuantity1Choice)
	return s.Quantity
}

func (s *SubBalanceInformation1) AddSubBalanceType() *SecuritiesBalanceType1Choice {
	s.SubBalanceType = new(SecuritiesBalanceType1Choice)
	return s.SubBalanceType
}

func (s *SubBalanceInformation1) AddAdditionalBalanceBreakdownDetails() *AdditionalBalanceInformation {
	newValue := new(AdditionalBalanceInformation)
	s.AdditionalBalanceBreakdownDetails = append(s.AdditionalBalanceBreakdownDetails, newValue)
	return newValue
}
