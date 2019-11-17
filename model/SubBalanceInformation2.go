package model

// Net position of a segregated holding of a single security within the overall position held in the securities account, eg, sub-balance per status.
type SubBalanceInformation2 struct {

	// Quantity of securities in the sub-balance.
	Quantity *SubBalanceQuantity1Choice `xml:"Qty"`

	// Reason for the sub-balance.
	SubBalanceType *SecuritiesBalanceType1Code `xml:"SubBalTp"`

	// Reason a security is not available or additional information about the financial instrument for which the balance is given, for example, unregistered, registered in nominee name.
	ExtendedSubBalanceType *Extended350Code `xml:"XtndedSubBalTp"`

	// Net position of a segregated holding of a single security within the overall position held in a securities account, eg, sub-balance per status.
	AdditionalBalanceBreakdownDetails []*AdditionalBalanceInformation2 `xml:"AddtlBalBrkdwnDtls,omitempty"`
}

func (s *SubBalanceInformation2) AddQuantity() *SubBalanceQuantity1Choice {
	s.Quantity = new(SubBalanceQuantity1Choice)
	return s.Quantity
}

func (s *SubBalanceInformation2) SetSubBalanceType(value string) {
	s.SubBalanceType = (*SecuritiesBalanceType1Code)(&value)
}

func (s *SubBalanceInformation2) SetExtendedSubBalanceType(value string) {
	s.ExtendedSubBalanceType = (*Extended350Code)(&value)
}

func (s *SubBalanceInformation2) AddAdditionalBalanceBreakdownDetails() *AdditionalBalanceInformation2 {
	newValue := new(AdditionalBalanceInformation2)
	s.AdditionalBalanceBreakdownDetails = append(s.AdditionalBalanceBreakdownDetails, newValue)
	return newValue
}
