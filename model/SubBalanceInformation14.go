package model

// Net position of a segregated holding of a single security within the overall position held in the securities account, for example sub-balance per status.
type SubBalanceInformation14 struct {

	// Reason for the sub-balance.
	SubBalanceType *SubBalanceType11Choice `xml:"SubBalTp"`

	// Quantity of securities in the sub-balance.
	Quantity *SubBalanceQuantity6Choice `xml:"Qty"`

	// Provides additional subbalance information.
	SubBalanceAdditionalDetails *Max140Text `xml:"SubBalAddtlDtls,omitempty"`

	// Provides additional instrument sub-balance information on all or parts of the reported financial instrument (unregistered, tax exempt, etc.).
	AdditionalBalanceBreakdownDetails []*AdditionalBalanceInformation14 `xml:"AddtlBalBrkdwnDtls,omitempty"`
}

func (s *SubBalanceInformation14) AddSubBalanceType() *SubBalanceType11Choice {
	s.SubBalanceType = new(SubBalanceType11Choice)
	return s.SubBalanceType
}

func (s *SubBalanceInformation14) AddQuantity() *SubBalanceQuantity6Choice {
	s.Quantity = new(SubBalanceQuantity6Choice)
	return s.Quantity
}

func (s *SubBalanceInformation14) SetSubBalanceAdditionalDetails(value string) {
	s.SubBalanceAdditionalDetails = (*Max140Text)(&value)
}

func (s *SubBalanceInformation14) AddAdditionalBalanceBreakdownDetails() *AdditionalBalanceInformation14 {
	newValue := new(AdditionalBalanceInformation14)
	s.AdditionalBalanceBreakdownDetails = append(s.AdditionalBalanceBreakdownDetails, newValue)
	return newValue
}
