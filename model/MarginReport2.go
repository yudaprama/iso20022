package model

// Provides the details on the margin report per margin account, and optionally per non-clearing member.
type MarginReport2 struct {

	// Specifies if the margin is related to equities or fixed income.
	MarginProduct []*MarginProductType1Choice `xml:"MrgnPdct,omitempty"`

	// Identifies the clearing member's account.
	MarginAccount *SecuritiesAccount18 `xml:"MrgnAcct"`

	// Used to indicate whether the reported margin account is collateralised or not. If not collateralised, the account is configured for informational reporting.
	CollateralisedMarginAccountIndicator *YesNoIndicator `xml:"CollsdMrgnAcctInd,omitempty"`

	// Provides details about the non clearing member identification and account.
	NonClearingMember []*PartyIdentificationAndAccount31 `xml:"NonClrMmb,omitempty"`

	// Provides the margin calculation summary per margin account.
	MarginCalculationSummary *MarginCalculation1 `xml:"MrgnClctnSummry,omitempty"`

	// Provides the margin details such as the exposure amount and the initial margin.
	MarginCalculation []*MarginCalculation2 `xml:"MrgnClctn"`
}

func (m *MarginReport2) AddMarginProduct() *MarginProductType1Choice {
	newValue := new(MarginProductType1Choice)
	m.MarginProduct = append(m.MarginProduct, newValue)
	return newValue
}

func (m *MarginReport2) AddMarginAccount() *SecuritiesAccount18 {
	m.MarginAccount = new(SecuritiesAccount18)
	return m.MarginAccount
}

func (m *MarginReport2) SetCollateralisedMarginAccountIndicator(value string) {
	m.CollateralisedMarginAccountIndicator = (*YesNoIndicator)(&value)
}

func (m *MarginReport2) AddNonClearingMember() *PartyIdentificationAndAccount31 {
	newValue := new(PartyIdentificationAndAccount31)
	m.NonClearingMember = append(m.NonClearingMember, newValue)
	return newValue
}

func (m *MarginReport2) AddMarginCalculationSummary() *MarginCalculation1 {
	m.MarginCalculationSummary = new(MarginCalculation1)
	return m.MarginCalculationSummary
}

func (m *MarginReport2) AddMarginCalculation() *MarginCalculation2 {
	newValue := new(MarginCalculation2)
	m.MarginCalculation = append(m.MarginCalculation, newValue)
	return newValue
}
