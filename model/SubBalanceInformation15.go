package model

// Net position of a segregated holding of a single security within the overall position held in the securities account, eg, sub-balance per status.
type SubBalanceInformation15 struct {

	// Reason for the sub-balance.
	SubBalanceType *SubBalanceType11Choice `xml:"SubBalTp"`

	// Quantity of securities in the sub-balance.
	Quantity *Balance9 `xml:"Qty"`

	// Provides additional subbalance information.
	SubBalanceAdditionalDetails *Max140Text `xml:"SubBalAddtlDtls,omitempty"`

	// Breakdown of the aggregate quantity reported into significant lots, for example, tax lots.
	QuantityBreakdown []*QuantityBreakdown27 `xml:"QtyBrkdwn,omitempty"`

	// Provides additional instrument sub-balance information on all or parts of the reported financial instrument (unregistered, tax exempt, etc.).
	AdditionalBalanceBreakdownDetails []*AdditionalBalanceInformation15 `xml:"AddtlBalBrkdwnDtls,omitempty"`
}

func (s *SubBalanceInformation15) AddSubBalanceType() *SubBalanceType11Choice {
	s.SubBalanceType = new(SubBalanceType11Choice)
	return s.SubBalanceType
}

func (s *SubBalanceInformation15) AddQuantity() *Balance9 {
	s.Quantity = new(Balance9)
	return s.Quantity
}

func (s *SubBalanceInformation15) SetSubBalanceAdditionalDetails(value string) {
	s.SubBalanceAdditionalDetails = (*Max140Text)(&value)
}

func (s *SubBalanceInformation15) AddQuantityBreakdown() *QuantityBreakdown27 {
	newValue := new(QuantityBreakdown27)
	s.QuantityBreakdown = append(s.QuantityBreakdown, newValue)
	return newValue
}

func (s *SubBalanceInformation15) AddAdditionalBalanceBreakdownDetails() *AdditionalBalanceInformation15 {
	newValue := new(AdditionalBalanceInformation15)
	s.AdditionalBalanceBreakdownDetails = append(s.AdditionalBalanceBreakdownDetails, newValue)
	return newValue
}
