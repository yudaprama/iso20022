package model

// Card used to represent a financial account for the purpose of payment settlement.
type FinancialCard1 struct {

	// Monetary value of the credit limit for this financial card.
	CreditLimitAmount []*CurrencyAndAmount `xml:"CdtLmtAmt,omitempty"`

	// Monetary value of the credit available for this financial card.
	CreditAvailableAmount []*CurrencyAndAmount `xml:"CdtAvlblAmt,omitempty"`

	// Interest rate expressed as a percentage for this financial card.
	InterestRatePercent *PercentageRate `xml:"IntrstRatePct,omitempty"`
}

func (f *FinancialCard1) AddCreditLimitAmount(value, currency string) {
	f.CreditLimitAmount = append(f.CreditLimitAmount, NewCurrencyAndAmount(value, currency))
}

func (f *FinancialCard1) AddCreditAvailableAmount(value, currency string) {
	f.CreditAvailableAmount = append(f.CreditAvailableAmount, NewCurrencyAndAmount(value, currency))
}

func (f *FinancialCard1) SetInterestRatePercent(value string) {
	f.InterestRatePercent = (*PercentageRate)(&value)
}
