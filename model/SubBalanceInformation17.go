package model

// Net position of a segregated holding of a single security within the overall position held in the securities account, eg, sub-balance per status.
type SubBalanceInformation17 struct {

	// Reason for the sub-balance.
	SubBalanceType *SubBalanceType13Choice `xml:"SubBalTp"`

	// Quantity of securities in the sub-balance.
	Quantity *Balance13 `xml:"Qty"`

	// Provides additional subbalance information.
	SubBalanceAdditionalDetails *RestrictedFINXMax140Text `xml:"SubBalAddtlDtls,omitempty"`

	// Breakdown of the aggregate quantity reported into significant lots, for example, tax lots.
	QuantityBreakdown []*QuantityBreakdown40 `xml:"QtyBrkdwn,omitempty"`

	// Provides additional instrument sub-balance information on all or parts of the reported financial instrument (unregistered, tax exempt, etc.).
	AdditionalBalanceBreakdownDetails []*AdditionalBalanceInformation17 `xml:"AddtlBalBrkdwnDtls,omitempty"`
}

func (s *SubBalanceInformation17) AddSubBalanceType() *SubBalanceType13Choice {
	s.SubBalanceType = new(SubBalanceType13Choice)
	return s.SubBalanceType
}

func (s *SubBalanceInformation17) AddQuantity() *Balance13 {
	s.Quantity = new(Balance13)
	return s.Quantity
}

func (s *SubBalanceInformation17) SetSubBalanceAdditionalDetails(value string) {
	s.SubBalanceAdditionalDetails = (*RestrictedFINXMax140Text)(&value)
}

func (s *SubBalanceInformation17) AddQuantityBreakdown() *QuantityBreakdown40 {
	newValue := new(QuantityBreakdown40)
	s.QuantityBreakdown = append(s.QuantityBreakdown, newValue)
	return newValue
}

func (s *SubBalanceInformation17) AddAdditionalBalanceBreakdownDetails() *AdditionalBalanceInformation17 {
	newValue := new(AdditionalBalanceInformation17)
	s.AdditionalBalanceBreakdownDetails = append(s.AdditionalBalanceBreakdownDetails, newValue)
	return newValue
}
